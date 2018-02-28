package seeder

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/icrowley/fake"
	"log"
	"time"
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

	for i := 0; i < 10; i++ {
		project := models.Project{
			UUID:      gocql.TimeUUID(),
			Name:      fmt.Sprintf("%s #%d", fake.WordsN(3), i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := models.ProjectDB.Insert(&project)
		if err != nil {
			log.Fatalf("Project was`n inserted during seeding. Error: %+v")
		}

		userProjects[project.UUID] = models.ROLE_OWNER
		projects = append(projects, project)
	}
}
