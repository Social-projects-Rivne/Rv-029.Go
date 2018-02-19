package validator

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SprintCreateRequestData struct {
	*baseValidator
	Goal   string `json:"goal"`
	Desc   string `json:"desc"`
	Status string
}

func (s *SprintCreateRequestData) Validate(r *http.Request) error {
	var err error

	s.Status = "TODO"

	err = s.ValidateRequired(s.Goal)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	err = s.ValidateRequired(s.Desc)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Invalid Board ID: %v\n", err.Error())
		return err
	}

	board := models.Board{}
	board.ID = boardId
	err = board.FindByID()
	if err != nil {
		return err
	}

	return nil
}
