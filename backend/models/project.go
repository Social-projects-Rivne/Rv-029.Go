package models

import (
	"time"
	"github.com/gocql/gocql"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"log"
)

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



func (project *Project) Insert() {

	if err := db.GetInstance().Session.Query(INSERT_PROJECT,project.UUID, project.Name,  project.CreatedAt, project.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) Update() {

	if err := db.GetInstance().Session.Query(UPDATE_PROJECT,
		 project.Name, project.UpdatedAt , project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) Delete() {

	if err := db.GetInstance().Session.Query(DELETE_PROJECT,	project.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

func (project *Project) FindByID() error {
	return db.GetInstance().Session.Query(FIND_PROJECT,project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.CreatedAt, &project.UpdatedAt)
}

func (project *Project) GetAll() ([]Project, error) {
	var projects []Project
	var row map[string]interface{}

	iterator := db.GetInstance().Session.Query(GET_PROJECTS).Consistency(gocql.One).Iter()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			projects = append(projects, Project{
				UUID: 		row["id"].(gocql.UUID),
				Name: 		row["name"].(string),
				CreatedAt: 	row["created_at"].(time.Time),
				UpdatedAt: 	row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Can`t fetch all projects from DB. Error: %s", err.Error())
	}

	return projects, nil
}
