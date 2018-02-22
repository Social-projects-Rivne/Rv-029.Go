package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"log"
	"github.com/icrowley/fake"
	"fmt"
)

//ProjectTableSeeder model
type ProjectTableSeeder struct {
}

var projects []models.Project
var userProjects map[gocql.UUID]string

//Run .
func (ProjectTableSeeder) Run() {

	projects = []models.Project{}
	userProjects = make(map[gocql.UUID]string)

	for i:=0; i<10; i++ {
		project := models.Project{
			UUID:      gocql.TimeUUID(),
			Name:      fmt.Sprintf("%s #%d", fake.WordsN(3), i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := project.Insert()
		if err != nil {
			log.Fatalf("Project was`n inserted during seeding. Error: %+v")
		}

		userProjects[project.UUID] = models.ROLE_OWNER
		projects = append(projects, project)
	}
}
