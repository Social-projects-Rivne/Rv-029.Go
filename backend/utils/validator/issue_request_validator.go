package validator

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
	//"time"
)

//CreateIssueRequestData struct
type CreateIssueRequestData struct {
	*baseValidator
	Name        string     `json:"name"`
	Description string     `json:"description"`
	UserID      gocql.UUID `json:"user_id"`
	Estimate    int        `json:"estimate"`
	Status      string     `json:"status"`
	SprintID    gocql.UUID `json:"sprint_id"`
}

//Validate .
func (b *CreateIssueRequestData) Validate(r *http.Request) error {
	if b.Name == "" {
		return errors.New(fmt.Sprintln("while decoding json error: Issue.Name is empty"))
	}
	if b.Description == "" {
		return errors.New(fmt.Sprintln("while decoding json error: Issue.Description is empty"))
	}
	return nil
}
