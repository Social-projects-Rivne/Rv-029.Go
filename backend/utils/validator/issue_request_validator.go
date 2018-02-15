package validator

import (
	"net/http"
	"errors"
	"fmt"

	"github.com/gocql/gocql"
	//"time"
)

//CreateIssueRequestData struct
type CreateIssueRequestData struct {
	*baseValidator
	Name   string     `json:"name"`
	Status string     `json:"status"`
	UserID gocql.UUID `json:"user_id"`
}

//Validate .
func (b *CreateIssueRequestData) Validate(r *http.Request) error {
	if b.Name == "" {
		return errors.New(fmt.Sprintln("Issue.Name is empty"))
	}
	if b.Status == "" {
		return errors.New(fmt.Sprintln("Issue.Status is empty"))
	}
	if b.UserID.String() == "" {
		return errors.New(fmt.Sprintln("Issue.UserID is empty"))
	}
	return nil
}
