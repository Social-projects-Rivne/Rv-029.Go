package validator

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func validateProjectId(r *http.Request) error {
	vars := mux.Vars(r)
	projectId, _ := gocql.ParseUUID(vars["project_id"])

	var projectName string

	db.GetInstance().Session.
		Query(`SELECT name FROM projects where id = ? LIMIT 1;`, projectId).
		Consistency(gocql.One).Scan(&projectName)

	if projectName == "" {
		err := fmt.Errorf("There is no project with ID %v", projectId)
		log.Printf(err.Error())
		return err
	}

	return nil
}

func validateBoardId(r *http.Request) error {
	vars := mux.Vars(r)
	boardId, _ := gocql.ParseUUID(vars["board_id"])

	var boardName string

	db.GetInstance().Session.
		Query(`SELECT name FROM boards where id = ? LIMIT 1;`, boardId).
		Consistency(gocql.One).Scan(&boardName)

	if boardName == "" {
		err := fmt.Errorf("There is no board with ID %v", boardId)
		log.Printf(err.Error())
		return err
	}

	return nil
}

type boardRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type BoardBaseRequestData struct {
	*boardRequestData
}

func (b *BoardBaseRequestData) Validate(r *http.Request) error {
	return validateBoardId(r)
}

type BoardCreateRequestData struct {
	*boardRequestData
}

func (b *BoardCreateRequestData) Validate(r *http.Request) error {
	var err error

	err = b.ValidateRequired(b.Name)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	err = b.ValidateRequired(b.Desc)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return validateProjectId(r)
}
