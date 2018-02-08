package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"	
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
<<<<<<< HEAD
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/rs/cors"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"		
=======
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/rs/cors"
>>>>>>> origin/project
)

func main() {
	var cmd string
	// b := &models.BaseModel{}
	// b.Where("email","=","asdf@df.df")
	// user := b.FindUser()
	// fmt.Println(user)
	user := &models.User{}
	user.Insert()
	user.FindByID("d60b50eb-066d-11e8-8160-c85b76da292c")
	fmt.Println(user)

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
			log.Fatal(http.ListenAndServe(":8080", handler ))
			
	}
}
