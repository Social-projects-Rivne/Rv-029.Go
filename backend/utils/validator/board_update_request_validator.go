package validator

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type BoardUpdateRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (b *BoardUpdateRequestData) Validate(r *http.Request) error {
	return validateBoardId(r)
}

func validateBoardId(r *http.Request) error {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil{
		log.Printf("Error in utils/validator/board_update_request_validator.go error: %+v",err)
		return err
	}

	var boardName string

	db.GetInstance().Session.
		Query(`SELECT name FROM boards where id = ? LIMIT 1;`, boardId).
		Consistency(gocql.One).Scan(&boardName)

	if boardName == "" {
		err := fmt.Errorf("There is no board with ID %v", boardId)
		log.Printf("Error in utils/validator/board_update_request_validator.go error: %+v",err)
		return err
	}

	return nil
}
