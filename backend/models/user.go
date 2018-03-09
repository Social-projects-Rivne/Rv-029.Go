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
	UUID      gocql.UUID
	Email     string
	FirstName string
	LastName  string
	Password  string
	Salt      string
	Role      string
	Status    int
	Projects  map[gocql.UUID] string
	Permissions []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/models UserCRUD

type UserCRUD interface {
	Insert(*User) error
	Update(*User) error
	Delete(*User) error
	FindByID(*User) error
	FindByEmail(*User) error
	AddRoleToProject(projectId gocql.UUID,role string, userId gocql.UUID) error
	DeleteProject(projectId gocql.UUID , userId gocql.UUID) error
}


type UserStorage struct {
	DB *gocql.Session
}

var UserDB UserCRUD

func InitUserDB(crud UserCRUD) {
	UserDB = crud
}


//Insert func inserts user object in database
func (u *UserStorage) Insert(user *User) error {

	if err := Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,
		salt,role,status,projects,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?);	`,
		user.UUID, user.Email, user.FirstName, user.LastName, user.Password,
		user.Salt, user.Role, user.Status,user.Projects, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {

		log.Printf("Error occured while inserting user %v", err)
		return err
	}
	return nil
}

//Update func finds user from database
func (u *UserStorage)  Update(user *User) error {

	if err := Session.Query(`Update users SET password = ? ,updated_at = ?, permissions = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.Permissions, user.UUID).Exec(); err != nil {

		log.Printf("Error occured while updating user %v", err)
		return err
	}
	return nil
}

//UpdateByID updates user by his id
func (u *UserStorage) UpdateByID(user *User) error {

	if err := Session.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {

		log.Printf("Error occured while updating user %v", err)
		return err
	}
	return nil
}

//Delete removes user by his id
func (u *UserStorage) Delete(user *User) error {

	if err := Session.Query(`DELETE FROM users WHERE id= ? ;`,
		user.UUID).Exec(); err != nil {

			log.Printf("Error occured in models/user.go, method: Delete, error: %v", err)
		return err
	}
	return nil
}

//FindByID finds user by id
func (u *UserStorage) FindByID(user *User) error {
	if err := Session.Query(`SELECT id, email, first_name, last_name,
		 projects, updated_at, created_at, password, salt, role, status FROM users WHERE id = ? LIMIT 1`, user.UUID).
		Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName,
		&user.Projects, &user.UpdatedAt, &user.CreatedAt, &user.Password, &user.Salt, &user.Role, &user.Status); err != nil {

		log.Printf("Error occured in models/user.go, method: FindByID, error: %v", err)
		return err
	}
	return nil
}

//FindByEmail finds user by email
func (u *UserStorage) FindByEmail(user *User) error {
	if err := Session.Query(`SELECT id, email, first_name, last_name, password, salt, role, status, 
		projects, created_at, updated_at FROM users WHERE email = ? LIMIT 1 ALLOW FILTERING`, user.Email).
		Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password,
		&user.Salt, &user.Role, &user.Status, &user.Projects, &user.CreatedAt, &user.UpdatedAt); err != nil {

		log.Printf("Error occured in models/user.go, method: FindByEmail, error: %v", err)
		return err
	}
	return nil
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

func (u *UserStorage) AddRoleToProject(projectId gocql.UUID,role string, userId gocql.UUID) error  {
	roleMap := make(map[gocql.UUID]string)
	roleMap[projectId] = role
	err := Session.Query(UPDATE_USER_PROJECT_ROLE,roleMap,userId).Exec()

	if err != nil {
		log.Printf("Error in method AddRoleToProject models/user.go: %s\n", err.Error())
		return err
	}

	return nil

}

func (u *UserStorage) DeleteProject(projectId gocql.UUID , userId gocql.UUID) error  {

	err := Session.Query(DELETE_USER_PROJECT_ROLE,projectId,userId).Exec()

	if err != nil {
		log.Printf("Error in method DeleteProject models/user.go: %s\n", err.Error())
		return err
	}

	return nil

}

func (u *User) HasPermission(permission string) bool {
	for _, value := range u.Permissions {
		if value == permission {
			return true
		}
	}

	return false
}

func (u *User) AddPermission (permission string) {
	if !u.HasPermission(permission) {
		u.Permissions = append(u.Permissions, permission)
	}
}

func (u *User) RemovePermission (permission string) {
	for index, value := range u.Permissions {
		if value == permission {
			u.Permissions = append(u.Permissions[:index], u.Permissions[index+1:]...)
		}
	}
}

func (u *User) SetPermissions (permissions []string) {
	u.Permissions = permissions
}

/***************************************************************************************/

const PERMISSION_CREATE_PROJECTS  = `project.create`
const PERMISSION_MANAGE_USER_PERMISSIONS  = `user.permissions.manage`

func GetPermissionsList() []string {
	return []string{
		PERMISSION_MANAGE_USER_PERMISSIONS,
		PERMISSION_CREATE_PROJECTS,
	}
}