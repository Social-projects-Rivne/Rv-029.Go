package models

import (
	"fmt"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/db"
	"github.com/gocql/gocql"
)

//User type
type User struct {
	UUID      gocql.UUID
	Email     string
	FirstName string
	LastName  string
	Password  string
	Salt      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Insert func insert user object in database
func (user *User) Insert() {
	defer db.Session.Close()
	if err := db.Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,salt,role,created_at,update_at) VALUES (?,?,?,?,?,?,?,?,?);	`,
		gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//UpdateEmail func updates user's email in database(don't touch this shit)
func (user *User) UpdateEmail(Session *gocql.Session) {

	defer db.Session.Close()
	if err := db.Session.Query(`UPDATE example.users 
		SET email = ?
		WHERE email id = ?;`,
		gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

// Return list of claims to generate jwt token
func (user *User) GetClaims() map[string]interface{} {
	claims := make(map[string]interface{})

	claims["id"] = user.UUID.Bytes()
	claims["email"] = user.Email

	return claims
}