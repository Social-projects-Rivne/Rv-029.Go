package middlewares

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
	"log"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)

// Check if user have permission in project methods
func ProjectAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := r.Context().Value("user").(models.User)
		vars := mux.Vars(r)

		//TODO refactor logic middleware
		if user.Role == models.ROLE_OWNER {
			next.ServeHTTP(w, r)
		} else  {
			//check if project id is not empty
			if vars["project_id"] != ""{

				projectId , err := gocql.ParseUUID(vars["project_id"])
				if err != nil {
					log.Printf("Can't parse uuid in project middlaware")
					// return error response
					response := helpers.Response{Message: "Invalid UUID "}
					response.Failed(w)
					return
				}

				//check if user have access
				if user.Projects[projectId] != ""{
					next.ServeHTTP(w, r)
				}else{
					// return error response
					response := helpers.Response{Message: "You haven't permission"}
					response.Failed(w)
					return
				}


			}else{
				// return error response
				response := helpers.Response{Message: "You haven't permission"}
				response.Failed(w)
				return
			}

		}
	})
}

