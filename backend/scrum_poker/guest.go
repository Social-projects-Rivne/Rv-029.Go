package scrum_poker

import (
	"github.com/gocql/gocql"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"

	"fmt"
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

	if _, ok := ActiveHubs[issueUUID]; ok {
		for k, v := range ActiveHubs[issueUUID].Clients {
			fmt.Println("££££££££££££££££££££££££££")
			fmt.Printf("key[%s] value[%s]\n", k, v.user)
			fmt.Println("££££££££££££££££££££££££££")



		}
		//fmt.Println(ActiveHubs[issueUUID].Clients)
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `GUEST`,
			Message: `Hello guest`,
			Data: "asdfsdf",
		})
		return
	}



}