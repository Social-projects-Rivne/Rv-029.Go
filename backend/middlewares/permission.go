package middlewares

import (
	"net/http"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)

type rules map[string][]string

var rulesMap = rules{
	"project.create":    []string{"project.create"},
	"user.permissions.add" : []string{"user.permissions.manage"},
	"user.permissions.remove" : []string{"user.permissions.manage"},
	"user.permissions.update" : []string{"user.permissions.manage"},
}

func CheckUserPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mux.CurrentRoute(r).GetName() != "" {
			user := r.Context().Value("user").(models.User)
			routePermissions := rulesMap[mux.CurrentRoute(r).GetName()]

			if user.Role != models.ROLE_OWNER && len(routePermissions) > 0 {
				for _, permission := range routePermissions {
					if !user.HasPermission(permission) {
						response := helpers.Response{Message: "Action denied. No permission", StatusCode: http.StatusForbidden}
						response.Failed(w)
						return
					}
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
