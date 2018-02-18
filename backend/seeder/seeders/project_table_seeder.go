package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
)

type ProjectTableSeeder struct {
}

func (ProjectTableSeeder) Run() {


	//id , err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	//if err != nil {
	//	log.Fatal("Can't parse uuid ",err)
	//}

	project := models.Project{
		UUID:      gocql.TimeUUID(),
		Name:      "project number one",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	project.Insert()

	project = models.Project{
		UUID:      gocql.TimeUUID(),
		Name:      "project number two",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	project.Insert()

	



}
