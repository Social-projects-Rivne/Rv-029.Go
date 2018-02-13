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
func (b *Board) Insert() {
	if err := Session.Query(`INSERT INTO boards (id, project_id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`,
		b.ID, b.ProjectID, b.Name, b.Desc, b.CreatedAt, b.UpdatedAt).Exec(); err != nil {
		log.Fatal(err)
	}
}

//Update func updates board name and description by id
func (b *Board) Update() {
	if err := Session.Query(`UPDATE boards SET name = ?, description = ?, updated_at = ? WHERE id = ?;`,
		b.Name, b.Desc, b.UpdatedAt, b.ID).Exec(); err != nil {
		log.Fatal(err)
	}
}

//Delete removes board by id
func (b *Board) Delete() {
	if err := Session.Query(`DELETE FROM boards where id = ?;`, b.ID).Exec(); err != nil {
		log.Fatal(err)
	}
}

//FindByID func finds board by id
func (b *Board) FindByID() {
	if err := Session.Query(`SELECT id, project_id, name, description, created_at, updated_at FROM boards WHERE id = ? LIMIT 1`,
		b.ID).Consistency(gocql.One).Scan(&b.ID, &b.ProjectID, &b.Name, &b.Desc, &b.CreatedAt, &b.UpdatedAt); err != nil {
		log.Fatal(err)
	}
}

//List func returns all boards in list
func (b *Board) List() []map[string]interface{} {
	boardsList, err := Session.Query(`SELECT * from boards`).Iter().SliceMap()

	if err != nil {
		log.Fatal(err)
	}

	return boardsList
}
