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

	boardRouter := Router
	applyBoardRoutes(boardRouter)
	//boardRouter.Use(middlewares.AuthenticatedMiddleware)
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

func applyBoardRoutes(r *mux.Router) {
	r.HandleFunc("/project/{project_id}/board/create/", controllers.StoreBoard).Methods("POST")
	r.HandleFunc("/project/{project_id}/board/create", controllers.StoreBoard).Methods("POST")

	r.HandleFunc("/project/board/update/{board_id}/", controllers.UpdateBoard).Methods("PUT")
	r.HandleFunc("/project/board/update/{board_id}", controllers.UpdateBoard).Methods("PUT")

	r.HandleFunc("/project/board/delete/{board_id}/", controllers.DeleteBoard).Methods("DELETE")
	r.HandleFunc("/project/board/delete/{board_id}", controllers.DeleteBoard).Methods("DELETE")

	r.HandleFunc("/project/board/select/{board_id}/", controllers.SelectBoard).Methods("GET")
	r.HandleFunc("/project/board/select/{board_id}", controllers.SelectBoard).Methods("GET")

	r.HandleFunc("/project/{project_id}/board/list/", controllers.BoardsList).Methods("GET")
	r.HandleFunc("/project/{project_id}/board/list", controllers.BoardsList).Methods("GET")
}

//func applyAdminRoutes(r *mux.Router)  {
//	r.HandleFunc("/users", controllers.Users)
//}
