package models

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
	"log"
	"time"
)

type Board struct {
	ID        gocql.UUID `cql:"id" key:"primary"`
	ProjectID gocql.UUID `cql:"project_id"`
	Name      string     `cql:"name"`
	Desc      string     `cql:"description"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}

var Session = db.GetInstance().Session

func (b *Board) Insert() error {
	err := Session.Query(`INSERT INTO boards (id, project_id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`,
		b.ID, b.ProjectID, b.Name, b.Desc, b.CreatedAt, b.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Invalid Insert inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

func (b *Board) Update() error {
	err := Session.Query(`UPDATE boards SET name = ?, description = ?, updated_at = ? WHERE id = ?;`,
		b.Name, b.Desc, b.UpdatedAt, b.ID).Exec()

	if err != nil {
		log.Printf("Invalid Update inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

func (b *Board) Delete() error {
	err := Session.Query(`DELETE FROM boards where id = ?;`, b.ID).Exec()

    if err != nil {
		log.Printf("Invalid Delete inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

func (b *Board) FindByID() error {
	err := Session.Query(`SELECT id, project_id, name, description, created_at, updated_at FROM boards WHERE id = ? LIMIT 1`,
		b.ID).Consistency(gocql.One).Scan(&b.ID, &b.ProjectID, &b.Name, &b.Desc, &b.CreatedAt, &b.UpdatedAt)

	if err != nil {
		log.Printf("Invalid FindById inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

func (b *Board) List() ([]map[string]interface{}, error) {
	boardsList, err := Session.Query(`SELECT * from boards`).Iter().SliceMap()

	if err != nil {
		log.Printf("Invalid List inside models/board.go: %s\n", err.Error())
		return nil, err
	}

	return boardsList, nil
}
