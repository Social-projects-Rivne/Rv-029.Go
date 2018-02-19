package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)



func main() {
	var cmd string

	f, err := os.OpenFile(fmt.Sprint("./logs/", time.Now().Format("2006_01_02"), "_task_manager.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
