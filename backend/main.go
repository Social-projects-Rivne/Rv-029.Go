package main

import (
	"log"
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
