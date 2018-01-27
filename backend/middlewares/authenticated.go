package middlewares

import (
	"net/http"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"fmt"
	"encoding/json"
)

// Check if user authenticated
func AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if jwt present & valid
		err := jwtMiddleware.CheckJWT(w, r)
		if err != nil {
			// return error response
			response := struct {
				Status bool
				Message string
			}{
				Status: false,
				Message: "You are not authorized.",
			}
			jsonResponse, _ := json.Marshal(response)

			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		} else {
			// get user claims from token
			user := r.Context().Value("user")
			claims := user.(*jwt.Token).Claims.(jwt.MapClaims)

			fmt.Println(claims["id"])
			fmt.Println(claims["email"])

			//TODO: get id or email from claims
			//TODO: check if user exists
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}
	})
}

// Decode jwt token and check if it is valid
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt2.Config.Secret), nil
	},
	//TODO: make custom error handler
	SigningMethod: jwt2.Config.Algo,
})