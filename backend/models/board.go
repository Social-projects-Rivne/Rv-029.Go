package models

import (
	"github.com/gocql/gocql"
	"log"
	"time"
)

//Board model
type Board struct {
	ID        gocql.UUID `cql:"id" key:"primary"`
	ProjectID gocql.UUID `cql:"project_id"`
	Name      string     `cql:"name"`
	Desc      string     `cql:"description"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}

//Insert func inserts board object in database
func (b *Board) Insert() error {
	err := Session.Query(`INSERT INTO boards (id, project_id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`,
		b.ID, b.ProjectID, b.Name, b.Desc, b.CreatedAt, b.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Error in method Insert inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

//Update func updates board name and description by id
func (b *Board) Update() error {
	err := Session.Query(`UPDATE boards SET name = ?, description = ?, updated_at = ? WHERE id = ?;`,
		b.Name, b.Desc, b.UpdatedAt, b.ID).Exec()

	if err != nil {
		log.Printf("Error in method Update inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

//Delete removes board by id
func (b *Board) Delete() error {
	err := Session.Query(`DELETE FROM boards where id = ?;`, b.ID).Exec()

    if err != nil {
		log.Printf("Error in method Delete inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

//FindByID func finds board by id
func (b *Board) FindByID() error {
	err := Session.Query(`SELECT id, project_id, name, description, created_at, updated_at FROM boards WHERE id = ? LIMIT 1`,
		b.ID).Consistency(gocql.One).Scan(&b.ID, &b.ProjectID, &b.Name, &b.Desc, &b.CreatedAt, &b.UpdatedAt)

	if err != nil {
		log.Printf("Error in method inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

//List func return list of boards orger by project_id
func (b *Board) List(projectId gocql.UUID) ([]map[string]interface{}, error) {

	boardsList, err := Session.Query(`SELECT * FROM boardslist WHERE project_id = ?;`, projectId).Iter().SliceMap()

	if err != nil {
		log.Printf("Error in method List inside models/board.go: %s\n", err.Error())
		return nil, err
	}

	return boardsList, nil
}
