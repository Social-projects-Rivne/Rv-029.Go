package scrum_poker

import (
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
)

func GetClients(req map[string]interface{}, client *Client)  {

	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.conn.WriteJSON(SocketResponse{
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
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST`,
			Message: `issue not found`,
		})
		return
	}

	if hub, ok := ActiveHubs[issueUUID]; ok {

		hub = ActiveHubs[issueUUID]
		hub.RegisterGuest <- client

		users := make(map[gocql.UUID]*models.User, 0)
		for v, _ := range ActiveHubs[issueUUID].Clients {
			users[v.user.UUID] = v.user
		}
		client.conn.WriteJSON(SocketResponse{
			Status: true,
			Action: `GUEST`,
			Message: `Welcome in estimation room`,
			Data: users,
		})
		return
	}
}
