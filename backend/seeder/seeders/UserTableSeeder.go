package seeder

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
)

type UsersTableSeeder struct {
}

func (UsersTableSeeder) Run() {

	user := &models.User{ UUID: gocql.TimeUUID() }
	user.Insert()


}