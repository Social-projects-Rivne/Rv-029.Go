package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocql/gocql"

	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	//"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/rs/cors"
)

func main() {
	var cmd string
	board := &models.Board{}
	board.ID = gocql.TimeUUID()
	board.ProjectID = gocql.ParseUUID("4aa8434e-1177-11e8-ba8e-c85b76da292c")
	board.ProjectName = "project number two"
	board.Name = "board"


	f, err := os.OpenFile("log_file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "listen"
	}

	issue := &models.Issue{}
	issue.UUID, err = gocql.ParseUUID("ac93ac66-0d10-11e8-a6cd-c85b76da292c")
	issue.FindByID()
	fmt.Println(issue)

	switch cmd {
	case "db:seed":
		seeder.Run()
	default:
		handler := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
			AllowedHeaders: []string{"*"},
		}).Handler(router.Router)
<<<<<<< HEAD
		defer db.GetInstance().Session.Close()
		log.Fatal(http.ListenAndServe(":8080", handler))

=======

		defer db.GetInstance().Session.Close()
		log.Fatal(http.ListenAndServe(":8080", handler))
>>>>>>> origin/f52
	}
}
