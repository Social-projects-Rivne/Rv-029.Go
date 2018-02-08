package main

import (
	"log"
	"net/http"
	"os"

	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/rs/cors"
)

func main() {
	var cmd string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "listen"
	}

	switch cmd {
	case "db:seed":
		seeder.Run()
	default:
		handler := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
			AllowedHeaders: []string{"*"},
		}).Handler(router.Router)
		defer db.GetInstance().Session.Close()
		log.Fatal(http.ListenAndServe(":8080", handler))

	}
}
