package models

import (
	"log"
	"time"

	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
)

//Role .
type Role string

//ROLE_ADMIN .
const ROLE_ADMIN = "Admin"

//ROLE_STAFF .
const ROLE_STAFF = "Staff"

//ROLE_OWNER .
const ROLE_OWNER = "Owner"

//ROLE_USER .
const ROLE_USER = "User"

//User type
type User struct {
	UUID      gocql.UUID `cql:"id" key:"primery"`
	Email     string     `cql:"email"`
	FirstName string     `cql:"first_name"`
	LastName  string     `cql:"last_name"`
	Password  string     `cql:"password"`
	Salt      string     `cql:"salt"`
	Role      string     `cql:"role"`
	Status    int        `cql:"status"`
	CreatedAt time.Time  `cql:"created_at"`
	UpdatedAt time.Time  `cql:"updated_at"`
}

//Userer is interface for user struct
type Userer interface {
	Insert() error
	Update() error
	Delete() error
	FindByEmail(string) error
	FindByID(string) error
}

//Insert func inserts user object in database
func (user *User) Insert() error {

	if err := Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,
		salt,role,status,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?);	`,
		user.UUID, user.Email, user.FirstName, user.LastName, user.Password,
		user.Salt, user.Role, user.Status, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		log.Printf("Error occured while inserting %v", err)
		return err
	}
	return nil
}

//Update updates user by id
func (user *User) Update() error {

	if err := Session.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {
		log.Printf("Error occured while updating %v", err)
		return err
	}
	return nil
}

//Delete removes user by id
func (user *User) Delete() error {

	if err := Session.Query(`DELETE FROM users WHERE id= ? ;`,
		user.UUID).Exec(); err != nil {
		log.Printf("Error occured while deleting %v", err)
		return err
	}
	return nil
}

//FindByID finds user by id
func (user *User) FindByID(id string) error {
	return Session.Query(`SELECT id, email, first_name, last_name, updated_at, created_at, password, salt, role, status 
		FROM users WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName,
		&user.UpdatedAt, &user.CreatedAt, &user.Password, &user.Salt, &user.Role, &user.Status)
}

//FindByEmail finds user by email
func (user *User) FindByEmail(email string) error {
	return Session.Query(`SELECT id, email, first_name, last_name, updated_at, created_at, password, salt, role, status 
		FROM users WHERE email = ? LIMIT 1`, email).Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName,
		&user.UpdatedAt, &user.CreatedAt, &user.Password, &user.Salt, &user.Role, &user.Status)
}

//GetAll returns all users
func (user *User) GetAll() ([]map[string]interface{}, error) {

	return Session.Query(`SELECT * FROM users`).Iter().SliceMap()

}

//GetClaims Return list of claims to generate jwt token
func (user *User) GetClaims() map[string]interface{} {
	claims := make(map[string]interface{})

	claims["UUID"] = user.UUID

	return claims
}
