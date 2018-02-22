package validator

import (
	"context"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
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
		log.Printf("Error in utils/validator/board_create_request_validator.go error: %+v",err)
		return err
	}

	err = b.ValidateRequired(b.Desc)
	if err != nil {
		log.Printf("Error in utils/validator/board_create_request_validator.go error: %+v",err)
		return err
	}

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		log.Printf("Error in utils/validator/board_create_request_validator.go error: %+v",err)
		return err
	}

	project := models.Project{}
	project.UUID = projectId
	err = project.FindByID()

	if err != nil {
		log.Printf("Error in utils/validator/board_create_request_validator.go error: %+v",err)		
		return err
	}

	// Adding project data to request
	ctx := context.WithValue(r.Context(), "project", project)
	r.WithContext(ctx)

	return nil
}
