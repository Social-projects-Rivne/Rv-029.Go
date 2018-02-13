package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
)

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
		res := baseResponse{false, "Error while accessing to db"}
		res.Failed(w)
		return
	}

	res := baseResponse{true, "Sprint has created"}
	res.Success(w)

}
