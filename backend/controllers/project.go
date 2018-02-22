package controllers

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"time"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)


func ProjectsList(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}

	projects , err := project.GetProjectList()
	if err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error %s", err.Error())}
		response.send(w)
		return
	}

	response := successResponse{true,"Projects list",projects}
	response.send(w)
}

func ShowProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error %s", err.Error())}
		response.send(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	if err = project.FindByID();err !=nil{
		log.Printf("Error in controllers/project error: %+v",err)
		return
	}

	response := successResponse{true,"Projects list",project}
	response.send(w)

}


func CreateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		response := failedResponse{false, err.Error()}
		response.send(w)
		return
	}

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(),projectRequestData.Name,time.Now(),time.Now()}
	if err = project.Insert();err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		return
	}

	user.AddRoleToProject(project.UUID,models.ROLE_OWNER)

	response := successResponse{true, "Project has created",nil}
	response.send(w)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		response := failedResponse{false, err.Error()}
		response.send(w)
		return
	}

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error %s", err.Error())}
		response.send(w)
		return
	}

	project := models.Project{}

	project.UUID = projectId
	project.Name = projectRequestData.Name
	project.UpdatedAt = time.Now()

	if err = project.Update();err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		return
	}

	response := successResponse{true, "Project has updated",nil}
	response.send(w)

}

func DeleteProject(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error %s", err.Error())}
		response.send(w)
		return
	}
	project := models.Project{}
	project.UUID = projectId

	user := r.Context().Value("user").(models.User)
	user.DeleteProject(project.UUID)

	if err = project.Delete(); err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		return
	}

	response := successResponse{true, "Project has deleted",nil}
	response.send(w)
}

