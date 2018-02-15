package controllers

import (
	"log"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardCreateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := failedResponse{false, err.Error()}
		response.send(w)
		return
	}

<<<<<<< HEAD
	vars := mux.Vars(r)
	projectID, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		response := baseResponse{false, "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{
		gocql.TimeUUID(),
		projectID,
=======
	// Getting project data from request from "Validation"
	project := r.Context().Value("project")

	board := models.Board{
		gocql.TimeUUID(),
		project.ID,
		project.Name,
>>>>>>> origin/f52
		boardRequestData.Name,
		boardRequestData.Desc,
		time.Now(),
		time.Now(),
	}

	err = board.Insert()

	if err != nil {
<<<<<<< HEAD
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		log.Printf("Error while accessing to database: %v",err)
=======
		response := failedResponse{false, "Error while accessing to database"}
		response.send(w)
>>>>>>> origin/f52
		return
	}

	response := successResponse{true, "Board has created", nil}
	response.send(w)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardUpdateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := failedResponse{false, err.Error()}
		response.send(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := failedResponse{false, "Board ID is not valid"}
		response.send(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	board.FindByID()

	if boardRequestData.Name != "" {
		board.Name = boardRequestData.Name
	}

	if boardRequestData.Desc != "" {
		board.Desc = boardRequestData.Desc
	}

	board.UpdatedAt = time.Now()
	err = board.Update()

	if err != nil {
		response := failedResponse{false, "Error while accessing to database"}
		response.send(w)
		return
	}

	response := successResponse{true, "Board has updated", nil}
	response.send(w)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := failedResponse{false, "Project ID is not valid"}
		response.send(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = board.Delete()

	if err != nil {
		response := failedResponse{false, "Error while accessing to database"}
		response.send(w)
		return
	}

	response := successResponse{true, "Board has deleted", nil}
	response.send(w)
}

func SelectBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := failedResponse{false, "Board ID is not valid"}
		response.send(w)
		return
	}

	board := models.Board{}
	board.ID = id
	err = board.FindByID()

	if err != nil {
		response := failedResponse{false, "Error while accessing to database"}
		response.send(w)
		return
	}

	// TODO: refactor
	jsonResponse, _ := json.Marshal(board)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func BoardsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		response := failedResponse{false, "Project ID is not valid"}
		response.send(w)
		return
	}

	board := models.Board{}

	boardsList, err := board.List(projectId)

	if err != nil {
		response := failedResponse{false, "Error while accessing to database"}
		response.send(w)
		return
	}

	// TODO: refactor
	jsonResponse, _ := json.Marshal(boardsList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
