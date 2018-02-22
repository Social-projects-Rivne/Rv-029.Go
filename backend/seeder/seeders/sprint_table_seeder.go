package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
)

//SprintTableSeeder model
type SprintTableSeeder struct {
}

//Run .
func (SprintTableSeeder) Run() {

	id, err := gocql.ParseUUID("152ac2c0-129b-11e8-b642-0ed5f89f718b")
	if err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	projectID, err := gocql.ParseUUID("fc3a1850-0f46-11e8-b192-d8cb8ac536c8")
	if err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	boardID, err := gocql.ParseUUID("9325624a-0ba2-22e8-ba34-c06ebf83499a")
	if err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	sprint := &models.Sprint{
		ID:          id,
		ProjectId:   projectID,
		ProjectName: "project number one",
		BoardId:     boardID,
		BoardName:   "Seeder board 1",
		Goal:        "Seeder goal 1",
		Desc:        "Seeder desc 1",
		Status:      "Started",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := sprint.Insert(); err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	id, err = gocql.ParseUUID("152ac54a-129b-11e8-b642-0ed5f89f718b")
	if err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	projectID, err = gocql.ParseUUID("fc3aab50-0f46-11e8-b194-d8cb8ac536c8")
	if err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	boardID, err = gocql.ParseUUID("93ab624a-1cb2-228a-ba34-c06ebf83322c")
	if err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

	sprint = &models.Sprint{
		ID:          id,
		ProjectId:   projectID,
		ProjectName: "project number two",
		BoardId:     boardID,
		BoardName:   "Seeder board 2",
		Goal:        "Seeder goal 2",
		Desc:        "Seeder desc 2",
		Status:      "Ended",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := sprint.Insert(); err != nil {
		log.Printf("Error in seeder/seeders/sprint_table_seeder.go error: %+v",err)
		return
	}

}
