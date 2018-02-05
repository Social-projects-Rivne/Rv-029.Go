package models

import (
	"time"
	"github.com/gocql/gocql"
)





type Project struct {
	UUID      gocql.UUID	`cql:"id" key:"primery"` 
	UserId    gocql.UUID	`cql:"user_id"`
	Name 	  string		`cql:"name"`
	CreatedAt time.Time		`cql:"created_at"`
	UpdatedAt time.Time		`cql:"updated_at"`

}

