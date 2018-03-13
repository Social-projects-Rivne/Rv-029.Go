package models

import (
	"log"
	"time"
	//"log"

	"github.com/gocql/gocql"
)

//go:generate mockgen -destination=../mocks/mock_issue.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/models IssueCRUD

const (
	//STATUS_TODO uses when issue in TODO list
	STATUS_TODO = "TODO"

	//STATUS_IN_PROGRESS uses when issue in progress
	STATUS_IN_PROGRESS = "In Progress"

	//STATUS_ON_HOLD uses when issue on hold
	STATUS_ON_HOLD = "On Hold"

	//STATUS_ON_REVIEW uses when issue on review1
	STATUS_ON_REVIEW = "On Review"

	//STATUS_DONE uses when issue done
	STATUS_DONE = "Done"

	INSERT_iSSUE = "INSERT INTO issues (id,name,status,description,estimate,user_id,user_first_name,user_last_name,sprint_id,board_id,board_name,project_id,project_name,parent,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

	UPDATE_ISSUE = "Update issues SET name = ?, status = ?, description = ?,estimate = ?, user_id = ?,user_first_name = ?,user_last_name = ?, board_name = ?, project_name = ?, parent = ?, updated_at = ? WHERE id = ? AND board_id = ? AND sprint_id = ? AND project_id = ?;"

	DELETE_ISSUE = "DELETE FROM issues WHERE id = ?;"

	FIND_ISSUE_BY_ID = "SELECT id, name, status, description,estimate, user_id,user_first_name, user_last_name,sprint_id, board_id, board_name, project_id, project_name, parent, created_at, updated_at FROM issues WHERE id = ? LIMIT 1"

	GET_BOARD_ISSUE_LIST          = "SELECT id, name, status, description, estimate, user_id,user_first_name,user_last_name, sprint_id, board_id, board_name, project_id, project_name, parent, created_at, updated_at FROM board_issues WHERE board_id = ?"

	GET_BOARD_BACKLOG_ISSUES_LIST = "SELECT id, name, status, description, estimate, user_id,user_first_name,user_last_name, sprint_id, board_id, board_name, project_id, project_name, parent, created_at, updated_at FROM board_issues WHERE board_id = ? AND sprint_id = 00000000-0000-0000-0000-000000000000"

	GET_SPRINT_ISSUE_LIST = "SELECT id, name, status, description, estimate, user_id,user_first_name,user_last_name, sprint_id, board_id, board_name, project_id, project_name, parent, created_at, updated_at from sprint_issues WHERE sprint_id = ? ;"
)


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
	Parent	      gocql.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type IssueCRUD interface {
	Insert(*Issue) error
	Update(*Issue) error
	Delete(*Issue) error
	FindByID(*Issue) error
	GetBoardIssueList(*Issue) ([]Issue, error)
	GetSprintIssueList(*Issue) ([]Issue, error)
	GetBoardBacklogIssuesList(*Issue) ([]Issue, error)
}

type IssueStorage struct {
	DB *gocql.Session
}

var IssueDB IssueCRUD

func InitIssueDB(crud IssueCRUD) {
	IssueDB = crud
}
//Insert func inserts user object in database
func (s *IssueStorage) Insert(issue *Issue) error {

	if err := s.DB.Query(INSERT_iSSUE,

		issue.UUID, issue.Name, issue.Status, issue.Description, issue.Estimate, issue.UserID, issue.UserFirstName, issue.UserLastName,
		issue.SprintID, issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName, issue.Parent,
		issue.CreatedAt, issue.UpdatedAt).Exec(); err != nil {

		log.Printf("Error in models/issue.go error: %+v", err)
		return err
	}
	return nil
}

//Update updates issue by UUID
func (s *IssueStorage) Update(issue *Issue) error {

	if err := s.DB.Query(UPDATE_ISSUE,

		issue.Name, issue.Status, issue.Description, issue.Estimate, issue.UserID, issue.UserFirstName, issue.UserLastName,
		issue.BoardName, issue.ProjectName, issue.Parent, issue.UpdatedAt, issue.UUID, issue.BoardID, issue.SprintID, issue.ProjectID).Exec(); err != nil {
		log.Printf("Error in models/issue.go error: %+v",err)
		return err
	}

	return nil
}

// FIXME
// bug possible
// see previous commit to figure out changes
func (s *IssueStorage) Delete(issue *Issue) error {

	if err := s.DB.Query(DELETE_ISSUE, issue.UUID).Exec(); err != nil {
		log.Printf("Error occured inside models/issue.go, method: Delete, error: %+v", err)
		return err
	}

	return nil
}

//FindByID finds issue by id
func (s *IssueStorage) FindByID(issue *Issue) error {

	if err := s.DB.Query(FIND_ISSUE_BY_ID, issue.UUID).Consistency(gocql.One).Scan(&issue.UUID, &issue.Name, &issue.Status, &issue.Description, &issue.Estimate, &issue.UserID,
		&issue.UserFirstName, &issue.UserLastName, &issue.SprintID, &issue.BoardID, &issue.BoardName,
		&issue.ProjectID, &issue.ProjectName, &issue.Parent, &issue.CreatedAt, &issue.UpdatedAt); err != nil {

		log.Printf("Error in models/issue.go error: %+v", err)
		return err
	}
	return nil
}

//GetBoardIssueList returns all issues by board_id
func (s *IssueStorage) GetBoardIssueList(issue *Issue) ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := s.DB.Query(GET_BOARD_ISSUE_LIST, issue.BoardID).Iter()

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
				Parent:        row["parent"].(gocql.UUID),
				CreatedAt:     row["created_at"].(time.Time),
				UpdatedAt:     row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in method GetBoardIssueList inside models/issue.go, method:GetBoardIssueList, error: %s\n", err.Error())
		return nil, err
	}

	return issues, nil
}

//GetBoardIssueList returns all issues by board_id what is in backlog
func (s *IssueStorage) GetBoardBacklogIssuesList(issue *Issue) ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := Session.Query(GET_BOARD_BACKLOG_ISSUES_LIST, issue.BoardID).Iter()

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
				Parent:        row["parent"].(gocql.UUID),
				CreatedAt:     row["created_at"].(time.Time),
				UpdatedAt:     row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in method GetBoardIssueList inside models/issue.go, method:GetBoardBacklogIssuesList, error: %s\n", err.Error())
		return nil, err
	}

	return issues, nil
}

//GetSprintIssueList returns all issues by board_id
func (s *IssueStorage) GetSprintIssueList(issue *Issue) ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := s.DB.Query(GET_SPRINT_ISSUE_LIST, issue.SprintID).Iter()

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
				Parent:        row["parent"].(gocql.UUID),
				CreatedAt:     row["created_at"].(time.Time),
				UpdatedAt:     row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in models/issue.go error: %+v", err)
		return nil, err
	}

	return issues, nil
}
