package scrum_poker

import (
	"github.com/gorilla/websocket"
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"

	"fmt"
)

func GetClients(req map[string]interface{}, conn *websocket.Conn)  {

	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
fmt.Println(ActiveHubs[issueUUID])


	if err != nil {
		conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST`,
			Message: `invalid issue id`,
		})
		return
	}

	issue := models.Issue{
		UUID: issueUUID,
	}

	err = models.IssueDB.FindByID(&issue)
	if err != nil {
		conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST`,
			Message: `issue not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID]; ok {
		fmt.Println(ActiveHubs[issueUUID].Clients)
		conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST`,
			Message: `Hello guest`,
			Data: "asdfsdf",
		})
		return
	}



}