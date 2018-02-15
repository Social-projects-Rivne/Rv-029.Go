package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/gocql/gocql"
	"log"
)

type UsersTableSeeder struct {

}

func (UsersTableSeeder) Run() {

	id , err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}
	uuid, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	if err != nil {
		log.Fatalf("Invalid gocql.UUID inputed during user seeding. Error: %s", err.Error())
	}
	projects := map[gocql.UUID]string{uuid: "project number one"}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID:      gocql.TimeUUID(),
		Email:     "user@gmail.com",
		FirstName: "Jon",
		LastName:  "Jones",
		Password:  password.EncodePassword(password.EncodeMD5("qwerty1234"), salt),
		Salt:      salt,
		Role:      models.ROLE_USER,
		Status:	   1,
		Projects:  projects,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.Insert()

	salt = password.GenerateSalt(8)
	user = models.User{
		UUID:      id,
		Email:     "owner@gmail.com",
		FirstName: "Daniel",
		LastName:  "Rigs",
		Salt:      salt,
		Status:		1,
		Password:  password.EncodePassword(password.EncodeMD5("qwerty1234"), salt),
		Role:      models.ROLE_OWNER,
		Projects:  projects,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}


	user.Insert()


}
