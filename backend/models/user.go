package models

import (
	"time"

	"github.com/gocql/gocql"
	"log"
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


//Projects queries
const UPDATE_USER_PROJECT_ROLE  = "UPDATE users SET projects = projects +  ? WHERE id = ?"
const DELETE_USER_PROJECT_ROLE  = "DELETE projects[?] FROM users WHERE id= ?"

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
	Projects  map[gocql.UUID] string       `cql:"status"`
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
		salt,role,status,projects,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?);	`,
		user.UUID, user.Email, user.FirstName, user.LastName, user.Password,
		user.Salt, user.Role, user.Status,user.Projects, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {

			log.Printf("Error in models/user.go error: %+v",err)
			return err
	}
	return nil
}

//Update func finds user from database
func (user *User) Update() error {

	if err := Session.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {

			log.Printf("Error in models/user.go error: %+v",err)
		return err
	}
	return nil
}

//UpdateByID updates user by his id
func (user *User) UpdateByID() error {

	if err := Session.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {

			log.Printf("Error in models/user.go error: %+v",err)
		return err
	}
	return nil
}

//Delete removes user by his id
func (user *User) Delete() error {

	if err := Session.Query(`DELETE FROM users WHERE id= ? ;`,
		user.UUID).Exec(); err != nil {

			log.Printf("Error in models/user.go error: %+v",err)
		return err
	}
	return nil
}

//FindByID finds user by id
func (user *User) FindByID() error {
	if err := Session.Query(`SELECT id, email, first_name, last_name,
		 projects, updated_at, created_at, password, salt, role, status FROM users WHERE id = ? LIMIT 1`, user.UUID).
		Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName,
		&user.Projects, &user.UpdatedAt, &user.CreatedAt, &user.Password, &user.Salt, &user.Role, &user.Status); err != nil {

			log.Printf("Error in models/user.go error: %+v",err)
		return err
	}
	return nil
}

//FindByEmail finds user by email
func (user *User) FindByEmail() error {
	if err := Session.Query(`SELECT id, email, first_name, last_name, password, salt, role, status, 
		projects, created_at, updated_at FROM users WHERE email = ? LIMIT 1 ALLOW FILTERING`, user.Email).
		Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password,
		&user.Salt, &user.Role, &user.Status, &user.Projects, &user.CreatedAt, &user.UpdatedAt); err != nil {

			log.Printf("Error in models/user.go error: %+v",err)
		return err
	}
	return nil
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

/*
* Projects methods
*/

func (user *User) AddRoleToProject(projectId gocql.UUID,role string) error  {
	roleMap := make(map[gocql.UUID]string)
	roleMap[projectId] = role
	err := Session.Query(UPDATE_USER_PROJECT_ROLE,roleMap,user.UUID).Exec()

	if err != nil {
		log.Printf("Error in models/user.go error: %+v",err)
		return err
	}

	return nil

}

func (user *User) DeleteProject(projectId gocql.UUID) error  {

	err := Session.Query(DELETE_USER_PROJECT_ROLE,projectId,user.UUID).Exec()

	if err != nil {
		log.Printf("Error in models/user.go error: %+v",err)
		return err
	}

	return nil

}