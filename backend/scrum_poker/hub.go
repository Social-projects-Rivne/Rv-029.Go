package scrum_poker

import (
	"github.com/Shopify/sarama"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
)

type Hub struct {
	Producer sarama.SyncProducer
	Consumer sarama.PartitionConsumer
	Clients map[gocql.UUID]*Client
	Register chan *Client
	Unregister chan *Client
	Sprint models.Sprint
	Broadcast chan []byte
}

func newHub(producer sarama.SyncProducer, consumer sarama.PartitionConsumer, sprint models.Sprint) Hub {
	return Hub{
		Producer: producer,
		Consumer: consumer,
		Clients:    make(map[gocql.UUID]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Sprint: sprint,
	}
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
