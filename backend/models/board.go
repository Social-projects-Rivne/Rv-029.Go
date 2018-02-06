package models

import (
	"github.com/gocql/gocql"
	"time"
)

type Board struct {
	UUID      gocql.UUID `cql:"id" key:"primary"`
	ProjectID gocql.UUID `cql:"project_id"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}


