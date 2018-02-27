package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/gocql/gocql"
	"github.com/icrowley/fake"
	"log"
)

//UsersTableSeeder model
type UsersTableSeeder struct {

}

var users []models.User

//Run .
func (UsersTableSeeder) Run() {


	users = []models.User{}

	var email string

	for i := 0; i < 10; i++ {
		salt := password.GenerateSalt(8)

		if i == 0 {
			email = "owner@gmail.com"
		} else if i == 1 {
			email = "user@gmail.com"
		} else {
			email = fake.EmailAddress()
		}

		user := models.User{
			UUID:      gocql.TimeUUID(),
			Email:     email,
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Password:  password.EncodePassword(password.EncodeMD5("qwerty1234"), salt),
			Salt:      salt,
			Role:      models.ROLE_OWNER,
			Status:    1,
			Projects:  userProjects,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err := models.UserDB.Insert(&user)
		if err != nil {
			log.Fatalf("User was`n inserted during seeding. Error: %+v", err)
		}

		users = append(users, user)
	}

}