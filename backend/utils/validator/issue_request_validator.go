package validator

import (
	"github.com/gocql/gocql"
	"net/http"
	//"time"
)

//IssueRequestData struct
type IssueRequestData struct {
	*baseValidator
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	UserID    gocql.UUID `json:"user_id"`
	SprintID  gocql.UUID `json:"sprint_id"`
	BoardID   gocql.UUID `json:"board_id"`
}

//Validate .
func (b *IssueRequestData) Validate(r *http.Request) error {
	// just for Interface implementation
	return nil
}