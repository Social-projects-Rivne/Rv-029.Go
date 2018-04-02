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
		hub.Register <- client

		users := make([]*models.User, 0)
		for _, v := range ActiveHubs[issueUUID].Clients {
			users = append(users, v.user)
		}
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST`,
			Message: `Hello guest`,
			Data: users,
		})
		return
	}



}

func RegisterGuest (req map[string]interface{}, client *Client){
	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))

	if err != nil {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST_REGISTER`,
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
			Action: `GUEST_REGISTER`,
			Message: `issue not found`,
		})
		return
	}

	if hub, ok := ActiveHubs[issueUUID]; ok {
		hub = ActiveHubs[issueUUID]
		hub.RegisterGuest <- client

		client.send(SocketResponse{
			Status:  false,
			Action:  `GUEST_REGISTER`,
			Message: `qqqqqqq`,
		})
	} else {
		client.send(SocketResponse{
			Status:  false,
			Action:  `GUEST_REGISTER`,
			Message: `room not found`,
		})
		return
	}

}