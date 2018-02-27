package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/seeders"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/mail"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/router"
)

type App struct {
	Config *AppConfig
	DB *gocql.Session
	Mailer *mail.SmtpMailer
}

type AppConfig struct {
	DB db.DBConfig `yaml:"db"`
	Mailer mail.SmtpMailerConfig `yaml:"mail"`
	JWT jwt.JWTConfig `yaml:"jwt"`
}

func (app *App) InitApp(path string) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	appConfig := &AppConfig{}

	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	app.Config = appConfig

	app.DB = db.InitFromConfig(app.Config.DB)
	app.Mailer = mail.InitFromConfig(&app.Config.Mailer)
}

var APP *App

func init()  {
	APP = &App{}
	APP.InitApp("./backend/config/app.yml")
}

func main() {
	models.InitBoardDB(&models.BoardStorage{ APP.DB })
	models.InitProjectDB(&models.ProjectStorage{ APP.DB })
	models.InitUserDB(&models.UserStorage{APP.DB})

	var cmd string

	jwt.Config = &APP.Config.JWT
	models.Session = APP.DB
	mail.Mailer = APP.Mailer

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

		defer APP.DB.Close()
		log.Fatal(http.ListenAndServe(":8080", handler))
	}
}