package seeder

import (
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"log"

	"github.com/gocql/gocql"
	"fmt"
	"github.com/icrowley/fake"
)

//IssueTableSeeder model
type IssueTableSeeder struct {
}

var issues []models.Issue

//Run .
func (IssueTableSeeder) Run() {

	issues = []models.Issue{}
	for _, board := range boards {
		for i:=0; i < random(5,25); i++ {

			issue := models.Issue{
				UUID: gocql.TimeUUID(),
				Name: fmt.Sprintf("Backlog issue %d", i),
				Status: models.STATUS_TODO,
				Description: fake.SentencesN(2),
				BoardID: board.ID,
				BoardName: board.Name,
				ProjectID: board.ProjectID,
				ProjectName: board.ProjectName,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := issue.Insert(); err != nil {
				log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %+v", err)
				return
			}

			issues = append(issues, issue)
		}
	}

	for _, sprint := range sprints {
		for i:=0; i < random(5,10); i++ {
			issue := models.Issue{
				UUID: gocql.TimeUUID(),
				Name: fmt.Sprintf("Sprint %s issue %d", models.STATUS_TODO, i),
				Status: models.STATUS_TODO,
				Description: fake.SentencesN(2),
				BoardID: sprint.BoardId,
				BoardName: sprint.BoardName,
				ProjectID: sprint.ProjectId,
				ProjectName: sprint.ProjectName,
				SprintID: sprint.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := issue.Insert(); err != nil {
				log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %+v", err)
				return
			}

			issues = append(issues, issue)
		}
		for i:=0; i < random(5,10); i++ {
			issue := models.Issue{
				UUID: gocql.TimeUUID(),
				Name: fmt.Sprintf("Sprint %s issue %d", models.STATUS_IN_PROGRESS, i),
				Status: models.STATUS_IN_PROGRESS,
				Description: fake.SentencesN(2),
				BoardID: sprint.BoardId,
				BoardName: sprint.BoardName,
				ProjectID: sprint.ProjectId,
				ProjectName: sprint.ProjectName,
				SprintID: sprint.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := issue.Insert(); err != nil {
				log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %+v", err)
				return
			}

			issues = append(issues, issue)
		}
		for i:=0; i < random(5,10); i++ {
			issue := models.Issue{
				UUID: gocql.TimeUUID(),
				Name: fmt.Sprintf("Sprint %s issue %d", models.STATUS_DONE, i),
				Status: models.STATUS_DONE,
				Description: fake.SentencesN(2),
				BoardID: sprint.BoardId,
				BoardName: sprint.BoardName,
				ProjectID: sprint.ProjectId,
				ProjectName: sprint.ProjectName,
				SprintID: sprint.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := issue.Insert(); err != nil {
				log.Printf("Error occured in seeder/seeders/issue_table_seeder.go method: Run,where: board.Insert error: %+v", err)
				return
			}

			issues = append(issues, issue)
		}
	}
}
