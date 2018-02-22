package middlewares

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
	"log"
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

				projectId , err := gocql.ParseUUID(vars["id"])
				if err != nil {
					log.Printf("Can't parse uuid in project middlaware")
					// return error response
					//TODO refactor error response make some helper
					response := struct {
						Status bool
						Message string
					}{
						Status: false,
						Message: "You haven't permission",
					}
					jsonResponse, err := json.Marshal(response)
					if err != nil{
						log.Printf("Error in middlewares/project.go error: %+v",err)
						return
					}

					w.WriteHeader(http.StatusForbidden)
					w.Header().Set("Content-Type", "application/json")
					w.Write(jsonResponse)
					return
				}

				//check if user have access
				if user.Projects[projectId] != ""{
					next.ServeHTTP(w, r)
				}else{
					// return error response
					//TODO refactor error response make some helper
					response := struct {
						Status bool
						Message string
					}{
						Status: false,
						Message: "You haven't permission",
					}
					jsonResponse, err := json.Marshal(response)
					if err != nil{
						log.Printf("Error in middlewares/project.go error: %+v",err)
						return
					}

					w.WriteHeader(http.StatusForbidden)
					w.Header().Set("Content-Type", "application/json")
					w.Write(jsonResponse)
					return
				}


			}else{
				// return error response
				//TODO refactor error response make some helper
				response := struct {
					Status bool
					Message string
				}{
					Status: false,
					Message: "You haven't permission",
				}
				jsonResponse, err := json.Marshal(response)
				if err != nil{
					log.Printf("Error in middlewares/project.go error: %+v",err)
					return
				}

				w.WriteHeader(http.StatusForbidden)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonResponse)
				return
			}




		}
	})
}

