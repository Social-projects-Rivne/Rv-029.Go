package validator

import (
	"context"
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
	Status string `json:"status"`
}

func (s *SprintCreateRequestData) Validate(r *http.Request) error {
	var err error

	err = s.ValidateRequired(s.Goal)

	if err != nil {
		log.Printf("Error in utils/validator/sprint_create_request_validator.go error: %+v",err)
		return err
	}

	err = s.ValidateRequired(s.Desc)

	if err != nil {
		log.Printf("Error in utils/validator/sprint_create_request_validator.go error: %+v",err)
		return err
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in utils/validator/sprint_create_request_validator.go error: %+v",err)
		return err
	}

	board := models.Board{}
	board.ID = boardId
	err = board.FindByID()

	if err != nil {
		log.Printf("Error in utils/validator/sprint_create_request_validator.go error: %+v",err)		
		return err
	}

	ctx := context.WithValue(r.Context(), "board", board)
	r.WithContext(ctx)

	return nil
}
