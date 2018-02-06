package controllers

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"encoding/json"
	"reflect"
	//"github.com/gocql/gocql"
	//"time"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"fmt"
)


type projectResponse struct {
	Status  bool
	Message string
}


func IndexProject(w http.ResponseWriter, r *http.Request) {

}

func StoreProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			Status:  false,
			Message: err.Error(),
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	user := r.Context().Value("user")
	w.Write([]byte("Project"))
	//
	userUUID := reflect.ValueOf(user).FieldByName("UUID")
	fmt.Println(userUUID)
	//q := gocql.UUID{}
	//project := models.Project{gocql.TimeUUID(),q.String(userUUID),"firstProject",time.Now(),time.Now()}
	//
	//b := models.BaseModel{}
	//
	//b.Insert("projects",project)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

}

func DestroyProject(w http.ResponseWriter, r *http.Request)  {

}

