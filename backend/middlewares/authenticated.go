package middlewares

import (
	"context"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	jwt2 "github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
	"net/http"
	"log"
)

// Check if user authenticated
func AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if jwt present & valid
		err := jwtMiddleware.CheckJWT(w, r)
		if err != nil {
			// return error response
			response := struct {
				Status  bool
				Message string
			}{
				Status:  false,
				Message: "You are not authorized.",
			}
			jsonResponse, err := json.Marshal(response)
			if err != nil{
				log.Printf("Error in middlewares/authenticated.go error: %+v",err)
				return
			}

			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		} else {
			// get user claims from token
			userContext := r.Context().Value("user")
			claims := userContext.(*jwt.Token).Claims.(jwt.MapClaims)

			UUID, err := gocql.ParseUUID(claims["UUID"].(string))
			if err != nil {
				response := struct {
					Status  bool
					Message string
				}{
					Status:  false,
					Message: "Invalid user JWT token claims",
				}
				jsonResponse, err := json.Marshal(response)
				if err != nil{
					log.Printf("Error in middlewares/authenticated.go error: %+v",err)
					return
				}

				w.WriteHeader(http.StatusForbidden)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonResponse)
			}

			currentUser := models.User{}
			currentUser.UUID = UUID
			err = currentUser.FindByID()
			if err != nil {
				response := struct {
					Status  bool
					Message string
				}{
					Status:  false,
					Message: "You are not authorized.",
				}
				jsonResponse, err := json.Marshal(response)
				if err != nil{
					log.Printf("Error in middlewares/authenticated.go error: %+v",err)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonResponse)

				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				//Add User instance to context
				ctx := context.WithValue(r.Context(), "user", currentUser)

				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
	})
}

// Decode jwt token and check if it is valid
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt2.Config.Secret), nil
	},
	//TODO: make custom error handler
	SigningMethod: jwt.SigningMethodHS256,
})
