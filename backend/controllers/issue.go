package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

//StoreIssue creates issue in database
func StoreIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.CreateIssueRequestData

	err := decodeAndValidate(r, &issueRequestData)

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardID, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = gocql.TimeUUID()
	issue.Name = issueRequestData.Name
	issue.Status = issueRequestData.Status
	issue.Description = issueRequestData.Description
	issue.Estimate = issueRequestData.Estimate
	issue.SprintID = issueRequestData.SprintID
	issue.BoardID = boardID

	board := &models.Board{}
	board.ID = issue.BoardID
	if err := models.BoardDB.FindByID(board); err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	issue.BoardName = board.Name
	issue.ProjectID = board.ProjectID

	issue.ProjectName = board.ProjectName
	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()

	if err := models.IssueDB.Insert(issue); err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Issue has created", StatusCode: http.StatusOK}
	response.Success(w)
}

//UpdateIssue controller updates issue in database
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issueRequestData validator.CreateIssueRequestData

	err := decodeAndValidate(r, &issueRequestData)
	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)

	issueID, err := gocql.ParseUUID(vars["issue_id"])
	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	if err := models.IssueDB.FindByID(issue); err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	issue.Name = issueRequestData.Name
	issue.Description = issueRequestData.Description
	issue.UserID = issueRequestData.UserID
	issue.Estimate = issueRequestData.Estimate
	issue.Status = issueRequestData.Status
	issue.UpdatedAt = time.Now()

	if err = models.IssueDB.Update(issue); err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Issue has updated"}
	response.Success(w)

}

// Add issue to active sprint
func AddIssueToSprint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	issueID, err := gocql.ParseUUID(vars["issue_id"])
	if err != nil {
		log.Printf("Error occured in controllers/issue.go while parsing issue_id, method: AddIssueToSprint, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing issue_id, metod: UpdateIssue, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	sprintID, err := gocql.ParseUUID(vars["sprint_id"])
	if err != nil {
		log.Printf("Error occured in controllers/issue.go while parsing issue_id, method: AddIssueToSprint, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go while parsing issue_id, metod: UpdateIssue, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID
	if err := models.IssueDB.FindByID(issue); err != nil {
		log.Printf("Error occured in controllers/issue.go method: AddIssueToSprint, where: issue.FindByID, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: AddIssueToSprint, where: issue.FindByID, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	sprint := models.Sprint{}
	sprint.ID = sprintID
	if err := models.SprintDB.FindByID(&sprint); err != nil {
		log.Printf("Error occured in controllers/issue.go method: AddIssueToSprint, where: sprint.FindById, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: AddIssueToSprint, where: issue.FindById, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	if err := models.IssueDB.Delete(issue); err != nil {
		log.Printf("Error occured in controllers/issue.go method: DeleteIssue, where: issue.Delete, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: DeleteIssue, where: issue.Delete, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	issue.SprintID = sprint.ID
	issue.UpdatedAt = time.Now()
	if err = models.IssueDB.Update(issue); err != nil {
		log.Printf("Error occured in controllers/issue.go method: AddIssueToSprint, where: issue.Update, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: AddIssueToSprint, where: issue.Update, error: %s", err.Error())}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Issue has updated"}
	response.Success(w)
}

//DeleteIssue controller deletes issue from database
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	issueID, err := gocql.ParseUUID(vars["issue_id"])

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = issueID

	if err := models.IssueDB.Delete(issue); err != nil {
		log.Printf("Error occured in controllers/issue.go method: DeleteIssue, where: issue.FindByID, error: %s", err.Error())
		response := helpers.Response{Message: fmt.Sprintf("Error occured in controllers/issue.go metod: UpdateIssue, where: issue.FindByID, error: %s", err.Error())}
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
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	issue := models.Issue{}
	issue.BoardID = id

	boardIssueList, err := models.IssueDB.GetBoardBacklogIssuesList(&issue)
	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Done", Data: boardIssueList}
	response.Success(w)
}

//SprintIssueslist returns list of issues order by sprint_id
func SprintIssueslist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := gocql.ParseUUID(vars["sprint_id"])
	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	issue := models.Issue{}
	issue.SprintID = id

	sprintIssueList, err := models.IssueDB.GetSprintIssueList(&issue)

	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusInternalServerError}
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
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	issue := &models.Issue{}
	issue.UUID = id
	if err := models.IssueDB.FindByID(issue); err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/issue.go error: %+v", err)}
		response.Failed(w)
		return
	}

	jsonResponse, err := json.Marshal(issue)
	if err != nil {
		log.Printf("Error in controllers/issue error: %+v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func SetParentIssue(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Caught error (controllers/issue.SetParentIssue): %+v", err)

		response := helpers.Response{
			StatusCode: http.StatusInternalServerError,
		}

		response.Failed(w)
		return
	}

	req := make([]struct {
		Child  gocql.UUID
		Parent gocql.UUID
	}, 0)

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("Caught error (controllers/issue.SetParentIssue): %+v", err)

		response := helpers.Response{
			StatusCode: http.StatusInternalServerError,
		}

		response.Failed(w)
		return
	}

	for _, value := range req {
		err := models.IssueDB.SetParentIssue(value.Child, value.Parent)

		if err != nil {
			response := helpers.Response{
				StatusCode: http.StatusInternalServerError,
			}

			response.Failed(w)
			return

		}
	}
}

func AddIssueLog(w http.ResponseWriter, r *http.Request) {

	// TODO: move to decodeAndValidate
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Caught error (controllers/issue.AddIssueLog): %+v", err)

		response := helpers.Response{
			StatusCode: http.StatusInternalServerError,
		}

		response.Failed(w)
		return
	}

	req := new(struct {
		IssueID gocql.UUID `json:"issueID"`
		UserID  gocql.UUID `json:"userID"`
		Log     string     `json:"log"`
	})

	err = json.Unmarshal(body, &req)

	if err != nil {
		log.Printf("Caught error (controllers/issue.AddIssueLog): %+v", err)

		response := helpers.Response{
			StatusCode: http.StatusInternalServerError,
		}

		response.Failed(w)
		return
	}

	//newLog := make(map[gocql.UUID]string)
	//newLog[req.UserID] = req.Log

	issue := &models.Issue{}
	issue.UUID = req.IssueID

	err = models.IssueDB.FindByID(issue) // todo ERROR handler
	if err != nil {
		fmt.Println(err)
	}

	issue.Logs = string(body)

	err = models.IssueDB.AddLog(issue)
	if err != nil {
		response := helpers.Response{
			StatusCode: http.StatusInternalServerError,
		}

		response.Failed(w)
		return
	}
}
