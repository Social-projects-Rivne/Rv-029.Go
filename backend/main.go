package main

import (
	"github.com/dgrijalva/jwt-go"
	//"time"
	"gopkg.in/yaml.v2"
	"log"
	"path/filepath"
	"io/ioutil"
	"fmt"
)

type JWTConfig struct {
	Claims map[string]interface{}
	Secret string
	Ttl int
	Refresh_ttl int
	Algo *jwt.SigningMethodHMAC
}

var config *JWTConfig

func main()  {
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
//
//type tokenizeable interface {
//	getClaims()
//}
//
//func generateToken(t *tokenizeable) (string, error)  {
//	/* Create the token */
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	/* Create a map to store our claims */
//	claims := token.Claims.(jwt.MapClaims)
//
//	if val, ok := config.Claims["iss"]; ok {
//		claims["iss"] = val
//	}
//
//	if val, ok := config.Claims["iat"]; ok {
//		if val == "" {
//			claims["iat"] = time.Now()
//		} else {
//			claims["iat"] = val
//		}
//	}
//
//	if val, ok := config.Claims["exp"]; ok {
//		if val == "" {
//			claims["exp"] = time.Duration(config.Ttl) * time.Minute
//		} else {
//			claims["exp"] = val
//		}
//	}
//
//	if val, ok := config.Claims["nbf"]; ok {
//		if val == "" {
//			claims["nbf"] = claims["iat"]
//		} else {
//			claims["nbf"] = time.Now().Add(time.Duration(int(val)) * time.Minute)
//		}
//	}
//
//	/* Set token claims */
//	claims["admin"] = true
//	claims["name"] = "Ado Kukic"
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	/* Sign the token with our secret */
//	return token.SignedString(config.Secret)
//}