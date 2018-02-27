package helpers

import (
	"github.com/gocql/gocql"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"time"

)

func InitFakeUser() (models.User,error) {
	projectId , err := gocql.ParseUUID("fc3a1850-0f46-11e8-b192-d8cb8ac536c8")
	if err != nil {
		return models.User{}, fmt.Errorf("invalid UUID ")
	}

	userId, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba15-b06ebf83499f")
	if err != nil {
		return models.User{}, fmt.Errorf("invalid UUID ")
	}
	projects := map[gocql.UUID]string{projectId: "project number one"}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID:      userId,
		Email:     "owner@gmail.com",
		FirstName: "Jon",
		LastName:  "Jones",
		Password:  password.EncodePassword(password.EncodeMD5("qwerty1234"), salt),
		Salt:      salt,
		Role:      models.ROLE_OWNER,
		Status:	   1,
		Projects:  projects,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user , nil
}