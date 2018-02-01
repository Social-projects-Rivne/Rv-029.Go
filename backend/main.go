package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
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
		dbConnection := db.GetInstance()
		defer dbConnection.Session.Close()

		log.Fatal(http.ListenAndServe(":8080", router.Router))
	}
}
