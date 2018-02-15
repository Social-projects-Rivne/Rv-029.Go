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

	projectId1 , err := gocql.ParseUUID("fc3a1850-0f46-11e8-b192-d8cb8ac536c8")
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}
	projectId2 , err := gocql.ParseUUID("fc3aab50-0f46-11e8-b194-d8cb8ac536c8")
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}

	userId1 , err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}
	userId2, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba15-b06ebf83499f")
	if err != nil {
		log.Fatalf("Invalid gocql.UUID inputed during user seeding. Error: %s", err.Error())
	}
	projects := map[gocql.UUID]string{projectId1: "project number one", projectId2:"project number two"}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID:      userId1,
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
		UUID:      userId2,
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
