package main

import (
    "net/http"
    "log"
	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
	"encoding/json"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string
	var response map[string]interface{}
	var tweets map[gocql.UUID]string
	response = make(map[string]interface{})
	tweets = make(map[gocql.UUID]string)

	// list all tweets
	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	for iter.Scan(&id, &text) {
		tweets[id] = text
	}
	
	if err := iter.Close(); err != nil {
		response["status"] = false
		response["message"] = err.Error()
		log.Fatal(err)
	} else {
		response["status"] = true
		response["message"] = `Backend Network`
		response["tweets"] = tweets
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8080", r))
}