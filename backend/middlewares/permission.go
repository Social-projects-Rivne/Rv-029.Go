package middlewares

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)

type rules map[string][]string

var rulesMap = rules{
	"project.create":    []string{models.PERMISSION_CREATE_PROJECTS},
	"project.update":    []string{models.PERMISSION_UPDATE_PROJECTS},
	"project.delete":    []string{models.PERMISSION_DELETE_PROJECTS},

	"board.create":    []string{models.PERMISSION_CREATE_BOARDS},
	"board.update":    []string{models.PERMISSION_UPDATE_BOARDS},
	"board.delete":    []string{models.PERMISSION_DELETE_BOARDS},

	"sprint.create":    []string{models.PERMISSION_CREATE_SPRINTS},
	"sprint.update":    []string{models.PERMISSION_UPDATE_SPRINTS},
	"sprint.delete":    []string{models.PERMISSION_DELETE_SPRINTS},
	"sprint.issue.add":    []string{models.PERMISSION_ADD_ISSUES_TO_SPRINTS},

	"issue.create":    []string{models.PERMISSION_CREATE_ISSUES},
	"issue.update":    []string{models.PERMISSION_UPDATE_ISSUES},
	"issue.delete":    []string{models.PERMISSION_DELETE_ISSUES},

	"user.permissions.add" : []string{models.PERMISSION_MANAGE_USER_PERMISSIONS},
	"user.permissions.remove" : []string{models.PERMISSION_MANAGE_USER_PERMISSIONS},
	"user.permissions.update" : []string{models.PERMISSION_MANAGE_USER_PERMISSIONS},
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
