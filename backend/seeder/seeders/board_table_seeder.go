package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
)

//BoardTableSeeder model
type BoardTableSeeder struct {
}

//Run .
func (BoardTableSeeder) Run() {

	id, err := gocql.ParseUUID("9325624a-0ba2-22e8-ba34-c06ebf83499a")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	projectID, err := gocql.ParseUUID("fc3a1850-0f46-11e8-b192-d8cb8ac536c8")
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

	if err := models.BoardDB.Insert(board); err != nil {
		log.Printf("Error occured in seeder/seeders/board_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
		return
	}




	id, err = gocql.ParseUUID("93ab624a-1cb2-228a-ba34-c06ebf83322c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	projectID, err = gocql.ParseUUID("fc3aab50-0f46-11e8-b194-d8cb8ac536c8")
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

	if err := models.BoardDB.Insert(board); err != nil {
		log.Printf("Error occured in seeder/seeders/board_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
		return
	}

}
