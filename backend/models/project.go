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
	UserId    gocql.UUID
	Name 	  string
	CreatedAt time.Time
	UpdatedAt time.Time

}

//Session const variable
var Session = db.GetInstance().Session

//Insert func inserts project obj into table
func (project *Project) Insert() {

	if err := Session.Query(`INSERT INTO projects (id,user_id,name,created_at,updated_at) VALUES (?,?,?,?,?);	`,
		gocql.TimeUUID(), project.UserId, project.Name,  project.CreatedAt, project.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Update func updates name of the project by id
func (project *Project) Update() {

	if err := Session.Query(`Update projects SET name = ? ,updated_at = ? WHERE id= ? ;`,
		 project.Name, project.UpdatedAt , project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Delete func deletes project by id
func (project *Project) Delete() {

	if err := Session.Query(`DELETE FROM projects WHERE id= ? ;`,
		project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

//FindByID func finds project by id
func (project *Project) FindByID() error {
	return Session.Query(`SELECT id, name, user_id, created_at, updated_at FROM projects WHERE id = ? LIMIT 1`,
		project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.UserId,&project.CreatedAt, &project.UpdatedAt)
}

//GetAll func gets all projects from 
func (project *Project) GetAll() ([]map[string]interface{}, error){

	return Session.Query(`SELECT * from projects`).Iter().SliceMap()

}
