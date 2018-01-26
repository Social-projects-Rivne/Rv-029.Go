package jwt

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"time"

	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
	"gopkg.in/yaml.v2"
)

type JWTConfig struct {
	Claims      map[string]interface{}
	Secret      string
	Ttl         int
	Refresh_ttl int
	Algo        *gojwt.SigningMethodHMAC
}

var config *JWTConfig

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *gojwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	},
	SigningMethod: gojwt.SigningMethodHS256,
})

func init() {
	filename, _ := filepath.Abs("./backend/config/jwt.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config = &JWTConfig{
		Algo: gojwt.SigningMethodHS256,
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

type WithClaims interface {
	GetClaims() map[string]interface{}
}

func GenerateToken(wc WithClaims) (string, error) {
	/* Create the token */
	token := gojwt.New(gojwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(gojwt.MapClaims)

	if val, ok := config.Claims["iss"]; ok {
		claims["iss"] = val
	}

	if val, ok := config.Claims["iat"]; ok {
		if val == nil {
			claims["iat"] = time.Now().Unix()
		} else {
			claims["iat"] = val
		}
	}

	if val, ok := config.Claims["exp"]; ok {
		if val == nil {
			claims["exp"] = time.Now().Add(time.Duration(config.Ttl) * time.Minute).Unix()
		} else {
			claims["exp"] = val
		}
	}

	if val, ok := config.Claims["nbf"]; ok {
		if val == nil {
			claims["nbf"] = claims["iat"]
		} else {
			claims["nbf"] = time.Now().Add(time.Duration(val.(int)) * time.Minute).Unix()
		}
	}

	for key, val := range wc.GetClaims() {
		ref := reflect.TypeOf(wc)
		if _, ok := ref.MethodByName("Get" + key); !ok {
			claims[key] = val
		} else {
			claims[key] = reflect.ValueOf(wc).MethodByName("Get" + key).Call([]reflect.Value{})
		}
	}

	/* Sign the token with our secret */
	return token.SignedString([]byte(config.Secret))
}
