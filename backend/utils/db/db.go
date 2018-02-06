package db

import (
	"github.com/gocql/gocql"
	"path/filepath"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"sync"
)

type dbConfig struct {
	Hosts 			[]string
	Port 			int
	Keyspace 		string
	Authenticator 	gocql.PasswordAuthenticator
}

type DB struct {
	Session *gocql.Session
}

var instance *DB
var once sync.Once

func GetInstance() *DB {
	once.Do(func() {
		instance = &DB{}
		instance.init()
	})

	return instance
}


func (db *DB) init() {
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

	db.Session = session
}