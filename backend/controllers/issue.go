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

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Issue has created",
	})

	setSuccessResHeaders(w, jsonResponse)
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

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Issue has updated",
	})

	setSuccessResHeaders(w, jsonResponse)
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
		log.Printf("Error occured in controller.DeleteIssue while parsing id variable: %v", err)
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	issue.Delete()

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Issue has deleted",
	})

	setSuccessResHeaders(w, jsonResponse)
}

//BoardIssueslist returns list of issues order by board_id
func BoardIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error occured in controller.BoardIssuesList while parsing id variable: %v", err)
	}

	issue := models.Issue{}
	issue.BoardID = id

	boardIssueList := issue.GetBoardIssueList()

	jsonResponse, _ := json.Marshal(boardIssueList)

	setSuccessResHeaders(w, jsonResponse)
}

//SprintIssueslist returns list of issues order by sprint_id
func SprintIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["sprint_id"])

	if err != nil {
		log.Printf("Error occured in controller.SprintIssueslist while parsing id variable: %v", err)
	}

	issue := models.Issue{}
	issue.SprintID = id

	boardIssueList := issue.GetSprintIssueList()

	jsonResponse, _ := json.Marshal(boardIssueList)

	setSuccessResHeaders(w, jsonResponse)
}
