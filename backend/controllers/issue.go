package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)

//StoreIssue creates issue in database
func StoreIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.CreateIssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		log.Printf("Error occured in controllers/issue.go while decoding JSON, method: StoreIssue where: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: StoreIssue where: %s", err.Error())}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardID, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error occured in controllers/issue.go method: StoreIssue, where: while parsing board_id %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: StoreIssue where: parsing board_id %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = gocql.TimeUUID()
	issue.Name = issueRequestData.Name
	issue.Status = issueRequestData.Status
	issue.UserID = issueRequestData.UserID

	user := &models.User{}
	user.UUID = issue.UserID
	if err := user.FindByID(); err != nil {
		log.Printf("Error occured in controllers/issue.go, method:StoreIssue, where: user.FindByID, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go, method:StoreIssue, where: user.FindByID, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue.UserFirstName = user.FirstName
	issue.UserLastName = user.LastName
	issue.BoardID = boardID

	board := &models.Board{}
	board.ID = issue.BoardID
	if err := board.FindByID(); err != nil {
		log.Printf("Error occured in controllers/issue.go, method:StoreIssue, where: board.FindByID, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go, method:StoreIssue, where: board.FindByID, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue.BoardName = board.Name
	issue.ProjectID = board.ProjectID
	issue.ProjectName = board.ProjectName
	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()

	if err := issue.Insert(); err != nil {
		log.Printf("Error occured in controllers/issue.go, method:StoreIssue, where: issue.Insert, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go, method:StoreIssue, where: issue.Insert, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Issue has created"}
	response.Success(w)
}

//UpdateIssue controller updates issue in database
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.CreateIssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		log.Printf("Error occured in controllers/issue.go while decoding JSON, method: UpdateIssue where: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while decoding JSON, metod: UpdateIssue, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error occured in controllers/issue.go while parsing issue_id, method: UpdateIssue, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing issue_id, metod: UpdateIssue, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	if err := issue.FindByID(); err != nil {
		log.Printf("Error occured in controllers/issue.go method: UpdateIssue, where: issue.FindByID, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: UpdateIssue, where: issue.FindByID, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue.Name = issueRequestData.Name
	issue.Description = issueRequestData.Description
	issue.UserID = issueRequestData.UserID
	issue.Estimate = issueRequestData.Estimate
	issue.Status = issueRequestData.Status
	issue.SprintID = issueRequestData.SprintID
	issue.UpdatedAt = time.Now()

	if err = issue.Update(); err != nil {
		log.Printf("Error occured in controllers/issue.go method: UpdateIssue, where: issue.Update, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: UpdateIssue, where: issue.Update, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Issue has updated"}
	response.Success(w)

}

//DeleteIssue controller deletes issue from database
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	// var issueRequestData validator.CreateIssueRequestData

	// err := decodeAndValidate(r, &issueRequestData)

	// if err != nil {
	// 	log.Printf("Error occured in controllers/issue.go while decoding JSON, method: DeleteIssue where: %s", err.Error())
	// 	response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while decoding JSON, metod: DeleteIssue, error: %s",err.Error())}
	// 	response.Failed(w)
	// 	return
	// }

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error occured in controllers/issue.go while parsing issue_id, method: DeleteIssue, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing issue_id, metod: DeleteIssue, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID

	if err := issue.Delete(); err != nil {
		log.Printf("Error occured in controllers/issue.go method: DeleteIssue, where: issue.Delete, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: DeleteIssue, where: issue.Delete, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Issue has deleted"}
	response.Success(w)

}

//BoardIssueslist returns list of issues order by board_id
func BoardIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error occured in controllers/issue.go while parsing board_id, method: BoardIssueList where: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing board_id, metod: BoardIssueList, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := models.Issue{}
	issue.BoardID = id

	boardIssueList, err := issue.GetBoardIssueList()

	if err != nil {
		log.Printf("Error occured in controllers/issue.go method: BoardIssueslist, where: issue.GetBoardIssueslist, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: BoardIssueslist, where: issue.GetBoardIssueslist, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	jsonResponse, _ := json.Marshal(boardIssueList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

//SprintIssueslist returns list of issues order by sprint_id
func SprintIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["sprint_id"])

	if err != nil {
		log.Printf("Error occured in controllers/issue.go  while parsing sprint_id, method: SprintIssueList where: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing sprint_id, metod: SprintIssueList, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := models.Issue{}
	issue.SprintID = id

	sprintIssueList, err := issue.GetSprintIssueList()
	if err != nil {
		log.Printf("Error occured in controllers/issue.go method: SprintIssueList, where: issue.GetSprintIssueList, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: SprintIssueList, where: issue.GetSprintIssueList, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	res := helpers.Response{Message: "Done", Data: sprintIssueList}
	res.Success(w)
}

//ShowIssue Failed issue obj
func ShowIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error occured in controllers/issue.go while parsing issue_id, method: ShowIssue, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing issue_id, metod: ShowIssue, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = id
	if err := issue.FindByID(); err != nil {
		log.Printf("Error occured in controllers/issue.go method: ShowIssue, where: issue.FindByID, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: ShowIssue, where: issue.FindByID, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	jsonResponse, _ := json.Marshal(issue)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
