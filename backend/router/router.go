package router

import (
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/controllers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/middlewares"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	authRouter := Router.PathPrefix("/auth").Subrouter()
	applyAuthRoutes(authRouter)

	userRouter := Router.PathPrefix("/user").Subrouter()
	applyUserRoutes(userRouter)
	userRouter.Use(middlewares.AuthenticatedMiddleware)
	//userRouter.Use(middlewares.OwenrMiddleware)//TODO:owner middleware

	projectRouter := Router.PathPrefix("/project").Subrouter()
	applyProjectsRoutes(projectRouter)
	applyBoardRoutes(projectRouter)

	boardRouter := projectRouter.PathPrefix("/board").Subrouter()
	applyIssueRoutes(boardRouter)
	applySprintRoutes(boardRouter)

	projectRouter.Use(middlewares.AuthenticatedMiddleware)
	//projectRouter.Use(middlewares.ProjectAccessMiddleware)
	projectRouter.Use(middlewares.CheckUserPermission)
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

func applyUserRoutes(r *mux.Router) {
	r.HandleFunc("/{user_id}/add/permission", controllers.AddUserPermission).Name(`user.permissions.add`)
	r.HandleFunc("/{user_id}/remove/permission", controllers.RemoveUserPermissions).Name(`user.permissions.remove`)
	r.HandleFunc("/{user_id}/set/permissions", controllers.SetUserPermissions).Name(`user.permissions.update`)
}

func applyProjectsRoutes(r *mux.Router) {
	r.HandleFunc("/create/", controllers.CreateProject).Methods("POST").Name(`project.create`)
	r.HandleFunc("/create", controllers.CreateProject).Methods("POST").Name(`project.create`)

	r.HandleFunc("/update/{project_id}/", controllers.UpdateProject).Methods("PUT").Name(`project.update`)
	r.HandleFunc("/update/{project_id}", controllers.UpdateProject).Methods("PUT").Name(`project.update`)

	r.HandleFunc("/delete/{project_id}/", controllers.DeleteProject).Methods("DELETE").Name(`project.delete`)
	r.HandleFunc("/delete/{project_id}", controllers.DeleteProject).Methods("DELETE").Name(`project.delete`)

	r.HandleFunc("/show/{project_id}/", controllers.ShowProject).Methods("GET").Name(`project.show`)
	r.HandleFunc("/show/{project_id}", controllers.ShowProject).Methods("GET").Name(`project.show`)

	r.HandleFunc("/list/", controllers.ProjectsList).Methods("GET").Name(`project.list`)
	r.HandleFunc("/list", controllers.ProjectsList).Methods("GET").Name(`project.list`)
}

func applyBoardRoutes(r *mux.Router) {
	r.HandleFunc("/{project_id}/board/create/", controllers.CreateBoard).Methods("POST")
	r.HandleFunc("/{project_id}/board/create", controllers.CreateBoard).Methods("POST")

	r.HandleFunc("/board/update/{board_id}/", controllers.UpdateBoard).Methods("PUT")
	r.HandleFunc("/board/update/{board_id}", controllers.UpdateBoard).Methods("PUT")

	r.HandleFunc("/board/delete/{board_id}/", controllers.DeleteBoard).Methods("DELETE")
	r.HandleFunc("/board/delete/{board_id}", controllers.DeleteBoard).Methods("DELETE")

	r.HandleFunc("/board/select/{board_id}/", controllers.SelectBoard).Methods("GET")
	r.HandleFunc("/board/select/{board_id}", controllers.SelectBoard).Methods("GET")

	r.HandleFunc("/{project_id}/board/list/", controllers.BoardsList).Methods("GET")
	r.HandleFunc("/{project_id}/board/list", controllers.BoardsList).Methods("GET")
}

func applyIssueRoutes(r *mux.Router) {
	r.HandleFunc("/{board_id}/issue/create/", controllers.StoreIssue).Methods("POST")
	r.HandleFunc("/{board_id}/issue/create", controllers.StoreIssue).Methods("POST")

	r.HandleFunc("/issue/update/{issue_id}/", controllers.UpdateIssue).Methods("PUT")
	r.HandleFunc("/issue/update/{issue_id}", controllers.UpdateIssue).Methods("PUT")

	r.HandleFunc("/issue/delete/{issue_id}/", controllers.DeleteIssue).Methods("DELETE")
	r.HandleFunc("/issue/delete/{issue_id}", controllers.DeleteIssue).Methods("DELETE")

	r.HandleFunc("/{board_id}/issue/list/", controllers.BoardIssueslist).Methods("GET")
	r.HandleFunc("/{board_id}/issue/list", controllers.BoardIssueslist).Methods("GET")

	r.HandleFunc("/sprint/{sprint_id}/issue/list/", controllers.SprintIssueslist).Methods("GET")
	r.HandleFunc("/sprint/{sprint_id}/issue/list", controllers.SprintIssueslist).Methods("GET")

	r.HandleFunc("/issue/show/{issue_id}/", controllers.ShowIssue).Methods("GET")
	r.HandleFunc("/issue/show/{issue_id}", controllers.ShowIssue).Methods("GET")

}

func applySprintRoutes(r *mux.Router) {
	r.HandleFunc("/{board_id}/sprint/create", controllers.CreateSprint).Methods("POST")
	r.HandleFunc("/sprint/update/{sprint_id}", controllers.UpdateSprint).Methods("PUT")
	r.HandleFunc("/sprint/show/{sprint_id}", controllers.SelectSprint).Methods("GET")
	r.HandleFunc("/sprint/delete/{sprint_id}", controllers.DeleteSprint).Methods("DELETE")
	r.HandleFunc("/{board_id}/sprint/list", controllers.SprintsList).Methods("GET")
	r.HandleFunc("/sprint/{sprint_id}/add/issue/{issue_id}", controllers.AddIssueToSprint).Methods("PUT")
}
