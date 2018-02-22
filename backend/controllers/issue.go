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
)

//StoreIssue creates issue in database
func StoreIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.CreateIssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	vars := mux.Vars(r)
	boardID, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
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
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue.UserFirstName = user.FirstName
	issue.UserLastName = user.LastName
	issue.BoardID = boardID

	board := &models.Board{}
	board.ID = issue.BoardID
	if err := board.FindByID(); err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue.BoardName = board.Name
	issue.ProjectID = board.ProjectID
	issue.ProjectName = board.ProjectName
	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()

	if err := issue.Insert(); err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	response := successResponse{true, "Issue has created", nil}
	response.send(w)
}

//UpdateIssue controller updates issue in database
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.CreateIssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	if err := issue.FindByID(); err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
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
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	response := successResponse{true, "Issue has updated", nil}
	response.send(w)

}

//DeleteIssue controller deletes issue from database
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	// var issueRequestData validator.CreateIssueRequestData

	// err := decodeAndValidate(r, &issueRequestData)

	// if err != nil {
	// 	log.Printf("Error occured in controllers/issue.go while decoding JSON, method: DeleteIssue where: %s", err.Error())
	// 	response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go while decoding JSON, metod: DeleteIssue, error: %s",err.Error())}
	// 	response.send(w)
	// 	return
	// }

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID

	if err := issue.Delete(); err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	response := successResponse{true, "Issue has deleted", nil}
	response.send(w)

}

//BoardIssueslist returns list of issues order by board_id
func BoardIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue := models.Issue{}
	issue.BoardID = id

	boardIssueList, err := issue.GetBoardIssueList()

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
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
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue := models.Issue{}
	issue.SprintID = id

	sprintIssueList, err := issue.GetSprintIssueList()

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	jsonResponse, err := json.Marshal(sprintIssueList)
	if err != nil{
		log.Printf("Error in controllers/issue error: %+v",err)
		return		
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

//ShowIssue send issue obj
func ShowIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = id
	if err := issue.FindByID(); err != nil {
		log.Printf("Error in controllers/issue error: %+v",err)
		response := failedResponse{false, fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.send(w)
		return
	}

	jsonResponse, err := json.Marshal(issue)
	if err != nil{
		log.Printf("Error in controllers/issue error: %+v",err)
		return		
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
