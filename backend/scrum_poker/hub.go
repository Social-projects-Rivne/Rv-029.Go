package scrum_poker

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
)

type Hub struct {
	Clients    map[gocql.UUID]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *SocketResponse
	Issue      models.Issue
	Summary    map[gocql.UUID]int
	Results    map[int]float32
}

func (h *Hub) Calculate() {
	marks := make(map[int]float32, 0)
	for _, mark := range h.Summary {
		marks[mark]++
	}

	for mark, count := range marks {
		h.Results[mark] = count / float32(len(h.Summary))
	}

	if len(h.Summary) >= len(h.Clients) {
		h.Broadcast <- &SocketResponse{
			Status:  true,
			Action:  `ESTIMATION_RESULTS`,
			Message: `estimation completed`,
			Data: struct {
				Summary map[gocql.UUID]int `json:"summary"`
				Results map[int]float32    `json:"results"`
			}{
				h.Summary,
				h.Results,
			},
		}
	}
}

func newHub(issue models.Issue) Hub {
	return Hub{
		Clients:    make(map[gocql.UUID]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *SocketResponse),
		Issue:      issue,
		Summary:    make(map[gocql.UUID]int, 0),
		Results:    make(map[int]float32, 0),
	}
}

func RegisterHub(req map[string]interface{}, client *Client) {
	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.send(SocketResponse{
			Status:  false,
			Action:  `CREATE_ESTIMATION_ROOM`,
			Message: `invalid issue id`,
		})
		return
	}

	issue := models.Issue{
		UUID: issueUUID,
	}
	err = models.IssueDB.FindByID(&issue)
	if err != nil {
		client.send(SocketResponse{
			Status:  false,
			Action:  `CREATE_ESTIMATION_ROOM`,
			Message: `issue not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID]; ok {
		client.send(SocketResponse{
			Status:  false,
			Action:  `CREATE_ESTIMATION_ROOM`,
			Message: `room already exists`,
		})
		return
	}

	hub := newHub(issue)

	ActiveHubs[issueUUID] = &hub

	go hub.run()

	client.send(SocketResponse{
		Status:  true,
		Action:  `CREATE_ESTIMATION_ROOM`,
		Message: `room was successfully created`,
	})
	return
}

func (h *Hub) run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client.user.UUID] = client

			h.Broadcast <- &SocketResponse{
				Status:  true,
				Action:  `NEW_USER_IN_ROOM`,
				Message: `new user connected to the room`,
				Data: client.user,
			}
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.user.UUID]; ok {
				delete(h.Clients, client.user.UUID)
				//close(client.send)
			}

			if len(h.Clients) == 0 {
				delete(ActiveHubs, h.Issue.UUID)
			}

			h.Broadcast <- &SocketResponse{
				Status:  true,
				Action:  `USER_DISCONNECT_FROM_ROOM`,
				Message: `user disconnected from the room`,
				Data: client.user,
			}
		case msg := <-h.Broadcast:
			for _, client := range h.Clients {
				client.send(msg)
			}
		}
	}
}
