package validator

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type BoardCreateRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"desc"`
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