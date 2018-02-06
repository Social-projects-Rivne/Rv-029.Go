package controllers

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"time"
	"fmt"
	"github.com/gorilla/mux"
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

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(),user.UUID,projectRequestData.Name,time.Now(),time.Now()}
	b := models.BaseModel{}

	b.Insert("projects",project)

	jsonResponse, _ := json.Marshal(projectResponse{
		Status:  true,
		Message: "Your project created",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	b := models.BaseModel{}
	project := models.Project{}

	b.Where("id","=",vars["id"])
	b.Select("projects",project)

	fmt.Println(project)

}

func DestroyProject(w http.ResponseWriter, r *http.Request)  {

}

