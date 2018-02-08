package models

import (
	"fmt"
	"time"

	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/gocql/gocql"
)

//Role .
type Role string

const ROLE_ADMIN = "Admin"
const ROLE_STAFF = "Staff"
const ROLE_OWNER = "Owner"
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

//Insert func inserts user object in database
func (user *User) Insert() {

	if err := Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,salt,role,status,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?);	`,
		user.UUID, user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.Status, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Update updates user by id
func (user *User) Update() {

	if err := Session.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

}

//Delete removes user by id
func (user *User) Delete() {

	if err := Session.Query(`DELETE FROM users WHERE id= ? ;`,
		user.UUID).Exec(); err != nil {
		fmt.Println(err)
	}

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
