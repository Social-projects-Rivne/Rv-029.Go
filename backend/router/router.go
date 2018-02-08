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

	projectRouter := Router.PathPrefix("/project").Subrouter()
	applyProjectsRoutes(projectRouter)
	projectRouter.Use(middlewares.AuthenticatedMiddleware)



}

func applyAuthRoutes(r *mux.Router)  {
	r.HandleFunc("/login/", controllers.Login)
	r.HandleFunc("/login", controllers.Login)

	r.HandleFunc("/register/", controllers.Register)
	r.HandleFunc("/register", controllers.Register)

	r.HandleFunc("/confirm", controllers.ConfirmRegistration)
	r.HandleFunc("/forget-password", controllers.ForgotPassword)
	r.HandleFunc("/new-password", controllers.ResetPassword)
}

func applyAuthorizedUserRoutes(r *mux.Router)  {
	r.HandleFunc("/", controllers.Dashboard)
	r.HandleFunc("", controllers.Dashboard)
}

func applyProjectsRoutes(r *mux.Router)  {

	r.HandleFunc("/create/", controllers.StoreProject).Methods("POST")
	r.HandleFunc("/create", controllers.StoreProject).Methods("POST")

	r.HandleFunc("/update/{id}/", controllers.UpdateProject).Methods("PUT")
	r.HandleFunc("/update/{id}", controllers.UpdateProject).Methods("PUT")

	r.HandleFunc("/delete/{id}/", controllers.DeleteProject).Methods("DELETE")
	r.HandleFunc("/delete/{id}", controllers.DeleteProject).Methods("DELETE")

	r.HandleFunc("/show/{id}/", controllers.ShowProjects).Methods("GET")
	r.HandleFunc("/show/{id}", controllers.ShowProjects).Methods("GET")

	r.HandleFunc("/list", controllers.ProjectsList).Methods("GET")
	r.HandleFunc("/list", controllers.ProjectsList).Methods("GET")
}

//func applyAdminRoutes(r *mux.Router)  {
//	r.HandleFunc("/users", controllers.Users)
//}
