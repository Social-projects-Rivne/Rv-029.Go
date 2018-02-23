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


func ProjectsList(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}

	projects , err := project.GetProjectList()
	if err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
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
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	if err = project.FindByID();err !=nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
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
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(),projectRequestData.Name,time.Now(),time.Now()}
	if err = project.Insert();err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	user.AddRoleToProject(project.UUID,models.ROLE_OWNER)

	response := helpers.Response{Message: "Project has created"}
	response.Success(w)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
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

	if err = project.Update();err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
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
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}
	project := models.Project{}
	project.UUID = projectId

	user := r.Context().Value("user").(models.User)
	user.DeleteProject(project.UUID)

	if err = project.Delete(); err != nil{
		log.Printf("Error in controllers/project error: %+v",err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has deleted"}
	response.Success(w)
}

