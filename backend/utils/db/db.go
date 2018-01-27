package db

import (
	"github.com/gocql/gocql"
	"path/filepath"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type dbConfig struct {
	Hosts 			[]string
	Port 			int
	Keyspace 		string
	Authenticator 	gocql.PasswordAuthenticator
}

//Session it's exporting pointer for using database connection
var Session *gocql.Session

func init() {
	filename, _ := filepath.Abs("./backend/config/db.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := &dbConfig{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Authenticator = config.Authenticator
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	
	Session = session
}
