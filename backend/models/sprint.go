package models

import (
	"github.com/gocql/gocql"
	"log"
	"time"
)

const (
	SPRINT_STAUS_TODO        = "TODO"
	SPRINT_STAUS_IN_PROGRESS = "In Progress"
	SPRINT_STAUS_DONE        = "Done"
)

const (
	PROJECT_SPRINTS_LIST               = `SELECT * FROM sprintslist WHERE board_id = ?;`
	GET_SPRINT_ISSUES_IN_PROGRESS_LIST = `SELECT id,name,status,description,estimate,user_id,user_first_name,user_last_name,sprint_id,board_id,board_name,project_id,project_name,created_at,updated_at FROM sprint_issues WHERE sprint_id = ? AND status IN (?, ?);`
)

type Sprint struct {
	ID          gocql.UUID `cql:"id" key:"primary"`
	ProjectId   gocql.UUID `cql:"project_id"`
	ProjectName string     `cql:"project_name"`
	BoardId     gocql.UUID `cql:"board_id"`
	BoardName   string     `cql:"board_name"`
	Goal        string     `cql:"goal"`
	Desc        string     `cql:"description"`
	Status      string     `cql:"status"`
	CreatedAt   time.Time  `cql:"created_at"`
	UpdatedAt   time.Time  `cql:"updated_at"`
}

//go:generate mockgen -destination=../mocks/mock_sprint.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/models SprintCRUD

type SprintCRUD interface {
	Insert(*Sprint) error
	Update(*Sprint) error
	Delete(*Sprint) error
	FindByID(*Sprint) error
	List(gocql.UUID) ([]map[string]interface{}, error)
	GetSprintIssuesInProgress(*Sprint) ([]Issue, error) // todo Sprint argument
}

type SprintStorage struct {
	DB *gocql.Session
}

var SprintDB SprintCRUD

func InitSprintDB(crud SprintCRUD) {
	SprintDB = crud
}

func (ss *SprintStorage) Insert(s *Sprint) error {
	err := Session.Query(`INSERT INTO sprints (id, project_id, project_name, status, board_id, board_name, goal, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		s.ID, s.ProjectId, s.ProjectName, s.Status, s.BoardId, s.BoardName, s.Goal, s.Desc, s.CreatedAt, s.UpdatedAt).Exec()

	if err != nil {
		log.Printf("Error in models/sprint.go error: %+v",err)
		return err
	}

	return nil
}

func (ss *SprintStorage) Update(s *Sprint) error {
	err := ss.DB.Query(`UPDATE sprints SET goal = ?, description = ?, status = ?, updated_at = ? WHERE id = ?;`,
		s.Goal, s.Desc, s.Status, s.UpdatedAt, s.ID).Exec()

	if err != nil {
		log.Printf("Error in models/sprint.go error: %+v",err)
		return err
	}

	return nil
}

func (ss *SprintStorage) Delete(s *Sprint) error {
	err := ss.DB.Query(`DELETE FROM sprints WHERE id = ?;`, s.ID).Exec()

	if err != nil {
		log.Printf("Error in models/sprint.go error: %+v",err)
		return err
	}

	return nil
}

func (ss *SprintStorage) FindById(s *Sprint) error {
	err := ss.DB.Query(`SELECT id, board_id, goal, description, status, created_at, updated_at FROM sprints WHERE id = ?;`, s.ID).
		Consistency(gocql.One).Scan(&s.ID, &s.BoardId, &s.Goal, &s.Desc, &s.Status, &s.CreatedAt, &s.UpdatedAt)

	if err != nil {
		log.Printf("Error in models/sprint.go error: %+v",err)
		return err
	}

	return nil
}

func (ss *SprintStorage) List(boardId gocql.UUID) ([]Sprint, error) {

	sprints := []Sprint{}
	var row map[string]interface{}

	iterator := ss.DB.Query(PROJECT_SPRINTS_LIST, boardId).Iter()
	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			sprints = append(sprints, Sprint{
				ID:          row["id"].(gocql.UUID),
				ProjectId:   row["project_id"].(gocql.UUID),
				ProjectName: row["project_name"].(string),
				BoardId:     row["board_id"].(gocql.UUID),
				BoardName:   row["board_name"].(string),
				Goal:        row["goal"].(string),
				Desc:        row["description"].(string),
				Status:      row["status"].(string),
				CreatedAt:   row["created_at"].(time.Time),
				UpdatedAt:   row["created_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in method List inside models/sprint.go: %s\n", err.Error())
		return nil, err
	}

	return sprints, nil
}

func (ss *SprintStorage) GetSprintIssuesInProgress(s *Sprint) ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := ss.DB.Query(GET_SPRINT_ISSUES_IN_PROGRESS_LIST, s.ID, SPRINT_STAUS_TODO, SPRINT_STAUS_IN_PROGRESS).Iter()
	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			issues = append(issues, Issue{
				UUID:          row["id"].(gocql.UUID),
				Name:          row["name"].(string),
				Status:        row["status"].(string),
				Description:   row["description"].(string),
				Estimate:      row["estimate"].(int),
				UserID:        row["user_id"].(gocql.UUID),
				UserFirstName: row["user_first_name"].(string),
				UserLastName:  row["user_last_name"].(string),
				SprintID:      row["sprint_id"].(gocql.UUID),
				BoardID:       row["board_id"].(gocql.UUID),
				BoardName:     row["board_name"].(string),
				ProjectID:     row["project_id"].(gocql.UUID),
				ProjectName:   row["project_name"].(string),
				CreatedAt:     row["created_at"].(time.Time),
				UpdatedAt:     row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in method List inside models/sprint.go: %s\n", err.Error())
		return nil, err
	}

	return issues, nil
}
