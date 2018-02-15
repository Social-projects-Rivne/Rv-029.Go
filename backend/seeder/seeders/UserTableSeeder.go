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

	b := models.BaseModel{}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID:      gocql.TimeUUID(),
		Email:     "test@gmail.com",
		FirstName: "User",
		LastName:  "Goodqwe",
		Salt:      salt,
		Status:    1,
		Password:  password.EncodePassword(password.EncodeMD5("qwerty1234"), salt),
		Role:      models.ROLE_USER,
		CreatedAt: time.Now(),
	}

	b.Insert("users", user)

}
