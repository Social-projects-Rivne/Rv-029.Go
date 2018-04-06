package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func ProjectsList(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}

	projects, err := models.ProjectDB.GetProjectList(&project)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Projects list", Data: projects}
	response.Success(w)
}

func ShowProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	err = models.ProjectDB.FindByID(&project)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Projects list", Data: project}
	response.Success(w)

}

func ProjectUsersList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	users, err := models.UserDB.GetProjectUsersList(projectId)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "List of users to assigned to current project", Data: users}
	response.Success(w)

}

func UsersToAddProjectList(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//
	//projectId , err := gocql.ParseUUID(vars["project_id"])
	//if err != nil {
	//	log.Printf("Error in controllers/project error: %+v",err)
	//	response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
	//	response.Failed(w)
	//	return
	//}

	users, err := models.UserDB.List()
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	notOwners := []models.User{}
	for _, u := range users {
		if u.Role != models.ROLE_OWNER {
			notOwners = append(notOwners, u)
		}
	}

	response := helpers.Response{Message: "List of users to assigned to current project", Data: notOwners}
	response.Success(w)

}

func CreateProject(w http.ResponseWriter, r *http.Request) {

	var projectRequestData validator.ProjectRequestData

	err := decodeAndValidate(r, &projectRequestData)
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)

	project := models.Project{gocql.TimeUUID(), projectRequestData.Name, time.Now(), time.Now()}

	// TODO some transaction for project Insert
	err = models.ProjectDB.Insert(&project)
	if err != nil {
		response := helpers.Response{Message: "Can't create project", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	err = models.UserDB.AddRoleToProject(project.UUID, models.ROLE_OWNER, user.UUID)
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
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error()), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	project := models.Project{}

	project.UUID = projectId
	project.Name = projectRequestData.Name
	project.UpdatedAt = time.Now()

	err = models.ProjectDB.Update(&project)
	if err != nil {
		response := helpers.Response{Message: "Can't update project", StatusCode: http.StatusInternalServerError}
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
		log.Printf("Error in controllers/project error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}
	project := models.Project{}
	project.UUID = projectId

	user := r.Context().Value("user").(models.User)
	err = models.UserDB.DeleteProject(project.UUID, user.UUID)
	if err != nil {
		response := helpers.Response{Message: "Can't delete user project access", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	err = models.ProjectDB.Delete(&project)
	if err != nil {
		response := helpers.Response{Message: "Can't delete project", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Project has deleted"}
	response.Success(w)
}

func ProjectAddUser(w http.ResponseWriter, r *http.Request) {
	var userToProjectRequestData validator.UserProjectRequestData
	err := decodeAndValidate(r, &userToProjectRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/project. Can`t parse project UUID. error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	userID, err := gocql.ParseUUID(userToProjectRequestData.UserID)
	if err != nil {
		log.Printf("Error in controllers/project. Can`t parse user UUID. error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	user := models.User{}
	user.UUID = userID
	err = models.UserDB.FindByID(&user)
	if err != nil {
		log.Printf("Error in controllers/project. Can`t find user by UUID. error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	if err, ok := user.Projects[projectId]; !ok {
		log.Printf("Error in controllers/project. No map with projectId key. error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %+v", err)}
		response.Failed(w)
		return
	}

	user.Projects[projectId] = userToProjectRequestData.Role

	models.UserDB.Update(&user)
	if err != nil {
		log.Printf("Error in controllers/project. Can`t update user. error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error %s", err.Error())}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: fmt.Sprintf("User %s was added to project", user.Email)}
	response.Success(w)
}
