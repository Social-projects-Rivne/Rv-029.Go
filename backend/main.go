package main

import (
	"log"
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)

func main() {
	defer db.Session.Close()

	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
