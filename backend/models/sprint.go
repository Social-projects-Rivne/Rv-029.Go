package models

import (
	"github.com/gocql/gocql"
	"log"
	"time"
)

type Sprint struct {
	ID        gocql.UUID `cql:"id" key:"primary"`
	BoardId   gocql.UUID `cql:"board_id"`
	Goal      string     `cql:"goal"`
	Desc      string     `cql:"description"`
	Status    string     `cql:"status"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}

func (s *Sprint) Insert() error {
	err := Session.Query(`INSERT INTO sprints (id, board_id, goal, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`,
		s.ID, s.BoardId, s.Goal, s.Desc, s.CreatedAt, s.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Error in method Insert inside models/sprint.go: %q\n", err.Error())
		return err
	}

	return nil
}

func (s *Sprint) Update() error {
	err := Session.Query(`UPDATE sprints SET goal = ?, description = ?, status = ?, updated_at = ? WHERE id = ?;`,
		s.Goal, s.Desc, s.Status, s.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Error in method Update inside models/sprint.go: %q\n", err.Error())
		return err
	}

	return nil
}
