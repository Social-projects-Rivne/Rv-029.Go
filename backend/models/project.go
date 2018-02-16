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

const INSERT_PROJECT = "INSERT INTO projects (id,user_id,name,created_at,updated_at) VALUES (?,?,?,?,?);"
const UPDATE_PROJECT = "Update projects SET name = ? ,updated_at = ? WHERE id= ? ;"
const DELETE_PROJECT = "DELETE FROM projects WHERE id= ? ;"
const FIND_PROJECT = "SELECT id, name, user_id, created_at, updated_at FROM projects WHERE id = ? LIMIT 1"
const GET_PROJECTS = "SELECT id,name,created_at,updated_at from projects"
const GET_USER_PROJECTS = "SELECT id, name, created_at, updated_at from user_projects_view WHERE user_id = ?"



//Insert func inserts project obj into table
func (project *Project) Insert() {

	if err := db.GetInstance().Session.Query(INSERT_PROJECT,gocql.TimeUUID(), project.UserId, project.Name,  project.CreatedAt, project.UpdatedAt).Exec(); err != nil {
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
	return db.GetInstance().Session.Query(FIND_PROJECT,project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.UserId,&project.CreatedAt, &project.UpdatedAt)
}

func (project *Project) GetAll() ([]map[string]interface{}, error) {

	return db.GetInstance().Session.Query(GET_PROJECTS).PageState(nil).PageSize(2).Iter().SliceMap()

}

func (Project) FindUserProjects (user *User) ([]Project, error) {
	row := map[string]interface{}{}
	var projects []Project

	iterator := db.GetInstance().Session.Query(GET_USER_PROJECTS).Consistency(gocql.One).Iter()
	if iterator.NumRows() > 0 {
		for iterator.MapScan(row) {
			projects = append(projects, Project{
				UUID: 			row["id"].(gocql.UUID),
				Name: 			row["name"].(string),
				CreatedAt: 		row["created_at"].(time.Time),
				UpdatedAt: 		row["updated_at"].(time.Time),
			})
		}
	}

	return projects, nil
}