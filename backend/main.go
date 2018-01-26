package main

import (
	"log"
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
	"fmt"
)

func main() {
	fmt.Printf("%T\n", router.Router)
	fmt.Printf("%v\n", router.Router)

	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
