package scrum_poker

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte
}

func RegisterClient(res map[string]interface{}, conn *websocket.Conn) {

	client := Client{}
	client.conn = conn
	client.send = make(chan []byte, 256)

	sprintUUID, _ := gocql.ParseUUID(res["sprintID"].(string))

	for _, hub := range ActiveEstimations {
		if hub.Sprint.ID == sprintUUID {
			client.hub = &hub
			client.hub.Register <- &client
		}
	}

	go client.ReadWorker()
	go client.WriteWorker()
}

func (c *Client) ReadWorker() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		req := make(map[string]interface{}, 0)
		json.Unmarshal(msg, &req)

		message := req["message"].(string)

		strTime := strconv.Itoa(int(time.Now().Unix()))
		producerMessage := &sarama.ProducerMessage{
			Partition: 0,
			Topic: "test-topic-1", // move to argument
			Key:   sarama.StringEncoder(strTime),
			Value: sarama.StringEncoder(message),
		}

		_, _, err = c.hub.Producer.SendMessage(producerMessage)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Client) WriteWorker() {
	const (
		writeWait = 10 * time.Second
	)

	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			err = w.Close()
			if err != nil { return }

		}
	}
}


