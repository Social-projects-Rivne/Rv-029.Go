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
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)


func ProjectsList(w http.ResponseWriter) {
	project := models.Project{}

	projects , err := project.GetProjectList()
	if err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()),StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message:"Projects list", Data: projects}
	response.Success(w)
}

func ShowProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	err = project.FindByID()
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()) ,StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message:"Projects list", Data: project}
	response.Success(w)

}


func CreateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(),projectRequestData.Name,time.Now(),time.Now()}

	// TODO some transaction for project Insert
	err = project.Insert()
	if err != nil {
		response := helpers.Response{Message: "Can't create project",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	err = user.AddRoleToProject(project.UUID,models.ROLE_OWNER)
	if err != nil {
		response := helpers.Response{Message: "Can't add role to project",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has created"}
	response.Success(w)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()),StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	project := models.Project{}

	project.UUID = projectId
	project.Name = projectRequestData.Name
	project.UpdatedAt = time.Now()

	err = project.Update()
	if err != nil {
		response := helpers.Response{Message: "Can't update project",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has updated"}
	response.Success(w)

}

func DeleteProject(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}
	project := models.Project{}
	project.UUID = projectId

	user := r.Context().Value("user").(models.User)
	err = user.DeleteProject(project.UUID)
	if err != nil {
		response := helpers.Response{Message: "Can't delete user project access",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	err = project.Delete()
	if err != nil {
		response := helpers.Response{Message: "Can't delete project",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has deleted"}
	response.Success(w)
}

