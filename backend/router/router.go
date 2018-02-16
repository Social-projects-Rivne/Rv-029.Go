package router

import (
	"net/http"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/controllers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/middlewares"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
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

	sprintRouter := Router
	applySprintRoutes(sprintRouter)
	//sprintRouter.Use(middlewares.AuthenticatedMiddleware)

	projectRouter := Router.PathPrefix("/project").Subrouter()
	applyProjectsRoutes(projectRouter)
	projectRouter.Use(middlewares.AuthenticatedMiddleware)
	projectRouter.Use(middlewares.RoleMiddleware)

	issueRouter := Router.PathPrefix("/project/board/").Subrouter()
	applyIssueRoutes(issueRouter)

}

func applyAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login/", controllers.Login)
	r.HandleFunc("/login", controllers.Login)

	r.HandleFunc("/register/", controllers.Register)
	r.HandleFunc("/register", controllers.Register)

	r.HandleFunc("/confirm", controllers.ConfirmRegistration)
	r.HandleFunc("/forget-password", controllers.ForgotPassword)
	r.HandleFunc("/new-password", controllers.ResetPassword)
}

func applyAuthorizedUserRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.Dashboard)
	r.HandleFunc("", controllers.Dashboard)
}

func applyProjectsRoutes(r *mux.Router) {

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

func applyBoardRoutes(r *mux.Router) {
	r.HandleFunc("/project/{project_id}/board/create/", controllers.CreateBoard).Methods("POST")
	r.HandleFunc("/project/{project_id}/board/create", controllers.CreateBoard).Methods("POST")

	r.HandleFunc("/project/board/update/{board_id}/", controllers.UpdateBoard).Methods("PUT")
	r.HandleFunc("/project/board/update/{board_id}", controllers.UpdateBoard).Methods("PUT")

	r.HandleFunc("/project/board/delete/{board_id}/", controllers.DeleteBoard).Methods("DELETE")
	r.HandleFunc("/project/board/delete/{board_id}", controllers.DeleteBoard).Methods("DELETE")

	r.HandleFunc("/project/board/select/{board_id}/", controllers.SelectBoard).Methods("GET")
	r.HandleFunc("/project/board/select/{board_id}", controllers.SelectBoard).Methods("GET")

	r.HandleFunc("/project/{project_id}/board/list/", controllers.BoardsList).Methods("GET")
	r.HandleFunc("/project/{project_id}/board/list", controllers.BoardsList).Methods("GET")
}

func applyIssueRoutes(r *mux.Router) {
	r.HandleFunc("{board_id}/issue/create/", controllers.StoreIssue).Methods("POST")
	r.HandleFunc("{board_id}/issue/create", controllers.StoreIssue).Methods("POST")

	r.HandleFunc("issue/update/{issue_id}/", controllers.UpdateIssue).Methods("PUT")
	r.HandleFunc("issue/update/{issue_id}", controllers.UpdateIssue).Methods("PUT")

	r.HandleFunc("issue/delete/{issue_id}/", controllers.DeleteIssue).Methods("DELETE")
	r.HandleFunc("issue/delete/{issue_id}", controllers.DeleteIssue).Methods("DELETE")

	r.HandleFunc("{board_id}/issue/list/", controllers.BoardIssueslist).Methods("GET")
	r.HandleFunc("{board_id}/issue/list", controllers.BoardIssueslist).Methods("GET")

	r.HandleFunc("sprint/{sprint_id}/issue/list/", controllers.SprintIssueslist).Methods("GET")
	r.HandleFunc("sprint/{sprint_id}/issue/list", controllers.SprintIssueslist).Methods("GET")

	r.HandleFunc("issue/show/{issue_id}/", controllers.ShowIssue).Methods("GET")
	r.HandleFunc("issue/show/{issue_id}", controllers.ShowIssue).Methods("GET")

}

func applySprintRoutes(r *mux.Router) {
	r.HandleFunc("/project/board/{board_id}/sprint/create/", controllers.CreateSprint).Methods("POST")
	r.HandleFunc("/project/board/{board_id}/sprint/create", controllers.CreateSprint).Methods("POST")

	r.HandleFunc("/project/board/sprint/update/{sprint_id}/", controllers.UpdateSprint).Methods("PUT")
	r.HandleFunc("/project/board/sprint/update/{sprint_id}", controllers.UpdateSprint).Methods("PUT")

	r.HandleFunc("/project/board/sprint/show/{sprint_id}/", controllers.SelectSprint).Methods("GET")
	r.HandleFunc("/project/board/sprint/show/{sprint_id}", controllers.SelectSprint).Methods("GET")

	r.HandleFunc("/project/board/sprint/delete/{sprint_id}/", controllers.DeleteSprint).Methods("DELETE")
	r.HandleFunc("/project/board/sprint/delete/{sprint_id}", controllers.DeleteSprint).Methods("DELETE")

	r.HandleFunc("/project/board/{board_id}/sprint/list/", controllers.SprintsList).Methods("GET")
	r.HandleFunc("/project/board/{board_id}/sprint/list", controllers.SprintsList).Methods("GET")
}
