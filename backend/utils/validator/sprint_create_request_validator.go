package validator

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"context"
)

type SprintCreateRequestData struct {
	*baseValidator
	Goal        string `json:"goal"`
	Desc        string `json:"desc"`
	Status      string `json:"status"`
}

func (s *SprintCreateRequestData) Validate(r *http.Request) error {
	var err error

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
	boardId, _ := gocql.ParseUUID(vars["board_id"])

	board := models.Board{}
	board.ID = boardId
	err = board.FindByID()

	if err != nil {
		return err
	}

	// todo: add project
	ctx := context.WithValue(r.Context(), "board", board)
	r.WithContext(ctx)






	return nil
}
