package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type boardSuccessResponse struct {
	Status bool
	Message string
}

func StoreBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			false,
			err.Error(),
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Fatal(err)
	}

	board := models.Board{
		gocql.TimeUUID(),
		projectId,
		boardRequestData.Name,
		boardRequestData.Desc,
		time.Now(),
		time.Now(),
	}

	board.Insert()

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Board has created",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			false,
			err.Error(),
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Fatal(err)
	}

	board := models.Board{}
	board.ID = boardId
	board.Name = boardRequestData.Name
	board.Desc = boardRequestData.Desc
	board.UpdatedAt = time.Now()
	board.Update()

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Board has updated",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Fatal(err)
	}

	board := models.Board{}
	board.ID = boardId
	board.Delete()

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Board has deleted",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
