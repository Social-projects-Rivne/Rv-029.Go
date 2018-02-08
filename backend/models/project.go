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

func (project *Project) Insert() {

	if err := Session.Query(`INSERT INTO projects (id,user_id,name,created_at,updated_at) VALUES (?,?,?,?,?);	`,
		gocql.TimeUUID(), project.UserId, project.Name,  project.CreatedAt, project.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) Update() {

	if err := Session.Query(`Update projects SET name = ? ,updated_at = ? WHERE id= ? ;`,
		 project.Name, project.UpdatedAt , project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) Delete() {

	if err := Session.Query(`DELETE FROM projects WHERE id= ? ;`,
		project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) FindByID() error {
	return Session.Query(`SELECT id, name, user_id, created_at, updated_at FROM projects WHERE id = ? LIMIT 1`,
		project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.UserId,&project.CreatedAt, &project.UpdatedAt)
}

func (project *Project) GetAll() ([]map[string]interface{}, error){

	return Session.Query(`SELECT * from projects`).Iter().SliceMap()

}
