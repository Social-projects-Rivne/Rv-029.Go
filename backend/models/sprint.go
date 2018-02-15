package models

import (
	"github.com/gocql/gocql"
	"time"
)

type Sprint struct {
	ID        gocql.UUID `cql:"id" key:"primary"`
	BoardId   gocql.UUID `cql:"board_id"`
	Goal      string     `cql:"goal"`
	Name      string     `cql:"name"`
	Desc      string     `cql:"description"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}
