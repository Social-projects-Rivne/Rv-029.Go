package scrum_poker

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
)

type Hub struct {
	Clients map[gocql.UUID]*Client
	Register chan *Client
	Unregister chan *Client
	Broadcast chan []byte
	Sprint models.Sprint
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

func RegisterHub(req map[string]interface{}) {
	sprintUUID, _ := gocql.ParseUUID(req["sprintID"].(string))
	sprint := models.Sprint{
		ID: sprintUUID,
	}
	_ = models.SprintDB.FindByID(&sprint)

	hub := newHub(sprint)

	ActiveHubs[sprintUUID] = &hub

	go hub.run()
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
