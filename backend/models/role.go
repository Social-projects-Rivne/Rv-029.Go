package models

import (
	"github.com/gocql/gocql"
	"log"
)

type Role struct {
	Name        string
	Permissions []string
}

type RoleCRUD interface {
	Insert(*Role) error
	Update(*Role) error
	Delete(*Role) error
	FindByName(*Role) error
	List() ([]Role, error)
}

type RoleStorage struct {
	DB *gocql.Session
}

var RoleDB RoleCRUD

func InitRoleDB(crud RoleCRUD) {
	RoleDB = crud
}

//Insert func inserts Role object in database
func (s *RoleStorage) Insert(r *Role) error {
	err := s.DB.Query(`INSERT INTO roles (name, permissions) VALUES (?, ?);`,
		r.Name, r.Permissions).Exec()

	if err != nil {
		log.Printf("Error in models/Role.go error: %+v", err)
		return err
	}

	return nil
}

//Update func updates Role name and description by id
func (s *RoleStorage) Update(r *Role) error {
	err := s.DB.Query(`UPDATE roles SET permissions = ? WHERE name = ?;`,
		r.Permissions, r.Name).Exec()

	if err != nil {
		log.Printf("Error in models/Role.go error: %+v", err)
		return err
	}

	return nil
}

//Delete removes Role by id
func (s *RoleStorage) Delete(r *Role) error {
	err := s.DB.Query(`DELETE FROM roles where name = ?;`, r.Name).Exec()

	if err != nil {
		log.Printf("Error in models/Role.go error: %+v", err)
		return err
	}

	return nil
}

//FindByName func finds Role by id
func (s *RoleStorage) FindByName(r *Role) error {
	err := s.DB.Query(`SELECT name, permissions FROM roles WHERE name = ? LIMIT 1`,
		r.Name).Consistency(gocql.One).Scan(&r.Name, &r.Permissions)

	if err != nil {
		log.Printf("Error in models/Role.go error: %+v", err)
		return err
	}

	return nil
}

//List func return list of Roles
func (s *RoleStorage) List() ([]Role, error) {

	roles := []Role{}
	var row map[string]interface{}

	iterator := s.DB.Query(`SELECT name, permissions FROM roles;`).Iter()

	if iterator.NumRows() > 0 {
		for {
			// New map each iteration
			row = make(map[string]interface{})
			if !iterator.MapScan(row) {
				break
			}

			roles = append(roles, Role{
				Name:          row["name"].(string),
				Permissions:   row["permissions"].([]string),
			})
		}
	}

	if err := iterator.Close(); err != nil {
		log.Printf("Error in models/issue.go error: %+v", err)
		return nil, err
	}

	return roles, nil
}

func (r *Role) HasPermission(permission string) bool {
	for _, value := range r.Permissions {
		if value == permission {
			return true
		}
	}

	return false
}


func (r *Role) AddPermission (permission string) {
	if !r.HasPermission(permission) {
		r.Permissions = append(r.Permissions, permission)
	}
}

func (r *Role) RemovePermission (permission string) {
	for index, value := range r.Permissions {
		if value == permission {
			r.Permissions = append(r.Permissions[:index], r.Permissions[index+1:]...)
		}
	}
}

func (r *Role) SetPermissions (permissions []string) {
	r.Permissions = permissions
}
