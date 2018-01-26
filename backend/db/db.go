package db

import (
	"github.com/gocql/gocql"
)

//Session it's exporting pointer for using database connection
var Session *gocql.Session

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	
	Session = session
}
