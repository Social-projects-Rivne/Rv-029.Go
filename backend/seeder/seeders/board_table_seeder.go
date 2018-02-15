package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
)

type BoardTableSeeder struct {
}

func (BoardTableSeeder) Run() {

	id, err := gocql.ParseUUID("9325624a-0ba2-22e8-ba34-c06ebf83499a")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	projectID, err := gocql.ParseUUID("4aa8434e-1177-11e8-ba8e-c85b76da292c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	board := &models.Board{
		ID:          id,
		ProjectID:   projectID,
		ProjectName: "project number one",
		Name:        "Seeder board 1",
		Desc:        "Some description 1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := board.Insert(); err != nil {
		log.Printf("Error occured in seeder/seeders/board_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
		return
	}




	id, err = gocql.ParseUUID("93ab624a-1cb2-228a-ba34-c06ebf83322c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	projectID, err = gocql.ParseUUID("78c0071e-1179-11e8-b672-c85b76da292c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	board = &models.Board{
		ID:          id,
		ProjectID:   projectID,
		ProjectName: "project number two",
		Name:        "Seeder board 2",
		Desc:        "Some description 2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := board.Insert(); err != nil {
		log.Printf("Error occured in seeder/seeders/board_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
		return
	}

}
