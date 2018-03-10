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
	userRouter.Use(middlewares.OwenrMiddleware)

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
	r.HandleFunc("/login/", controllers.Login).Name(`user.login`)
	r.HandleFunc("/login", controllers.Login).Name(`user.login`)

	r.HandleFunc("/register/", controllers.Register).Name(`user.register`)
	r.HandleFunc("/register", controllers.Register).Name(`user.register`)

	r.HandleFunc("/confirm", controllers.ConfirmRegistration).Name(`user.confirm`)
	r.HandleFunc("/forget-password", controllers.ForgotPassword).Name(`us	er.password.forget`)
	r.HandleFunc("/new-password", controllers.ResetPassword).Name(`user.password.reset`)
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
	r.HandleFunc("/{project_id}/board/create/", controllers.CreateBoard).Methods("POST").Name(`board.create`)
	r.HandleFunc("/{project_id}/board/create", controllers.CreateBoard).Methods("POST").Name(`board.create`)

	r.HandleFunc("/board/update/{board_id}/", controllers.UpdateBoard).Methods("PUT").Name(`board.update`)
	r.HandleFunc("/board/update/{board_id}", controllers.UpdateBoard).Methods("PUT").Name(`board.update`)

	r.HandleFunc("/board/delete/{board_id}/", controllers.DeleteBoard).Methods("DELETE").Name(`board.delete`)
	r.HandleFunc("/board/delete/{board_id}", controllers.DeleteBoard).Methods("DELETE").Name(`board.delete`)

	r.HandleFunc("/board/select/{board_id}/", controllers.SelectBoard).Methods("GET").Name(`board.show`)
	r.HandleFunc("/board/select/{board_id}", controllers.SelectBoard).Methods("GET").Name(`board.show`)

	r.HandleFunc("/{project_id}/board/list/", controllers.BoardsList).Methods("GET").Name(`board.list`)
	r.HandleFunc("/{project_id}/board/list", controllers.BoardsList).Methods("GET").Name(`board.list`)
}

func applyIssueRoutes(r *mux.Router) {
	r.HandleFunc("/{board_id}/issue/create/", controllers.StoreIssue).Methods("POST").Name(`issue.create`)
	r.HandleFunc("/{board_id}/issue/create", controllers.StoreIssue).Methods("POST").Name(`issue.create`)

	r.HandleFunc("/issue/update/{issue_id}/", controllers.UpdateIssue).Methods("PUT").Name(`issue.update`)
	r.HandleFunc("/issue/update/{issue_id}", controllers.UpdateIssue).Methods("PUT").Name(`issue.update`)

	r.HandleFunc("/issue/delete/{issue_id}/", controllers.DeleteIssue).Methods("DELETE").Name(`issue.delete`)
	r.HandleFunc("/issue/delete/{issue_id}", controllers.DeleteIssue).Methods("DELETE").Name(`issue.delete`)

	r.HandleFunc("/{board_id}/issue/list/", controllers.BoardIssueslist).Methods("GET").Name(`issue.backlog.list`)
	r.HandleFunc("/{board_id}/issue/list", controllers.BoardIssueslist).Methods("GET").Name(`issue.backlog.list`)

	r.HandleFunc("/sprint/{sprint_id}/issue/list/", controllers.SprintIssueslist).Methods("GET").Name(`issue.sprint.list`)
	r.HandleFunc("/sprint/{sprint_id}/issue/list", controllers.SprintIssueslist).Methods("GET").Name(`issue.sprint.list`)

	r.HandleFunc("/issue/show/{issue_id}/", controllers.ShowIssue).Methods("GET").Name(`issue.show`)
	r.HandleFunc("/issue/show/{issue_id}", controllers.ShowIssue).Methods("GET").Name(`issue.show`)
}

func applySprintRoutes(r *mux.Router) {
	r.HandleFunc("/{board_id}/sprint/create", controllers.CreateSprint).Methods("POST").Name(`sprint.create`)
	r.HandleFunc("/sprint/update/{sprint_id}", controllers.UpdateSprint).Methods("PUT").Name(`sprint.update`)
	r.HandleFunc("/sprint/show/{sprint_id}", controllers.SelectSprint).Methods("GET").Name(`sprint.show`)
	r.HandleFunc("/sprint/delete/{sprint_id}", controllers.DeleteSprint).Methods("DELETE").Name(`sprint.delete`)
	r.HandleFunc("/{board_id}/sprint/list", controllers.SprintsList).Methods("GET").Name(`sprint.list`)
	r.HandleFunc("/sprint/{sprint_id}/add/issue/{issue_id}", controllers.AddIssueToSprint).Methods("PUT").Name(`sprint.issue.add`)
}
