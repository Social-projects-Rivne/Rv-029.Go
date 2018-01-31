package models

import (
	"fmt"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
)

type Role string

const ROLE_ADMIN = "Admin"
const ROLE_STAFF = "Staff"
const ROLE_OWNER = "Owner"
const ROLE_USER = "User"

//User type
type User struct {
	UUID      gocql.UUID
	Email     string
	FirstName string
	LastName  string
	Password  string
	Salt      string
	Role      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Insert func inserts user object in database
func (user *User) Insert() {

	if err := db.Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,salt,role,status,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?);	`,
		gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.Status, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//FindByID func finds user from database
func (user *User) FindByID(id string) error {
	return db.Session.Query(`SELECT id, email, first_name, last_name, password, salt, role, status, created_at, updated_at FROM users WHERE id = ? LIMIT 1`,
		id).Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Salt, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
}

//FindByEmail func finds user from database
func (user *User) FindByEmail(email string) {

	if err := db.Session.Query(`SELECT id, email, first_name, last_name, password, salt, role, status, created_at, updated_at FROM users WHERE email = ? LIMIT 1 ALLOW FILTERING`,
		email).Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Salt, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		fmt.Println(err)
	}

}

func (user *User) FindByToken (token string) error {
	//TODO:

	return nil
}

//UpdateEmail func updates user's email in database(FIXME: don't touch this shit)
func (user *User) UpdateEmail() {

	if err := db.Session.Query(`UPDATE example.users 
		SET email = ?
		WHERE email id = ?;`,
		gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Update func updates user's email in database(don't touch this shit)
func (user *User) Update() {

	//TODO

}

// Return list of claims to generate jwt token
func (user *User) GetClaims() map[string]interface{} {
	claims := make(map[string]interface{})

	claims["UUID"] = user.UUID

	return claims
}
