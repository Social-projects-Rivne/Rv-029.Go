package jwt

import (
	gojwt "github.com/dgrijalva/jwt-go"
)

type JWTConfig struct {
	Claims      gojwt.MapClaims
	Secret      string
	Ttl         int
	Refresh_ttl int
}

type WithClaims interface {
	GetClaims() map[string]interface{}
}

var Config *JWTConfig

// Generate jwt token with default and custom claims
func GenerateToken(wc WithClaims) (string, error) {
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
