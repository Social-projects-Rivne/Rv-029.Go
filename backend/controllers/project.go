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
		log.Printf("Error in controllers/project.go . Can't get projects list, method: ProjectsList where: %s", err.Error())
		response := Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	response := Response{Message:"Projects list", Data: projects}
	response.Success(w)
}

func ShowProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project.go . Can't parse uuid, method: ShowProjects where: %s", err.Error())
		response := Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	project.FindByID()

	response := Response{Message:"Projects list", Data: project}
	response.Success(w)

}


func CreateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		response := Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(),projectRequestData.Name,time.Now(),time.Now()}
	project.Insert()

	user.AddRoleToProject(project.UUID,models.ROLE_OWNER)

	response := Response{Message: "Project has created"}
	response.Success(w)

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		response := Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project.go . Can't parse uuid, method: UpdateProject where: %s", err.Error())
		response := Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	project := models.Project{}

	project.UUID = projectId
	project.Name = projectRequestData.Name
	project.UpdatedAt = time.Now()

	project.Update()

	response := Response{Message: "Project has updated"}
	response.Success(w)

}

func DeleteProject(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	projectId , err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project.go . Can't parse uuid, method: DeleteProject where: %s", err.Error())
		response := Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}
	project := models.Project{}
	project.UUID = projectId

	user := r.Context().Value("user").(models.User)
	user.DeleteProject(project.UUID)

	project.Delete()

	response := Response{Message: "Project has deleted"}
	response.Success(w)
}

