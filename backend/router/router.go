package router

import (
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/controllers"
)

var Router *mux.Router

func init()  {
	Router = mux.NewRouter()

	authRouter := Router.PathPrefix("/auth").Subrouter()
	applyAuthRoutes(authRouter)

	//adminRouter := Router.PathPrefix("/admin").Subrouter()
	//applyAdminRoutes(adminRouter)
}

func applyAuthRoutes(r *mux.Router)  {
	r.HandleFunc("/login/", controllers.Login)
	r.HandleFunc("/login", controllers.Login)
	//r.HandleFunc("/register/", controllers.Register)
	//r.HandleFunc("/register", controllers.Register)
	//r.HandleFunc("/logout", controllers.Logout)
	//r.HandleFunc("/forget-password", controllers.ForgetPassword)
	//r.HandleFunc("/", controllers.ForgetPassword)
}

//func applyAdminRoutes(r *mux.Router)  {
//	r.HandleFunc("/users", controllers.Users)
//}
