package models

import (
	"log"
	"time"
	//"log"

	"github.com/gocql/gocql"
)
const(
	STATUS_TODO = "TODO"
	STATUS_IN_PROGRESS = "In Progress"
	STATUS_ON_HOLD = "On Hold"
	STATUS_ON_REVIEW = "On Review"
	STATUS_DONE = "Done"
	INSERT_iSSUE = "INSERT INTO issues (id,name,status,description,estimate,user_id,user_first_name,user_last_name,sprint_id,board_id,board_name,project_id,project_name,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	UPDATE_ISSUE = "Update issues SET name = ?, status = ?, description = ?,estimate = ?, user_id = ?,user_first_name = ?,user_last_name = ?,sprint_id = ?, board_id = ?,board_name = ?,project_id = ?, project_name = ?, updated_at = ? WHERE id= ? ;"
	DELETE_ISSUE = "DELETE FROM issues WHERE id= ? ;"
	FIND_ISSUE_BY_ID = "SELECT id, name, status, description,estimate, user_id,user_first_name, user_last_name,sprint_id, board_id, board_name, project_id,project_name, created_at, updated_at FROM issues WHERE id = ? LIMIT 1"
	GET_BOARD_ISSUE_LIST = "SELECT id, name, status, description, estimate, user_id,user_first_name,user_last_name, sprint_id, board_id, board_name, project_id,project_name,created_at, updated_at from issues WHERE board_id = ? AND sprint_id = 00000000-0000-0000-0000-000000000000 ALLOW FILTERING"
	GET_SPRINT_ISSUE_LIST = "SELECT id, name, status, description, estimate, user_id,user_first_name,user_last_name, sprint_id, board_id, board_name, project_id, project_name,created_at, updated_at from issues WHERE sprint_id = ? ALLOW FILTERING"
)

//go:generate mockgen -destination=../mocks/mock_board.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/models BoardCRUD

type IssueCRUD interface {
	Insert(*Issue) error
	Update(*Issue) error
	Delete(*Issue) error
	FindByID(*Issue) error
	GetBoardIssueList(gocql.UUID) ([]map[string]interface{}, error)
	GetSprintIssueList(gocql.UUID) ([]map[string]interface{}, error)	
}

type IssueStorage struct {
	DB *gocql.Session
}

var IssueDB IssueCRUD

func InitIssueDB(crud IssueCRUD) {
	IssueDB = crud
}


//Issue model
type Issue struct {
	UUID          gocql.UUID
	Name          string
	Status        string
	Description   string
	Estimate      int
	UserID        gocql.UUID
	UserFirstName string
	UserLastName  string
	SprintID      gocql.UUID
	BoardID       gocql.UUID
	BoardName     string
	ProjectID     gocql.UUID
	ProjectName   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//Insert func inserts user object in database
func (issue *Issue) Insert() error {

	if err := Session.Query(INSERT_iSSUE,

		issue.UUID, issue.Name, issue.Status, issue.Description, issue.Estimate, issue.UserID, issue.UserFirstName, issue.UserLastName,
		issue.SprintID, issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName,
		issue.CreatedAt, issue.UpdatedAt).Exec(); err != nil {

			log.Printf("Error in models/issue.go error: %+v",err)
			return err
	}
	return nil
}

//Update updates issue by UUID
func (issue *Issue) Update() error {

	if err := Session.Query(UPDATE_ISSUE,

		issue.Name, issue.Status, issue.Description, issue.Estimate, issue.UserID, issue.UserFirstName, issue.UserLastName, issue.SprintID,
		issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName, issue.UpdatedAt, issue.UUID).Exec(); err != nil {

			log.Printf("Error in models/issue.go error: %+v",err)
		return err
	}
	return nil
}

//Delete removes issue by id
func (issue *Issue) Delete() error {

	if err := Session.Query(DELETE_ISSUE,
		issue.UUID).Exec(); err != nil {
			log.Printf("Error in models/issue.go error: %+v",err)
		return err
	}
	return nil
}

//FindByID finds issue by id
func (issue *Issue) FindByID() error {

	if err := Session.Query(FIND_ISSUE_BY_ID, issue.UUID).Consistency(gocql.One).Scan(&issue.UUID, &issue.Name, &issue.Status, &issue.Description, &issue.Estimate, &issue.UserID,
		&issue.UserFirstName, &issue.UserLastName, &issue.SprintID, &issue.BoardID, &issue.BoardName,
		&issue.ProjectID, &issue.ProjectName, &issue.CreatedAt, &issue.UpdatedAt); err != nil {

			log.Printf("Error in models/issue.go error: %+v",err)
		return err
	}
	return nil
}

//GetBoardIssueList returns all issues by board_id
func GetBoardIssueList(BoardID gocql.UUID) ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := Session.Query(GET_BOARD_ISSUE_LIST, BoardID).Iter()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			issues = append(issues, Issue{
				UUID: row["id"].(gocql.UUID),
				Name: row["name"].(string),
				Status: row["status"].(string),
				Description: row["description"].(string),
				Estimate: row["estimate"].(int),
				UserID: row["user_id"].(gocql.UUID),
				UserFirstName: row["user_first_name"].(string),
				UserLastName: row["user_last_name"].(string),
				SprintID: row["sprint_id"].(gocql.UUID),
				BoardID: row["board_id"].(gocql.UUID),
				BoardName: row["board_name"].(string),
				ProjectID: row["project_id"].(gocql.UUID),
				ProjectName: row["project_name"].(string),
				CreatedAt: row["created_at"].(time.Time),
				UpdatedAt: row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in method GetBoardIssueList inside models/issue.go, method:GetBoardIssueList, error: %s\n", err.Error())
		return nil, err
	}

	return issues, nil
}

//GetSprintIssueList returns all issues by board_id
func GetSprintIssueList(SprintID gocql.UUID) ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := Session.Query(GET_SPRINT_ISSUE_LIST, SprintID).Iter()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			issues = append(issues, Issue{
				UUID: row["id"].(gocql.UUID),
				Name: row["name"].(string),
				Status: row["status"].(string),
				Description: row["description"].(string),
				Estimate: row["estimate"].(int),
				UserID: row["user_id"].(gocql.UUID),
				UserFirstName: row["user_first_name"].(string),
				UserLastName: row["user_last_name"].(string),
				SprintID: row["sprint_id"].(gocql.UUID),
				BoardID: row["board_id"].(gocql.UUID),
				BoardName: row["board_name"].(string),
				ProjectID: row["project_id"].(gocql.UUID),
				ProjectName: row["project_name"].(string),
				CreatedAt: row["created_at"].(time.Time),
				UpdatedAt: row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in models/issue.go error: %+v",err)
		return nil, err
	}

	return issues, nil
}