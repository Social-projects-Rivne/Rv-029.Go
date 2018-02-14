package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"encoding/json"
)

const DBError = "Error while accessing to database"

func CreateSprint(w http.ResponseWriter, r *http.Request) {
	var sprintRequestData validator.SprintCreateRequestData

	err := decodeAndValidate(r, &sprintRequestData)

	if err != nil {
		res := failedResponse{false, err.Error()}
		res.send(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		res := failedResponse{false, "Invalid Board ID"}
		res.send(w)
		return
	}

	sprint := models.Sprint{
		gocql.TimeUUID(),
		boardId,
		sprintRequestData.Goal,
		sprintRequestData.Desc,
		sprintRequestData.Status,
		time.Now(),
		time.Now(),
	}

	err = sprint.Insert()

	if err != nil {
		res := failedResponse{false, DBError}
		res.send(w)
		return
	}

	res := successResponse{true, "Sprint has created", nil}
	res.send(w)
}

func UpdateSprint(w http.ResponseWriter, r *http.Request) {
	var sprintRequestData validator.SprintUpdateRequestData

	err := decodeAndValidate(r, &sprintRequestData)

	if err != nil {
		res := failedResponse{false, err.Error()}
		res.send(w)
		return
	}

	vars := mux.Vars(r)
	sprintId, _ := gocql.ParseUUID(vars["sprint_id"])

	sprint := models.Sprint{}
	sprint.ID = sprintId
	sprint.FindById()

	if sprintRequestData.Goal != "" {
		sprint.Goal = sprintRequestData.Goal
	}

	if sprintRequestData.Desc != "" {
		sprint.Desc = sprintRequestData.Desc
	}

	if sprintRequestData.Status != "" {
		sprint.Status = sprintRequestData.Status
	}

	sprint.UpdatedAt = time.Now()

	err = sprint.Update()

	if err != nil {
		res := failedResponse{false, DBError}
		res.send(w)
		return
	}

	res := successResponse{true, "Sprint has updated", nil}
	res.send(w)
}

func DeleteSprint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])

	if err != nil {
		res := failedResponse{false, "Sprint ID is not valid"}
		res.send(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = sprint.Delete()

	if err != nil {
		res := failedResponse{false, DBError}
		res.send(w)
		return
	}

	res := successResponse{true, "Sprint has deleted", nil}
	res.send(w)
}

func SelectSprint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])

	if err != nil {
		response := failedResponse{false, "Sprint ID is not valid"}
		response.send(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = sprint.FindById()

	if err != nil {
		res := failedResponse{false, DBError}
		res.send(w)
		return
	}

	jsonResponse, _ := json.Marshal(sprint)

	res := successResponse{true, "Done", jsonResponse}
	res.send(w)
}

func SprintsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		res := failedResponse{false, "Board ID is not valid"}
		res.send(w)
		return
	}

	sprint := models.Sprint{}

	sprintsList, err := sprint.List(boardId)

	if err != nil {
		res := failedResponse{false, DBError}
		res.send(w)
		return
	}

	jsonResponse, _ := json.Marshal(sprintsList)

	res := successResponse{true, "Done", jsonResponse}
	res.send(w)
}
