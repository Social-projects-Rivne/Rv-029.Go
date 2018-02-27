package controllers

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ProjectsList(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}

	projects, err := models.ProjectDB.GetProjectList(&project)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Projects list", Data: projects}
	response.Success(w)
}

func ShowProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	err = models.ProjectDB.FindByID(&project)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Projects list", Data: project}
	response.Success(w)

}

func CreateProject(w http.ResponseWriter, r *http.Request) {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(), projectRequestData.Name, time.Now(), time.Now()}

	err = models.ProjectDB.Insert(&project)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: "Can't create project", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	err = user.AddRoleToProject(project.UUID, models.ROLE_OWNER)
	if err != nil {
		response := helpers.Response{Message: "Can't add role to project", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has created"}
	response.Success(w)

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	project := models.Project{}

	project.UUID = projectId
	project.Name = projectRequestData.Name
	project.UpdatedAt = time.Now()

	if err = models.ProjectDB.Update(&project);err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has updated"}
	response.Success(w)

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}
	project := models.Project{}
	project.UUID = projectId

	user := r.Context().Value("user").(models.User)
	err = user.DeleteProject(project.UUID)
	if err != nil {
		response := helpers.Response{Message: "Can't delete user project access", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	if err = models.ProjectDB.Delete(&project); err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has deleted"}
	response.Success(w)
}
