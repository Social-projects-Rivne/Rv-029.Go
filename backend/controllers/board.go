package controllers

import (
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func CreateBoard(w http.ResponseWriter, r *http.Request) {

	boardRequestData := new(validator.BoardCreateRequestData)
	err := decodeAndValidate(r, boardRequestData)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	project := models.Project{}
	project.UUID = projectId
	err = models.ProjectDB.FindByID(&project)

	if err != nil {
		response := helpers.Response{Message: "Project ID is not valid",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	board := models.Board{
		gocql.TimeUUID(),
		project.UUID,
		project.Name,
		boardRequestData.Name,
		boardRequestData.Desc,
		time.Now(),
		time.Now(),
	}

	err = models.BoardDB.Insert(&board)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)
		response := helpers.Response{Message: "Error while accessing to database",StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has created"}
	response.Success(w)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardUpdateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in controllers/board error: %+v", err)
		response := helpers.Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = models.BoardDB.FindByID(&board)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v", err)
		response := helpers.Response{Message: fmt.Sprintf("Error in controllers/board error: %+v", err)}
		response.Failed(w)
		return
	}

	board.Name = boardRequestData.Name
	board.Desc = boardRequestData.Desc
	board.UpdatedAt = time.Now()
	err = models.BoardDB.Update(&board)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has updated"}
	response.Success(w)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: "Project ID is not valid", StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = models.BoardDB.Delete(&board)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has deleted"}
	response.Success(w)
}

func SelectBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: "Board ID is not valid", StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = id
	err = models.BoardDB.FindByID(&board)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Data: board}
	response.Success(w)
}

func BoardsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)		
		response := helpers.Response{Message: "Project ID is not valid", StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	boardsList, err := models.BoardDB.List(projectId)

	if err != nil {
		log.Printf("Error in controllers/board error: %+v",err)	
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Data: boardsList}
	response.Success(w)
}
