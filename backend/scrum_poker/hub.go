package scrum_poker

import (
	"github.com/Shopify/sarama"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"log"
)

type Hub struct {
	Producer sarama.SyncProducer
	Consumer sarama.PartitionConsumer
	Clients map[*Client]bool
	Register chan *Client
	Unregister chan *Client
	Sprint models.Sprint
}

func newHub(producer sarama.SyncProducer, consumer sarama.PartitionConsumer, sprint models.Sprint) Hub {
	return Hub{
		Producer: producer,
		Consumer: consumer,
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Sprint: sprint,
	}
}

func (h *Hub) run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}

		case err := <-h.Consumer.Errors():
			log.Fatal(err)

		case m := <-h.Consumer.Messages():
			message := []byte(m.Value)

			for client := range h.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
