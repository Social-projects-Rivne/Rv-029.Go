package models

import (
	"github.com/gocql/gocql"
	"log"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
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

func (b *Board) Insert() {
	if err := Session.Query(`INSERT INTO boards (id, project_id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`,
		b.ID, b.ProjectID, b.Name, b.Desc, b.CreatedAt, b.UpdatedAt).Exec(); err != nil {
			log.Fatal(err)
		}
}

func (b *Board) Update() {
	if err := Session.Query(`Update boards SET name = ?, description = ?, updated_at = ? WHERE id = ?;`,
		b.Name, b.Desc, b.UpdatedAt, b.ID).Exec(); err != nil {
			log.Fatal(err)
		}
}
