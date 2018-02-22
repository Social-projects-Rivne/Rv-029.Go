package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"log"
)

//ProjectTableSeeder model
type ProjectTableSeeder struct {
}

//Run .
func (ProjectTableSeeder) Run() {


	id1 , err := gocql.ParseUUID("fc3a1850-0f46-11e8-b192-d8cb8ac536c8")
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}

	id2 , err := gocql.ParseUUID("fc3aab50-0f46-11e8-b194-d8cb8ac536c8")
	if err != nil {
		log.Fatal("Can't parse uuid ",err)
	}

	project := models.Project{
		UUID:      id1,
		Name:      "project number one",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	models.ProjectDB.Insert(&project)

	project = models.Project{
		UUID:      id2,
		Name:      "project number two",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	models.ProjectDB.Insert(&project)

}
