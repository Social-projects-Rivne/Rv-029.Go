package db

import (
	"github.com/gocql/gocql"
	"log"
)

type DBConfig struct {
	Hosts         []string                    `yaml:"hosts"`
	Port          int                         `yaml:"port"`
	Keyspace      string                      `yaml:"keyspace"`
	Authenticator gocql.PasswordAuthenticator `yaml:"authenticator"`
}

func InitFromConfig(config DBConfig) *gocql.Session {
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Authenticator = config.Authenticator
	cluster.Consistency = gocql.One

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("DB Init Error: %+v", err)
	}

	return session
}
