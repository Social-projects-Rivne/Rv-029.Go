package controllers

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
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
		res := helpers.Response{Message: "Board ID is not valid"}
		res.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = board.FindByID()

	if err != nil {
		res := helpers.Response{Message: DBError}
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

	err = sprint.Insert()

	if err != nil {
		res := helpers.Response{Message: DBError}
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
		res := helpers.Response{Message: "Sprint ID is not valid"}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId
	err = sprint.FindById()

	if err != nil {
		res := helpers.Response{Message: DBError}
		res.Failed(w)
		return
	}

	sprint.Goal = sprintRequestData.Goal
	sprint.Desc = sprintRequestData.Desc
	sprint.Status = sprintRequestData.Status
	sprint.UpdatedAt = time.Now()
	err = sprint.Update()

	if err != nil {
		res := helpers.Response{Message: DBError}
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
		res := helpers.Response{Message: "Sprint ID is not valid"}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = sprint.Delete()

	if err != nil {
		res := helpers.Response{Message: DBError}
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
		response := helpers.Response{Message: "Sprint ID is not valid"}
		response.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = sprint.FindById()
	if err != nil {
		res := helpers.Response{Message: DBError}
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
		res := helpers.Response{Message: "Board ID is not valid"}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}

	sprintsList, err := sprint.List(boardId)

	if err != nil {
		res := helpers.Response{Message: DBError}
		res.Failed(w)
		return
	}

	res := helpers.Response{Message: "Done", Data: sprintsList}
	res.Success(w)
}
