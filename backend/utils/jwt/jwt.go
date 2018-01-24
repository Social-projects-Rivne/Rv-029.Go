package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"io/ioutil"
	"log"
	"reflect"
)

type JWTConfig struct {
	Claims map[string]interface{}
	Secret string
	Ttl int
	Refresh_ttl int
	Algo *jwt.SigningMethodHMAC
}

var config *JWTConfig


func init()  {
	filename, _ := filepath.Abs("./config/jwt.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config = &JWTConfig{
		Algo: jwt.SigningMethodHS256,
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

type withClaims interface {
	GetClaims() map[string]interface{}
}

func generateToken(wc withClaims) (string, error)  {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

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
	return token.SignedString(config.Secret)
}