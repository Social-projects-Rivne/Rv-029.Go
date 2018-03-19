package middlewares

import (
	"net/http"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)

func OwenrMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mux.CurrentRoute(r).GetName() != "" {
			user := r.Context().Value("user").(models.User)

			if user.Role != models.ROLE_OWNER {
				response := helpers.Response{Message: "Action denied. You have not OWNER role", StatusCode: http.StatusForbidden}
				response.Failed(w)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
