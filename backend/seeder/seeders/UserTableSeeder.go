package seeder

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
)

type UsersTableSeeder struct {
}

func (UsersTableSeeder) Run() {

	user := &models.User{}
	user.Insert()


}