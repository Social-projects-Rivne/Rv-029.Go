package controllers

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"time"
	"github.com/gorilla/mux"
	"log"
)


type projectResponse struct {
	Status  bool
	Message string
}


func ProjectsList(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}

	projects , err := project.GetAll()
	if err != nil{
		log.Fatal("Can't get all projects ",err)
	}

	projectJsonResponse, _ := json.Marshal(projects)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(projectJsonResponse)
}

func ShowProjects(w http.ResponseWriter, r *http.Request) {


	vars := mux.Vars(r)

	id , err := gocql.ParseUUID(vars["id"])
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}
	project := models.Project{}

	project.UUID = id
	project.FindByID()

	projectJsonResponse, err := json.Marshal(project)
	if err != nil{
		log.Println("Fail encode json in ShowProject method ",err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(projectJsonResponse)
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

	project := models.Project{gocql.TimeUUID(),projectRequestData.Name,time.Now(),time.Now()}

	project.Insert()

	jsonResponse, err := json.Marshal(projectResponse{
		Status:  true,
		Message: "Your project created",
	})
	if err != nil{
		log.Println("Fail encode json in StoreProject method ",err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

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

	vars := mux.Vars(r)

	id , err := gocql.ParseUUID(vars["id"])
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}
	project := models.Project{}

	project.UUID = id
	project.Name = projectRequestData.Name
	project.UpdatedAt = time.Now()

	project.Update()

	jsonResponse, err := json.Marshal(projectResponse{
		Status:  true,
		Message: "Your project updated",
	})
	if err != nil{
		log.Println("Fail encode json in UpdateProject method ",err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func DeleteProject(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	id , err := gocql.ParseUUID(vars["id"])
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}
	project := models.Project{}

	project.UUID = id
	project.Delete()

	jsonResponse, err := json.Marshal(projectResponse{
		Status:  true,
		Message: "Your project deleted",
	})
	if err != nil{
		log.Println("Fail encode json in DeleteProject method ",err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

