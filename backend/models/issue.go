package models

import (
	"log"
	"time"
	//"fmt"
	//"log"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
)

//STATUS_TODO uses when issue in TODO list
const STATUS_TODO = "TODO"

//STATUS_IN_PROGRESS uses when issue in progress
const STATUS_IN_PROGRESS = "In_Progress"

//STATUS_ON_HOLD uses when issue on hold
const STATUS_ON_HOLD = "On_Hold"

//STATUS_ON_REVIEW uses when issue on review
const STATUS_ON_REVIEW = "On_Review"

//STATUS_DONE uses when issue done
const STATUS_DONE = "Done"

//Issue model
type Issue struct {
	UUID          gocql.UUID
	Name          string
	Status        string
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

	if err := db.GetInstance().Session.Query(`INSERT INTO issues (id,name,status,user_id,
		user_first_name,user_last_name,sprint_id,board_id,board_name,project_id,
		project_name,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?);`,

		issue.UUID, issue.Name, issue.Status, issue.UserID, issue.UserFirstName, issue.UserLastName,
		issue.SprintID, issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName,
		issue.CreatedAt, issue.UpdatedAt).Exec(); err != nil {

		log.Printf("Error occured inside models/issue.go, method:Insert, error: %v", err)
		return err
	}
	return nil
}

//Update updates issue by UUID
func (issue *Issue) Update() error {

	if err := db.GetInstance().Session.Query(`Update issues SET name = ?, status= ?, user_id = ?,
		user_first_name = ?,user_last_name = ?,sprint_id = ?, board_id = ?,board_name = ?,
		  project_id = ?, project_name = ?,updated_at = ? WHERE id= ? ;`,

		issue.Name, issue.Status, issue.UserID, issue.UserFirstName, issue.UserLastName, issue.SprintID,
		issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName, issue.UpdatedAt, issue.UUID).Exec(); err != nil {

		log.Printf("Error occured inside models/issue.go, method: Update, error: %v", err)
		return err
	}
	return nil
}

//Delete removes issue by id
func (issue *Issue) Delete() error {

	if err := db.GetInstance().Session.Query(`DELETE FROM issues WHERE id= ? ;`,
		issue.UUID).Exec(); err != nil {
		log.Printf("Error occured inside models/issue.go, method: Delete, error: %v", err)
		return err
	}
	return nil
}

//FindByID finds issue by id
func (issue *Issue) FindByID() error {

	if err := db.GetInstance().Session.Query(`SELECT id, name, status, user_id,
		user_first_name, user_last_name, sprint_id, board_id, board_name, project_id,
		project_name, created_at, updated_at
		FROM issues WHERE id = ? LIMIT 1`,

		issue.UUID).Consistency(gocql.One).Scan(&issue.UUID, &issue.Name, &issue.Status, &issue.UserID,
		&issue.UserFirstName, &issue.UserLastName, &issue.SprintID, &issue.BoardID, &issue.BoardName,
		 &issue.ProjectID, &issue.ProjectName, &issue.CreatedAt); err != nil {

		log.Printf("Error occured inside models/issue.go, method:FindByID, error: %v", err)
		return err
	}
	return nil
}

//GetBoardIssueList returns all issues by board_id
func (issue *Issue) GetBoardIssueList() ([]map[string]interface{}, error) {

	issueList, err := Session.Query("SELECT SELECT id, name, status, user_id,user_first_name, user_last_name, sprint_id, board_id, board_name, project_id,project_name, created_at, updated_at from issues WHERE board_id = ?", issue.BoardID).Iter().SliceMap()

	if err != nil {
		log.Printf("Error in method GetBoardIssueList inside models/issue.go, method:GetBoardIssueList, error: %s\n", err.Error())
		return nil, err
	}

	return issueList, nil

}

//GetSprintIssueList returns all issues by board_id
func (issue *Issue) GetSprintIssueList() ([]map[string]interface{}, error) {

	issueList, err := Session.Query("SELECT SELECT id, name, status, user_id,user_first_name, user_last_name, sprint_id, board_id, board_name, project_id,project_name, created_at, updated_at from issues WHERE sprint_id = ?", issue.SprintID).Iter().SliceMap()

	if err != nil {
		log.Printf("Error in method GetSprintIssueList inside models/issue.go: %s\n", err.Error())
		return nil, err
	}

	return issueList, nil

}

// //GetClaims Return list of claims to generate jwt token
// func (user *User) GetClaims() map[string]interface{} {
// 	claims := make(map[string]interface{})

// 	claims["UUID"] = user.UUID

// 	return claims
// }
