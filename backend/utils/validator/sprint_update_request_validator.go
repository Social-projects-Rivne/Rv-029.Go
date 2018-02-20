package validator

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SprintUpdateRequestData struct {
	*baseValidator
	Goal   string `json:"goal"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

func (s *SprintUpdateRequestData) Validate(r *http.Request) error {
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

	err = s.ValidateRequired(s.Status)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return validateSprintId(r)
}

func validateSprintId(r *http.Request) error {
	vars := mux.Vars(r)
	sprintId, _ := gocql.ParseUUID(vars["sprint_id"])

	var sprintGoal string

	db.GetInstance().Session.
		Query(`SELECT goal FROM sprints where id = ? LIMIT 1;`, sprintId).
		Consistency(gocql.One).Scan(&sprintGoal)

	if sprintGoal == "" {
		err := fmt.Errorf("There is no sprint with ID %q", sprintId)
		log.Printf(err.Error())
		return err
	}

	return nil
}
