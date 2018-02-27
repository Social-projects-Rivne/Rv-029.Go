package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
	"github.com/icrowley/fake"
)

//BoardTableSeeder model
type BoardTableSeeder struct {
}

var boards []models.Board

//Run .
func (BoardTableSeeder) Run() {

	boards = []models.Board{}

	for _, project := range projects {
		for i := 0; i < random(2, 5); i++ {
			board := models.Board{
				ID:          gocql.TimeUUID(),
				ProjectID:   project.UUID,
				ProjectName: project.Name,
				Name:        fake.WordsN(3),
				Desc:        fake.Sentences(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := models.BoardDB.Insert(&board); err != nil {
				log.Printf("Error occured in seeder/seeders/board_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
				return
			}

			boards = append(boards, board)
		}
	}
}
