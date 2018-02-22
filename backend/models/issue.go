package models

import (
	"log"
	"time"
	//"log"

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

	if err := Session.Query(`INSERT INTO issues (id,name,status,description,estimate,user_id,
		user_first_name,user_last_name,sprint_id,board_id,board_name,project_id,
		project_name,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`,

		issue.UUID, issue.Name, issue.Status, issue.Description, issue.Estimate, issue.UserID, issue.UserFirstName, issue.UserLastName,
		issue.SprintID, issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName,
		issue.CreatedAt, issue.UpdatedAt).Exec(); err != nil {

		log.Printf("Error occured inside models/issue.go, method:Insert, error: %v", err)
		return err
	}
	return nil
}

//Update updates issue by UUID
func (issue *Issue) Update() error {

	if err := Session.Query(`Update issues SET name = ?, status = ?, description = ?,
		 estimate = ?, user_id = ?,user_first_name = ?,
		 user_last_name = ?,sprint_id = ?, board_id = ?,board_name = ?,
		  project_id = ?, project_name = ?,updated_at = ? WHERE id= ? ;`,

		issue.Name, issue.Status, issue.Description, issue.Estimate, issue.UserID, issue.UserFirstName, issue.UserLastName, issue.SprintID,
		issue.BoardID, issue.BoardName, issue.ProjectID, issue.ProjectName, issue.UpdatedAt, issue.UUID).Exec(); err != nil {

		log.Printf("Error occured inside models/issue.go, method: Update, error: %v", err)
		return err
	}
	return nil
}

//Delete removes issue by id
func (issue *Issue) Delete() error {

	if err := Session.Query(`DELETE FROM issues WHERE id= ? ;`,
		issue.UUID).Exec(); err != nil {
		log.Printf("Error occured inside models/issue.go, method: Delete, error: %v", err)
		return err
	}
	return nil
}

//FindByID finds issue by id
func (issue *Issue) FindByID() error {

	if err := Session.Query(`SELECT id, name, status, description,
		estimate, user_id,user_first_name, user_last_name,
		sprint_id, board_id, board_name, project_id,
		project_name, created_at, updated_at
		FROM issues WHERE id = ? LIMIT 1`,

		issue.UUID).Consistency(gocql.One).Scan(&issue.UUID, &issue.Name, &issue.Status, &issue.Description, &issue.Estimate, &issue.UserID,
		&issue.UserFirstName, &issue.UserLastName, &issue.SprintID, &issue.BoardID, &issue.BoardName,
		&issue.ProjectID, &issue.ProjectName, &issue.CreatedAt, &issue.UpdatedAt); err != nil {

		log.Printf("Error occured inside models/issue.go, method:FindByID, error: %v", err)
		return err
	}
	return nil
}

//GetBoardIssueList returns all issues by board_id
func (issue *Issue) GetBoardIssueList() ([]map[string]interface{}, error) {

	issueList, err := Session.Query("SELECT id, name, status, description, estimate, user_id,user_first_name, user_last_name, sprint_id, board_id, board_name, project_id,project_name, created_at, updated_at from issues WHERE board_id = ? ALLOW FILTERING", issue.BoardID).Iter().SliceMap()

	if err != nil {
		log.Printf("Error in method GetBoardIssueList inside models/issue.go, method:GetBoardIssueList, error: %s\n", err.Error())
		return nil, err
	}

	return issueList, nil

}

//GetSprintIssueList returns all issues by board_id
func (issue *Issue) GetSprintIssueList() ([]Issue, error) {

	issues := []Issue{}
	var row map[string]interface{}

	iterator := Session.Query("SELECT id, name, status, description, estimate, user_id,user_first_name, user_last_name, sprint_id, board_id, board_name, project_id,project_name, created_at, updated_at from issues WHERE sprint_id = ? ALLOW FILTERING", issue.SprintID).Iter()

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
		log.Printf("Error in method GetSprintIssueList inside models/issue.go: %s\n", err.Error())
		return nil, err
	}

	return issues, nil
}

// //GetClaims Return list of claims to generate jwt token
// func (user *User) GetClaims() map[string]interface{} {
// 	claims := make(map[string]interface{})

// 	claims["UUID"] = user.UUID

// 	return claims
// }
