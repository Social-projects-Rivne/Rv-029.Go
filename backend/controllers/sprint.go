package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	//"fmt"
)

const DBError = "Error while accessing to database"

func CreateSprint(w http.ResponseWriter, r *http.Request) {

	var sprintRequestData validator.SprintCreateRequestData
	err := decodeAndValidate(r, &sprintRequestData)
	if err != nil {
		res := helpers.Response{Message: err.Error()}
		res.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Error in controllers/sprint/CreateSprint: %v", err)
		res := helpers.Response{Message: "Board ID is not valid"}
		res.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = models.BoardDB.FindByID(&board)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{
		gocql.TimeUUID(),
		board.ProjectID,
		board.ProjectName,
		board.ID,
		board.Name,
		sprintRequestData.Goal,
		sprintRequestData.Desc,
		sprintRequestData.Status,
		time.Now(),
		time.Now(),
	}

	err = models.SprintDB.Insert(&sprint)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	res := helpers.Response{Message: "Sprint has created"}
	res.Success(w)
}

func UpdateSprint(w http.ResponseWriter, r *http.Request) {

	var sprintRequestData validator.SprintUpdateRequestData
	err := decodeAndValidate(r, &sprintRequestData)
	if err != nil {
		res := helpers.Response{Message: err.Error()}
		res.Failed(w)
		return
	}

	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])
	if err != nil {
		log.Printf("Error in controllers/sprint/UpdateSprint: %+v", err)
		res := helpers.Response{Message: err.Error()}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId
	err = models.SprintDB.FindByID(&sprint)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	//If you want finish sprint
	if sprint.Status != sprintRequestData.Status && sprintRequestData.Status == models.SPRINT_STAUS_DONE {
		inProgressIssues, err := models.SprintDB.GetSprintIssuesInProgress(&sprint)
		if err != nil {
			res := helpers.Response{Message: DBError}
			res.Failed(w)
			return
		} else if len(inProgressIssues) > 0 {
			res := helpers.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Message: "Sprint contains not finished issues. Please finish them before finish the sprint",
			}
			res.Failed(w)
			return
		}
	}

	sprint.Goal = sprintRequestData.Goal
	sprint.Desc = sprintRequestData.Desc
	sprint.Status = sprintRequestData.Status
	sprint.UpdatedAt = time.Now()

	err = models.SprintDB.Update(&sprint)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	res := helpers.Response{Message: "Sprint has updated"}
	res.Success(w)
}

func DeleteSprint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])
	if err != nil {
		log.Printf("Error in controllers/sprint/DeleteSprint error: %+v", err)
		res := helpers.Response{Message: "Sprint ID is not valid"}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = models.SprintDB.Delete(&sprint)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	res := helpers.Response{Message: "Sprint has deleted"}
	res.Success(w)
}

func SelectSprint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])
	if err != nil {
		log.Printf("Error in controllers/sprint/SelectSprint error: %+v", err)
		res := helpers.Response{Message: "Sprint ID is not valid" }
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId
	err = models.SprintDB.FindByID(&sprint)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	res := helpers.Response{Message: "Done", Data: sprint}
	res.Success(w)
}

func SprintsList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Error in controllers/sprint/SprintsList: %+v", err)
		res := helpers.Response{Message: "Board ID is not valid"}
		res.Failed(w)
		return
	}

	sprintsList, err := models.SprintDB.List(boardId)
	if err != nil {
		res := helpers.Response{Message: DBError, StatusCode: http.StatusInternalServerError}
		res.Failed(w)
		return
	}

	res := helpers.Response{Message: "Done", Data: sprintsList}
	res.Success(w)
}
