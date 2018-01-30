package router

import (
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/controllers"
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/middlewares"
)

var Router *mux.Router

func init()  {
	Router = mux.NewRouter()
	Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/public"))))

	authRouter := Router.PathPrefix("/auth").Subrouter()
	applyAuthRoutes(authRouter)

	authorizedUserRouter := Router.PathPrefix("/dashboard").Subrouter()
	applyAuthorizedUserRoutes(authorizedUserRouter)
	authorizedUserRouter.Use(middlewares.AuthenticatedMiddleware)

}

func applyAuthRoutes(r *mux.Router)  {
	r.HandleFunc("/login/", controllers.Login)
	r.HandleFunc("/login", controllers.Login)

	r.HandleFunc("/register/", controllers.Register)
	r.HandleFunc("/register", controllers.Register)

	//r.HandleFunc("/logout", controllers.Logout)
	r.HandleFunc("/forget-password", controllers.ForgotPassword)
	//r.HandleFunc("/", controllers.ForgetPassword)
}

func applyAuthorizedUserRoutes(r *mux.Router)  {
	r.HandleFunc("/", controllers.Dashboard)
	r.HandleFunc("", controllers.Dashboard)
}

//func applyAdminRoutes(r *mux.Router)  {
//	r.HandleFunc("/users", controllers.Users)
//}
