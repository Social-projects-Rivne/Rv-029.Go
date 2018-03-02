package middlewares

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type rules map[string][]string

var rulesMap = rules{
	"login": []string{"asfa11sf", "a22sfasf", "a33sfasfsa"},
	"register": []string{"a44sfasf", "asfa55sf", "as66fasfsa"},
	"test": []string{"as77fasf", "asf88asf", "as999fasfsa"},
	"test2": []string{"as000fasf", "asf**asf", "asfas^^fsa"},
}

func CheckUserPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(mux.Vars(r))
		fmt.Println(mux.CurrentRoute(r))
		fmt.Println(mux.CurrentRoute(r).GetName())
		fmt.Println(rulesMap)
		next.ServeHTTP(w, r)
	})
}
