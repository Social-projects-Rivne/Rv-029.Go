package scrum_poker

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"fmt"
)

type Hub struct {
	Clients map[gocql.UUID]*Client
	Register chan *Client
	Unregister chan *Client
	Broadcast chan []byte
	Issue models.Issue
	Summary map[gocql.UUID]map[gocql.UUID]int // map[issueID]map[userID]estimate
	Results map[gocql.UUID]map[int]float32 // map[issueID]map[userID]estimate
}

func (h *Hub) Calculate() {
	for issueUUID, estimations := range h.Summary {
		marks := make(map[int]float32, 0)
		for _, mark := range estimations {
			marks[mark] ++
		}
		for mark, count := range marks {
			marks[mark] = count / float32(len(estimations))
		}
		h.Results[issueUUID] = marks
	}



	fmt.Println("Results:")
	fmt.Printf("%+v\n", h.Results)
	fmt.Println("Summary:")
	fmt.Printf("%+v\n", h.Summary)
	//todo: calculate result of estimation
	//todo: if len(Clients) == len(issue estimations) -> send message with result to broadcast
}

func newHub(issue models.Issue) Hub {
	return Hub{
		Clients:    make(map[gocql.UUID]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan []byte),
		Issue: issue,
	}
}

func RegisterHub(req map[string]interface{}, client *Client) {
	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `CREATE_ESTIMATION_ROOM`,
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
			Action: `CREATE_ESTIMATION_ROOM`,
			Message: `issue not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID]; ok {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `CREATE_ESTIMATION_ROOM`,
			Message: `room already exists`,
		})
		return
	}

	hub := newHub(issue)

	ActiveHubs[issueUUID] = &hub

	go hub.run()

	client.conn.WriteJSON(SocketResponse{
		Status: true,
		Action: `CREATE_ESTIMATION_ROOM`,
		Message: `room was successfully created`,
	})
	return
}

func (h *Hub) run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client.user.UUID] = client
			fmt.Printf("user %+v", h)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.user.UUID]; ok {
				delete(h.Clients, client.user.UUID)
				close(client.send)
				if len(h.Clients) == 0 {
					delete(ActiveHubs, h.Issue.UUID)
				}
			}
		case msg := <-h.Broadcast:
			for _, client := range h.Clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(h.Clients, client.user.UUID)
				}
			}
		}
	}
}
