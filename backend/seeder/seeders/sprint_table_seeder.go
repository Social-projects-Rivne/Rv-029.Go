package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
	"github.com/icrowley/fake"
)

//SprintTableSeeder model
type SprintTableSeeder struct {
}

var sprints []models.Sprint

//Run .
func (SprintTableSeeder) Run() {

	sprints = []models.Sprint{}

	var status string

	for _, board := range boards {
		for i:=0; i <= 3; i++ {
			if i != 3 {
				status = models.SPRINT_STAUS_DONE
			} else {
				status = models.SPRINT_STAUS_IN_PROGRESS
			}

			sprint := models.Sprint{
				ID:          gocql.TimeUUID(),
				ProjectId:   board.ProjectID,
				ProjectName: board.ProjectName,
				BoardId:     board.ID,
				BoardName:   board.Name,
				Goal:        fake.WordsN(5),
				Desc:        fake.Sentences(),
				Status:      status,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := sprint.Insert(); err != nil {
				log.Printf("Error occured in seeder/seeders/board_table_seeder.go method: Run,where: board.sprint error: %+v", err)
				return
			}

			sprints = append(sprints, sprint)
		}
	}

}
