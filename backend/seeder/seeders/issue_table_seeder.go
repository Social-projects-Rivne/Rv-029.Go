package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
)

//IssueTableSeeder model
type IssueTableSeeder struct {
}

//Run .
func (IssueTableSeeder) Run() {

	id, err := gocql.ParseUUID("9228322a-1ca2-33e8-ba28-c06e22a3322c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	projectID, err := gocql.ParseUUID("fc3a1850-0f46-11e8-b192-d8cb8ac536c8")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}
	userID, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba34-b06ebf83499f")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}
	sprintID, err := gocql.ParseUUID("152ac2c0-129b-11e8-b642-0ed5f89f718b")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}
	boardID, err :=gocql.ParseUUID("9325624a-0ba2-22e8-ba34-c06ebf83499a")
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
		UserLastName: "Jones", 
		SprintID: sprintID,     
		BoardID: boardID,      
		BoardName: "Seeder board 1",   
		ProjectID: projectID,  
		ProjectName: "project number one",
		CreatedAt: time.Now(),    
		UpdatedAt: time.Now(),    
	}

	if err := issue.Insert(); err != nil {
		log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
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
	userID, err = gocql.ParseUUID("9646324a-0aa2-11e8-ba15-b06ebf83499f")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}
	sprintID, err = gocql.ParseUUID("152ac54a-129b-11e8-b642-0ed5f89f718b")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}
	boardID, err = gocql.ParseUUID("93ab624a-1cb2-228a-ba34-c06ebf83322c")
	if err != nil {
		log.Fatal("Can't parse uuid ", err)
		return
	}

	issue = &models.Issue{
		UUID: id,         
		Name: "Seeder issue 2",         
		Status: models.STATUS_DONE,
		Description: "Seeder description 2",	 
		Estimate: 5,	  
		UserID: userID,     
		UserFirstName: "Nigga",
		UserLastName: "Shit", 
		SprintID: sprintID,     
		BoardID: boardID,      
		BoardName: "Seeder board 2",   
		ProjectID: projectID,  
		ProjectName: "project number two",
		CreatedAt: time.Now(),    
		UpdatedAt: time.Now(), 
	}

	if err := issue.Insert(); err != nil {
		log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %s", err.Error())
		return
	}

}
