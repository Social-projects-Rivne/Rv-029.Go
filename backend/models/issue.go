package models

import (
	"time"
	//"fmt"
	//"log"

	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
)

//Uses when issue in backlog
const STATUS_BACKLOG = "Backlog"
//Uses when issue in sprint backlog
const STATUS_SPRINT_BACKLOG = "Sprint_Backlog"
//Uses when issue in progress
const STATUS_IN_PROGRESS = "In_Progress"
//Uses when issue on hold
const STATUS_ON_HOLD = "On_Hold"
//Uses when issue on review
const STATUS_ON_REVIEW = "On_Review"
//Uses when issue done
const STATUS_DONE = "Done"

//Issue model
type Issue struct {
	UUID      gocql.UUID `cql:"id" key:"primery"`
	Name      string     `cql:"name"`
	Status    string     `cql:"status"`
	UserID    gocql.UUID `cql:"user_id"`
	ProjectID gocql.UUID `cql:"project_id"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}
