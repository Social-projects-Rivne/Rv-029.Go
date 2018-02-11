package models

import (
	"time"
	//"fmt"
	//"log"

	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
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
	UUID      gocql.UUID `cql:"id" key:"primery"`
	Name      string     `cql:"name"`
	Status    string     `cql:"status"`
	UserID    gocql.UUID `cql:"user_id"`
	SprintID  gocql.UUID `cql:"sprint_id"`
	BoardID   gocql.UUID `cql:"board_id"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}

//Insert func inserts user object in database
func (issue *Issue) Insert() {

	if err := Session.Query(`INSERT INTO issues (id,name,status,user_id,
		sprint_id,board_id,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?);`,
		issue.UUID, issue.Name, issue.Status, issue.UserID, issue.SprintID, issue.BoardID,
		issue.CreatedAt, issue.UpdatedAt).Exec(); err != nil {
	}

}

//Update updates user by id
// func (user *User) Update() {

// 	if err := Session.Query(`Update issues SET name = ? ,updated_at = ? WHERE id= ? ;`,
// 		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {
// 		fmt.Println(err)
// 	}

// }

//Delete removes issue by id
func (issue *Issue) Delete() {

	if err := Session.Query(`DELETE FROM issues WHERE id= ? ;`,
		issue.UUID).Exec(); err != nil {
	}

}

//FindByID finds issue by id
func (issue *Issue) FindByID(id string) error {
	return Session.Query(`SELECT id, name, status, user_id, sprint_id, board_id, created_at	
		FROM issues WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).Scan(&issue.UUID,&issue.Name,&issue.Status,&issue.UserID,
			&issue.SprintID,&issue.BoardID,&issue.CreatedAt)
}

//GetAll returns all users
func (issue *Issue) GetAll() ([]map[string]interface{}, error) {

	return Session.Query(`SELECT * FROM issues`).Iter().SliceMap()

}

// //GetClaims Return list of claims to generate jwt token
// func (user *User) GetClaims() map[string]interface{} {
// 	claims := make(map[string]interface{})

// 	claims["UUID"] = user.UUID

// 	return claims
// }
