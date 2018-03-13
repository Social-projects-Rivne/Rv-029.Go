package models

import (
	"time"

	"log"

	"github.com/gocql/gocql"
)

const (
	//ROLE_ADMIN .
	ROLE_ADMIN = "Admin"

	//ROLE_STAFF .
	ROLE_STAFF = "Staff"

	//ROLE_OWNER .
	ROLE_OWNER = "Owner"

	//ROLE_USER .
	ROLE_USER = "User"


	CHECK_USER_PASSWORD      = "SELECT id, email, first_name, last_name, projects, updated_at, created_at, password, salt, role, status FROM users WHERE email = ? LIMIT 1 allow filtering"

	UPDATE_USER_PROJECT_ROLE = "UPDATE users SET projects = projects +  ? WHERE id = ?"
	DELETE_USER_PROJECT_ROLE = "DELETE projects[?] FROM users WHERE id= ?"
	GET_PROJECT_USERS_LIST   = "SELECT id, email, first_name, last_name, projects, updated_at, created_at, password, salt, role, status from users WHERE projects CONTAINS KEY ?"
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
	Status    int
	Projects  map[gocql.UUID]string
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
	AddRoleToProject(projectId gocql.UUID, role string, userId gocql.UUID) error
	DeleteProject(projectId gocql.UUID, userId gocql.UUID) error
	GetProjectUsersList(projectId gocql.UUID)  ([]User, error)
	CheckUserPassword(User) (User, error)
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

	if err := u.DB.Query(`INSERT INTO users (id,email,first_name,last_name,password,
		salt,role,status,projects,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?);	`,
		user.UUID, user.Email, user.FirstName, user.LastName, user.Password,
		user.Salt, user.Role, user.Status,user.Projects, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {

		log.Printf("Error occured while inserting user %v", err)
		return err
	}
	return nil
}

//Update func finds user from database
func (u *UserStorage) Update(user *User) error {

	if err := u.DB.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {

		log.Printf("Error occured while updating user %v", err)
		return err
	}
	return nil
}

//UpdateByID updates user by his id
func (u *UserStorage) UpdateByID(user *User) error {

	if err := u.DB.Query(`Update users SET password = ? ,updated_at = ? WHERE id= ? ;`,
		user.Password, user.UpdatedAt, user.UUID).Exec(); err != nil {

		log.Printf("Error occured while updating user %v", err)
		return err
	}
	return nil
}

//Delete removes user by his id
func (u *UserStorage) Delete(user *User) error {

	if err := u.DB.Query(`DELETE FROM users WHERE id= ? ;`,
		user.UUID).Exec(); err != nil {

			log.Printf("Error occured in models/user.go, method: Delete, error: %v", err)
		return err
	}
	return nil
}

//FindByID finds user by id
func (u *UserStorage) FindByID(user *User) error {
	if err := u.DB.Query(`SELECT id, email, first_name, last_name,
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
	if err := u.DB.Query(`SELECT id, email, first_name, last_name, password, salt, role, status, 
		projects, created_at, updated_at FROM users WHERE email = ? LIMIT 1 ALLOW FILTERING`, user.Email).
		Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName, &user.Password,
		&user.Salt, &user.Role, &user.Status, &user.Projects, &user.CreatedAt, &user.UpdatedAt); err != nil {

		log.Printf("Error occured in models/user.go, method: FindByEmail, error: %v", err)
		return err
	}
	return nil
}

//GetClaims Return list of claims to generate jwt token
func (user User) GetClaims() map[string]interface{} {
	claims := make(map[string]interface{})

	claims["UUID"] = user.UUID

	return claims
}

// PROJECTS METHODS
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

func (u *UserStorage) GetProjectUsersList(projectId gocql.UUID) ([]User, error)  {


	var users []User
	var row map[string]interface{}
	var pageState []byte

	iterator := Session.Query(GET_PROJECT_USERS_LIST,projectId).Consistency(gocql.One).PageState(pageState).PageSize(5).Iter()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			users = append(users, User{
				UUID:      row["id"].(gocql.UUID),
				Email:     row["email"].(string),
				FirstName: row["first_name"].(string),
				LastName:  row["last_name"].(string),
				//Projects:  row["projects"].(map[gocql.UUID]string),
				CreatedAt: row["created_at"].(time.Time),
				UpdatedAt: row["updated_at"].(time.Time),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in models/user.go error: %+v",err)
	}

	return users, nil

}

// END PROJECTS METHODS

func (u *UserStorage) CheckUserPassword(user User) (User, error) {

	if err := u.DB.Query(CHECK_USER_PASSWORD, user.Email).
		Consistency(gocql.One).Scan(&user.UUID, &user.Email, &user.FirstName, &user.LastName,
		&user.Projects, &user.UpdatedAt, &user.CreatedAt, &user.Password, &user.Salt, &user.Role, &user.Status); err != nil {

		log.Printf("Error in models/user.go error: %+v", err)
		user.UUID, err = gocql.ParseUUID(" ")

		return user, err
	}

	return user, nil
}




func (u *UserStorage) CheckUserEmail(user User) (User, error) {

	if err := u.DB.Query(CHECK_USER_PASSWORD, user.Email).
		Consistency(gocql.One).Scan(&user.Password, &user.Salt, &user.UUID); err != nil {

		log.Printf("Error in models/user.go error: %+v", err)
		user.UUID, err = gocql.ParseUUID(" ")

		return user, err
	}

	return user, nil
}

