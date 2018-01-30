package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "data from server!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data", DataHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/public"))))

	http.ListenAndServe(":8080", r)
}
