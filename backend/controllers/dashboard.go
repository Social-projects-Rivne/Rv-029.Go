package controllers

import (
	"fmt"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request)  {
	user := r.Context().Value("user")
	fmt.Printf("%T\n", user)
	fmt.Printf("%v\n", user)
	w.Write([]byte("I'm Authorized"))
}
