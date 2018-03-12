package seeder

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"log"
)

//RolesTableSeeder model
type RolesTableSeeder struct {

}

//Run .
func (RolesTableSeeder) Run() {

	adminRole := models.Role{
		Name:      models.ROLE_ADMIN,
		Permissions: []string{
			models.PERMISSION_ADD_ISSUES_TO_SPRINTS,
			models.PERMISSION_CREATE_BOARDS,
			models.PERMISSION_CREATE_ISSUES,
			models.PERMISSION_CREATE_SPRINTS,

			models.PERMISSION_UPDATE_BOARDS,
			models.PERMISSION_UPDATE_ISSUES,
			models.PERMISSION_UPDATE_SPRINTS,

			models.PERMISSION_DELETE_BOARDS,
			models.PERMISSION_DELETE_ISSUES,
			models.PERMISSION_DELETE_SPRINTS,
		},
	}
	err := models.RoleDB.Insert(&adminRole)
	if err != nil {
		log.Fatalf("Admin Role was`n inserted during seeding. Error: %+v", err)
	}

	staffRole := models.Role{
		Name:      models.ROLE_STAFF,
		Permissions: []string{
			models.PERMISSION_ADD_ISSUES_TO_SPRINTS,
			models.PERMISSION_CREATE_ISSUES,
			models.PERMISSION_CREATE_SPRINTS,

			models.PERMISSION_UPDATE_ISSUES,
			models.PERMISSION_UPDATE_SPRINTS,

			models.PERMISSION_DELETE_ISSUES,
			models.PERMISSION_DELETE_SPRINTS,
		},
	}
	err = models.RoleDB.Insert(&staffRole)
	if err != nil {
		log.Fatalf("Staff Role was`n inserted during seeding. Error: %+v", err)
	}

	userRole := models.Role{
		Name:      models.ROLE_USER,
		Permissions: []string{
			models.PERMISSION_ADD_ISSUES_TO_SPRINTS,
			models.PERMISSION_CREATE_ISSUES,
			models.PERMISSION_UPDATE_ISSUES,
			models.PERMISSION_DELETE_ISSUES,
		},
	}
	err = models.RoleDB.Insert(&userRole)
	if err != nil {
		log.Fatalf("User Role was`n inserted during seeding. Error: %+v", err)
	}

}