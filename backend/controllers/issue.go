package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
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

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	boardID, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Fatal(err)
	}

	issue := &models.Issue{}
	issue.BoardID = boardID
	if issueRequestData.Name != ""{
		issue.Name = issueRequestData.Name
	}
	if issueRequestData.Status != ""{
		issue.Status = issueRequestData.Status
	}
	if issueRequestData.UserID.String() != ""{
		issue.UserID = issueRequestData.UserID
	}
	if issueRequestData.SprintID.String() != ""{
		issue.SprintID = issueRequestData.SprintID
	}
	if issueRequestData.BoardID.String() != ""{
		issue.BoardID = issueRequestData.BoardID
	}
	issue.UUID = gocql.TimeUUID()
	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()
	
	issue.Insert()

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Board has created",
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

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error occured while parsing id variable: %v",err)
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	issue.FindByID()

	if issueRequestData.Name != ""{
		issue.Name = issueRequestData.Name
	}
	if issueRequestData.Status != ""{
		issue.Status = issueRequestData.Status
	}
	if issueRequestData.UserID.String() != ""{
		issue.UserID = issueRequestData.UserID
	}
	if issueRequestData.SprintID.String() != ""{
		issue.SprintID = issueRequestData.SprintID
	}
	if issueRequestData.BoardID.String() != ""{
		issue.BoardID = issueRequestData.BoardID
	}

	issue.UpdatedAt = time.Now()
	issue.Update()

	jsonResponse, _ := json.Marshal(boardSuccessResponse{
		true, "Board has updated",
	})

	setSuccessResHeaders(w, jsonResponse)
}