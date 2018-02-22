package jwt

import (
	gojwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type JWTConfig struct {
	Claims      gojwt.MapClaims
	Secret      string
	Ttl         int
	Refresh_ttl int
	Algo        *gojwt.SigningMethodHMAC
}

var Config *JWTConfig

type WithClaims interface {
	GetClaims() map[string]interface{}
}

// Generate jwt token with default and custom claims
func GenerateToken(wc WithClaims) (string, error) {
	filename, _ := filepath.Abs("./backend/config/jwt.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	Config = &JWTConfig{
		Algo: gojwt.SigningMethodHS256,
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Merge default and custom claims
	claims := Config.Claims
	for key, value := range wc.GetClaims() {
		claims[key] = value
	}
	//TODO: use ttl to generate exp
	// Create the token
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString([]byte(Config.Secret))
}
