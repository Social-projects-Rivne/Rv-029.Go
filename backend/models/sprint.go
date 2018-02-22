package models

import (
	"github.com/gocql/gocql"
	"log"
	"time"
)

const (
	SPRINT_STAUS_TODO = "TODO"
	SPRINT_STAUS_IN_PROGRESS = "In Progress"
	SPRINT_STAUS_DONE = "Done"
)

type Sprint struct {
	ID          gocql.UUID `cql:"id" key:"primary"`
	ProjectId   gocql.UUID `cql:"project_id"`
	ProjectName string     `cql:"project_name"`
	BoardId     gocql.UUID `cql:"board_id"`
	BoardName   string     `cql:"board_name"`
	Goal        string     `cql:"goal"`
	Desc        string     `cql:"description"`
	Status      string     `cql:"status"`
	CreatedAt   time.Time  `cql:"created_at"`
	UpdatedAt   time.Time  `cql:"updated_at"`
}

func (s *Sprint) Insert() error {
	err := Session.Query(`INSERT INTO sprints (id, project_id, project_name, board_id, board_name, goal, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		s.ID, s.ProjectId, s.ProjectName, s.BoardId, s.BoardName, s.Goal, s.Desc, s.CreatedAt, s.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Error in method Insert inside models/sprint.go: %q\n", err.Error())
		return err
	}

	return nil
}

func (s *Sprint) Update() error {
	err := Session.Query(`UPDATE sprints SET goal = ?, description = ?, status = ?, updated_at = ? WHERE id = ?;`,
		s.Goal, s.Desc, s.Status, s.UpdatedAt, s.ID).Exec()

	if err != nil {
		log.Printf("Error in method Update inside models/sprint.go: %q\n", err.Error())
		return err
	}

	return nil
}

func (s *Sprint) Delete() error {
	err := Session.Query(`DELETE FROM sprints WHERE id = ?;`, s.ID).Exec()

	if err != nil {
		log.Printf("Error in method Delete inside models/sprint.go: %q\n", err.Error())
		return err
	}

	return nil
}

func (s *Sprint) FindById() error {
	err := Session.Query(`SELECT id, board_id, goal, description, status, created_at, updated_at FROM sprints WHERE id = ?;`, s.ID).
		Consistency(gocql.One).Scan(&s.ID, &s.BoardId, &s.Goal, &s.Desc, &s.Status, &s.CreatedAt, &s.UpdatedAt)

	if err != nil {
		log.Printf("Error in method FindById inside models/sprint.go: %q\n", err.Error())
		return err
	}

	return nil
}

func (s *Sprint) List(boardId gocql.UUID) ([]map[string]interface{}, error) {

	sprintsList, err := Session.Query(`SELECT * FROM sprintslist WHERE board_id = ?;`, boardId).Iter().SliceMap()

	if err != nil {
		log.Printf("Error in method List inside models/sprint.go: %q\n", err.Error())
		return nil, err
	}

	return sprintsList, nil
}
