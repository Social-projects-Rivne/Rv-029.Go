package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

//StoreIssue .
func StoreIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.IssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			false,
			err.Error(),
		})

		log.Printf("Error occured in controller.StoreIssue while validating: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	boardID, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error occured in controller.StoreIssue while parsing UUID: %v", err)
	}

	issue := &models.Issue{}
	issue.BoardID = boardID
	if issueRequestData.Name != "" {
		issue.Name = issueRequestData.Name
	}
	if issueRequestData.Status != "" {
		issue.Status = issueRequestData.Status
	}
	if issueRequestData.UserID.String() != "" {
		issue.UserID = issueRequestData.UserID
	}
	if issueRequestData.SprintID.String() != "" {
		issue.SprintID = issueRequestData.SprintID
	}
	if issueRequestData.BoardID.String() != "" {
		issue.BoardID = issueRequestData.BoardID
	}
	issue.UUID = gocql.TimeUUID()
	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()

	issue.Insert()

	response := baseResponse{true, "Issue has created"}
	response.Success(w)
}

//UpdateIssue .
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.IssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			false,
			err.Error(),
		})

		log.Printf("Error occured in controller.UpdateIssue while validating: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error occured in controller.UpdateIssue while parsing id variable: %v", err)
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	issue.FindByID()

	if issueRequestData.Name != "" {
		issue.Name = issueRequestData.Name
	}
	if issueRequestData.Status != "" {
		issue.Status = issueRequestData.Status
	}
	if issueRequestData.UserID.String() != "" {
		issue.UserID = issueRequestData.UserID
	}
	if issueRequestData.SprintID.String() != "" {
		issue.SprintID = issueRequestData.SprintID
	}
	if issueRequestData.BoardID.String() != "" {
		issue.BoardID = issueRequestData.BoardID
	}

	issue.UpdatedAt = time.Now()
	issue.Update()

	response := baseResponse{true, "Issue has updated"}
	response.Success(w)

}

func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.IssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			false,
			err.Error(),
		})

		log.Printf("Error occured in controller.DeleteIssue while validating: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		response := baseResponse{false, "Issue ID is not valid"}
		response.Failed(w)
		log.Printf("Issue ID is not valid: %v", err)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID

	if err := issue.Delete(); err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		log.Printf("Error while accessing to database: %v", err)
		return
	}

	response := baseResponse{true, "Issue has deleted"}
	response.Success(w)

}

//BoardIssueslist returns list of issues order by board_id
func BoardIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := baseResponse{false, "Board ID is not valid"}
		response.Failed(w)
		log.Printf("Board ID is not valid: %v", err)
		return
	}

	issue := models.Issue{}
	issue.BoardID = id

	boardIssueList, err := issue.GetBoardIssueList()

	if err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		log.Printf("Error while accessing to database: %v", err)
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
		response := baseResponse{false, "Sprint ID is not valid"}
		response.Failed(w)
		log.Printf("Error occured in controller.SprintIssues while parsing id variable: %v", err)
		return
	}

	issue := models.Issue{}
	issue.SprintID = id

	sprintIssueList, err := issue.GetSprintIssueList()

	if err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		log.Printf("Error while accessing to database: %v", err)
		return
	}

	jsonResponse, _ := json.Marshal(sprintIssueList)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

//ShowIssue send issue obj
func ShowIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		response := baseResponse{false, "Issue ID is not valid"}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = id
	if err := issue.FindByID(); err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		log.Printf("Error while accessing to database: %v", err)
		return
	}

	jsonResponse, _ := json.Marshal(issue)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
