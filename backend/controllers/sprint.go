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
		res := baseResponse{false, err.Error()}
		res.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		res := baseResponse{false, "Invalid Board ID"}
		res.Failed(w)
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
		res := baseResponse{false, DBError}
		res.Failed(w)
		return
	}

	res := baseResponse{true, "Sprint has created"}
	res.Success(w)
}

func UpdateSprint(w http.ResponseWriter, r *http.Request) {
	var sprintRequestData validator.SprintUpdateRequestData

	err := decodeAndValidate(r, &sprintRequestData)

	if err != nil {
		res := baseResponse{false, err.Error()}
		res.Failed(w)
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
		res := baseResponse{false, DBError}
		res.Failed(w)
		return
	}

	res := baseResponse{true, "Sprint has updated"}
	res.Success(w)
}

func DeleteSprint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])

	if err != nil {
		res := baseResponse{false, "Sprint ID is not valid"}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = sprint.Delete()

	if err != nil {
		res := baseResponse{false, DBError}
		res.Failed(w)
		return
	}

	res := baseResponse{true, "Sprint has deleted"}
	res.Success(w)
}

func SelectSprint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sprintId, err := gocql.ParseUUID(vars["sprint_id"])

	if err != nil {
		response := baseResponse{false, "Sprint ID is not valid"}
		response.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintId

	err = sprint.FindById()

	if err != nil {
		res := baseResponse{false, DBError}
		res.Failed(w)
		return
	}

	jsonResponse, _ := json.Marshal(sprint)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func SprintsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		res := baseResponse{false, "Board ID is not valid"}
		res.Failed(w)
		return
	}

	sprint := models.Sprint{}

	sprintsList, err := sprint.List(boardId)

	if err != nil {
		res := baseResponse{false, DBError}
		res.Failed(w)
		return
	}

	jsonResponse, _ := json.Marshal(sprintsList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
