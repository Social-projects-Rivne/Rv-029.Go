package db

import (
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	Session = session
}
