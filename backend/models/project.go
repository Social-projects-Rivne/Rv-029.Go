package models

import (
	"time"
	"github.com/gocql/gocql"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)

type Project struct {
	UUID      gocql.UUID
	UserId    gocql.UUID
	Name 	  string
	CreatedAt time.Time
	UpdatedAt time.Time

}
var Session = db.GetInstance().Session

const INSERT_PROJECT = "INSERT INTO projects (id,user_id,name,created_at,updated_at) VALUES (?,?,?,?,?);"
const UPDATE_PROJECT = "Update projects SET name = ? ,updated_at = ? WHERE id= ? ;"
const DELETE_PROJECT = "DELETE FROM projects WHERE id= ? ;"
const FIND_PROJECT = "SELECT id, name, user_id, created_at, updated_at FROM projects WHERE id = ? LIMIT 1"
const GET_PROJECTS = "SELECT id,name,created_at,updated_at from projects"



func (project *Project) Insert() {

	if err := Session.Query(INSERT_PROJECT,gocql.TimeUUID(), project.UserId, project.Name,  project.CreatedAt, project.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) Update() {

	if err := Session.Query(UPDATE_PROJECT,
		 project.Name, project.UpdatedAt , project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) Delete() {

	if err := Session.Query(DELETE_PROJECT,	project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) FindByID() error {
	return Session.Query(FIND_PROJECT,project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.UserId,&project.CreatedAt, &project.UpdatedAt)
}

func (project *Project) GetAll() ([]map[string]interface{}, error){

	return Session.Query(GET_PROJECTS).PageState(nil).PageSize(2).Iter().SliceMap()

}
