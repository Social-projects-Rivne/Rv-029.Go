package scrum_poker

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"fmt"
)

type Hub struct {
	Clients map[gocql.UUID]*Client
	Register chan *Client
	Unregister chan *Client
	Broadcast chan []byte
	Sprint models.Sprint
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

func newHub(sprint models.Sprint) Hub {
	return Hub{
		Clients:    make(map[gocql.UUID]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan []byte),
		Sprint: sprint,
	}
}

func RegisterHub(req map[string]interface{}, conn *websocket.Conn) {
	sprintUUID, err := gocql.ParseUUID(req["sprintID"].(string))
	if err != nil {
		conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `CREATE_ESTIMATION_ROOM`,
			Message: `invalid sprint id`,
		});
		return
	}
	
	sprint := models.Sprint{
		ID: sprintUUID,
	}
	
	err = models.SprintDB.FindByID(&sprint)
	if err != nil {
		conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `CREATE_ESTIMATION_ROOM`,
			Message: `sprint not found`,
		});
		return
	}

	if _, ok := ActiveHubs[sprintUUID]; ok {
		conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `CREATE_ESTIMATION_ROOM`,
			Message: `room already exists`,
		});
		return
	}

	hub := newHub(sprint)

	ActiveHubs[sprintUUID] = &hub

	go hub.run()

	conn.WriteJSON(SocketResponse{
		Status: true,
		Action: `CREATE_ESTIMATION_ROOM`,
		Message: `room was successfully created`,
	});
	return
}

func (h *Hub) run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client.user.UUID] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.user.UUID]; ok {
				delete(h.Clients, client.user.UUID)
				close(client.send)
				if len(h.Clients) == 0 {
					delete(ActiveHubs, h.Sprint.ID)
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
