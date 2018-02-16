package models

import (
	"time"
	"github.com/gocql/gocql"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)

//Project type
type Project struct {
	UUID      gocql.UUID
	Name 	  string
	CreatedAt time.Time
	UpdatedAt time.Time

}

const INSERT_PROJECT = "INSERT INTO projects (id,name,created_at,updated_at) VALUES (?,?,?,?);"
const UPDATE_PROJECT = "Update projects SET name = ? ,updated_at = ? WHERE id= ? ;"
const DELETE_PROJECT = "DELETE FROM projects WHERE id= ? ;"
const FIND_PROJECT = "SELECT id, name, created_at, updated_at FROM projects WHERE id = ? LIMIT 1"
const GET_PROJECTS = "SELECT id,name,created_at,updated_at from projects"



//Insert func inserts project obj into table
func (project *Project) Insert() {

	if err := db.GetInstance().Session.Query(INSERT_PROJECT,project.UUID, project.Name,  project.CreatedAt, project.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Update func updates name of the project by id
func (project *Project) Update() {

	if err := db.GetInstance().Session.Query(UPDATE_PROJECT,
		 project.Name, project.UpdatedAt , project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Delete func deletes project by id
func (project *Project) Delete() {

	if err := db.GetInstance().Session.Query(DELETE_PROJECT,	project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

//FindByID func finds project by id
func (project *Project) FindByID() error {
	return db.GetInstance().Session.Query(FIND_PROJECT,project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.CreatedAt, &project.UpdatedAt)
}

func (project *Project) GetAll() ([]map[string]interface{}, error) {

	return db.GetInstance().Session.Query(GET_PROJECTS).PageState(nil).PageSize(4).Iter().SliceMap()

}
