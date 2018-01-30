package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/gocql/gocql"
)

type UsersTableSeeder struct {
}

func (UsersTableSeeder) Run() {

	salt := password.GenerateSalt(8)
	user := &models.User{
		UUID:      gocql.TimeUUID(),
		Email:     "user@gmail.com",
		FirstName: "Some",
		LastName:  "User",
		Salt:      salt,
		Password:  password.EncodePassword(password.EncodeMD5("qwerty1234"), salt),
		Role:      models.ROLE_USER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.Insert()

}
