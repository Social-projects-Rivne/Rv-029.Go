package models

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

// projects queries
const (
	INSERT_PROJECT        = "INSERT INTO projects (id,name,created_at,updated_at) VALUES (?,?,?,?);"
	UPDATE_PROJECT        = "Update projects SET name = ? ,updated_at = ? WHERE id= ? ;"
	DELETE_PROJECT        = "DELETE FROM projects WHERE id= ? ;"
	FIND_PROJECT          = "SELECT id, name, created_at, updated_at FROM projects WHERE id = ? LIMIT 1"
	GET_PROJECTS          = "SELECT id, name, created_at, updated_at FROM projects"
	GET_PROJECTS_WHERE_IN = "SELECT id, name FROM projects WHERE id IN ("
)

type ProjSortName []ProjectName

func (s ProjSortName) Len() int {
	return len(s)
}
func (s ProjSortName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ProjSortName) Less(i, j int) bool {
	var si string = s[i].Name
	var sj string = s[j].Name
	var si_lower = strings.ToLower(si)
	var sj_lower = strings.ToLower(sj)
	if si_lower == sj_lower {
		return si < sj
	}
	return si_lower < sj_lower
}

//ProjectName is struct for getting project's names
type ProjectName struct {
	Name string
	ID gocql.UUID
}

//Project type
type Project struct {
	UUID      gocql.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//go:generate mockgen -destination=../mocks/mock_project.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/models ProjectCRUD

type ProjectCRUD interface {
	Insert(*Project) error
	Update(*Project) error
	Delete(*Project) error
	FindByID(*Project) error
	GetProjectList(*Project) ([]Project, error)
	GetProjectsNamesList([]gocql.UUID) ([]ProjectName, error)
}

type ProjectStorage struct {
	DB *gocql.Session
}

var ProjectDB ProjectCRUD

func InitProjectDB(crud ProjectCRUD) {
	ProjectDB = crud
}

//Insert func inserts project obj into table
func (p *ProjectStorage) Insert(project *Project) error {

	err := Session.Query(INSERT_PROJECT, project.UUID, project.Name, project.CreatedAt, project.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Error in models/project.go error: %+v", err)
		return err
	}

	return nil

}

//Update func updates name of the project by id
func (p *ProjectStorage) Update(project *Project) error {

	err := Session.Query(UPDATE_PROJECT, project.Name, project.UpdatedAt, project.UUID).Exec()

	if err != nil {
		log.Printf("Error in models/project.go error: %+v", err)
		return err
	}

	return nil

}

//Delete func deletes project by id
func (p *ProjectStorage) Delete(project *Project) error {

	err := Session.Query(DELETE_PROJECT, project.UUID).Exec()

	if err != nil {
		log.Printf("Error in models/project.go error: %+v", err)
		return err
	}

	return nil

}

//FindByID func finds project by id
func (p *ProjectStorage) FindByID(project *Project) error {

	err := Session.Query(FIND_PROJECT, project.UUID).Consistency(gocql.One).Scan(&project.UUID, &project.Name, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		log.Printf("Error in models/project.go error: %+v", err)
		return err
	}

	return nil
}

func (p *ProjectStorage) GetProjectList(project *Project) ([]Project, error) {
	var projects []Project
	var row map[string]interface{}
	var pageState []byte

	iterator := p.DB.Query(GET_PROJECTS).Consistency(gocql.One).PageState(pageState).PageSize(5).Iter()
	pageState = iterator.PageState()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			projects = append(projects, Project{
				UUID:      row["id"].(gocql.UUID),
				Name:      row["name"].(string),
				CreatedAt: row["created_at"].(time.Time),
				UpdatedAt: row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in models/project.go error: %+v", err)
	}

	return projects, nil
}

//GetProjectsNamesList returns project's names by project id's
func (p *ProjectStorage) GetProjectsNamesList(list []gocql.UUID) ([]ProjectName, error) {

	tail := ""
	for i := 0; i < len(list); i++ {
		if i == len(list)-1 {

			tail += fmt.Sprintf("%v);", list[i])

		} else {

			tail += fmt.Sprintf("%v,", list[i])

		}
	}

	var projects ProjSortName
	var row map[string]interface{}
	query := fmt.Sprintln(GET_PROJECTS_WHERE_IN + tail)
	iterator := p.DB.Query(query).Consistency(gocql.One).Iter()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			projects = append(projects, ProjectName{
				Name: row["name"].(string),
				ID	: row["id"].(gocql.UUID),
			})
		}
	}
	//listNames := ProjSortName{projects}

	sort.Sort(projects)
	return projects, nil
}
