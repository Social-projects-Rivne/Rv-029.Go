package models

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
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

//Insert func inserts user object in database
func (user *User) Insert() {

	if err := db.Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,salt,role,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?);	`,
		gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//FindByID func finds user from database
func (user *User)FindByID(id string){

	if err := db.Session.Query(`SELECT id, email, first_name, last_name, password, salt, role, created_at, updated_at FROM users WHERE id = ? LIMIT 1`,
		id).Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Salt, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		fmt.Println(err)
	}

}

//FindByEmail func finds user from database
func (user *User)FindByEmail(email string){

	if err := db.Session.Query(`SELECT id, email, first_name, last_name, password, salt, role, created_at, updated_at FROM users WHERE email = ? LIMIT 1 ALLOW FILTERING`,
		email).Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Salt, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		fmt.Println(err)
	}

}

//UpdateEmail func updates user's email in database(don't touch this shit)
func (user *User) UpdateEmail(Session *gocql.Session) {

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

	claims["UUID"] = user.UUID

	return claims
}