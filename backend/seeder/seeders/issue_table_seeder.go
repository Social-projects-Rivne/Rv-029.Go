package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
)

type IssueTableSeeder struct {
}

func (IssueTableSeeder) Run() {

	id, err := gocql.ParseUUID("9228322a-1ca2-33e8-ba28-c06e22a3322c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	projectID, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}
	userID, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	issue := &models.Issue{
		UUID: id,         
		Name: "Seeder issue 1",         
		Status: models.STATUS_TODO,
		Description: "Seeder description 1",	 
		Estimate: 3,	  
		UserID: userID,     
		UserFirstName: "Jon",
		UserLastName: "Jones" 
		SprintID:      
		BoardID      
		BoardName    
		ProjectID    
		ProjectName  
		CreatedAt    
		UpdatedAt    
	}

	if err := board.Insert(); err != nil {
		log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
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
		log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
		return
	}

}
