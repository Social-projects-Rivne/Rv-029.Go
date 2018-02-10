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
		log.Fatal("Can't parse uuid ", err)
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

	jsonResponse, _ := json.Marshal(struct {
		Status  bool
		Message string
	}{
		true,
		"Board has created",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
