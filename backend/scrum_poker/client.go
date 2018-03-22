package scrum_poker

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte
	user *models.User
}

func RegisterClient(req map[string]interface{}, conn *websocket.Conn) {

	client := Client{}
	client.conn = conn
	client.send = make(chan []byte, 256)

	sprintUUID, _ := gocql.ParseUUID(req["sprintID"].(string))

	userUUID, _ := gocql.ParseUUID(req["userID"].(string))
	user := &models.User{
		UUID: userUUID,
	}

	err := models.UserDB.FindByID(user)

	if err != nil {
		// TODO: error
	}

	client.user = user

	if _, ok := ActiveHubs[sprintUUID]; ok {
		client.hub = ActiveHubs[sprintUUID]
		client.hub.Register <- &client
	}

	go client.WriteWorker()
}

func SendEstimation(req string){
	strTime := strconv.Itoa(int(time.Now().Unix()))
	producerMessage := &sarama.ProducerMessage{
		Topic: "test-topic-1",
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(req),
	}

	_, _, err := Producer.SendMessage(producerMessage)

	if err != nil {
		fmt.Println(err)
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


