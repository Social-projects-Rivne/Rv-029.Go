package middlewares

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"encoding/json"
)

// Check if user authenticated
func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if user.Role == models.ROLE_OWNER {
			next.ServeHTTP(w,r)
		} else {
			//TODO: check if there are any project where user assigned
			// return error response
			response := struct {
				Status bool
				Message string
			}{
				Status: false,
				Message: "You haven't permission",
			}
			jsonResponse, _ := json.Marshal(response)

			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}
	})
}

