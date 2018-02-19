package models

import (
	"time"
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"log"
)

const INSERT_PROJECT 	= "INSERT INTO projects (id,name,created_at,updated_at) VALUES (?,?,?,?);"
const UPDATE_PROJECT 	= "Update projects SET name = ? ,updated_at = ? WHERE id= ? ;"
const DELETE_PROJECT 	= "DELETE FROM projects WHERE id= ? ;"
const FIND_PROJECT 		= "SELECT id, name, created_at, updated_at FROM projects WHERE id = ? LIMIT 1"
const GET_PROJECTS 		= "SELECT id,name,created_at,updated_at from projects"

//Project type
type Project struct {
	UUID      gocql.UUID
	Name 	  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Insert func inserts project obj into table
func (project *Project) Insert() error {

	err := db.GetInstance().Session.Query(INSERT_PROJECT,gocql.TimeUUID(),  project.Name,  project.CreatedAt, project.UpdatedAt).Exec();

	if err != nil {
		log.Printf("Error in method Insert models/project.go: %s\n", err.Error())
		return err
	}

	return nil

}

//Update func updates name of the project by id
func (project *Project) Update() error {

	err := db.GetInstance().Session.Query(UPDATE_PROJECT, project.Name, project.UpdatedAt , project.UUID).Exec()

	if err != nil {
		log.Printf("Error in method Update models/project.go: %s\n", err.Error())
		return err
	}

	return nil

}

//Delete func deletes project by id
func (project *Project) Delete() error {

	err := db.GetInstance().Session.Query(DELETE_PROJECT , project.UUID).Exec()

	if err != nil {
		log.Printf("Error in method Delete models/project.go: %s\n", err.Error())
		return err
	}

	return nil

}

//FindByID func finds project by id
func (project *Project) FindByID() error {

	err := db.GetInstance().Session.Query(FIND_PROJECT,project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		log.Printf("Error in method FindByID models/project.go: %s\n", err.Error())
		return err
	}

	return nil
}

func (project *Project) GetProjectList() ([]Project, error) {
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