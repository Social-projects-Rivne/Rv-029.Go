package scrum_poker

import (
	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"encoding/json"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	user *models.User
}

func RegisterClient(req map[string]interface{}, client *Client) {

	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `REGISTER_CLIENT`,
			Message: `invalid issue id`,
		})
		return
	}

	if hub, ok := ActiveHubs[issueUUID]; ok {
		hub = ActiveHubs[issueUUID]
		hub.Register <- client
	} else {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `REGISTER_CLIENT`,
			Message: `room not found`,
		})
		return
	}

	client.conn.WriteJSON(SocketResponse{
		Status: true,
		Action: `REGISTER_CLIENT`,
		Message: `you was successfully connected to the estimation room`,
	})
	return
}

func SendEstimation(req map[string]interface{}, client *Client){
	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `ESTIMATION`,
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
			Action: `ESTIMATION`,
			Message: `issue not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID]; !ok {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `ESTIMATION`,
			Message: `room not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID].Clients[client.user.UUID]; !ok {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `ESTIMATION`,
			Message: `user is not connected to the room`,
		})
		return
	}

	if value, ok := req["estimate"]; !ok || value.(int) < 0 || value.(int) > 10 {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `ESTIMATION`,
			Message: `estimation is not set or have invalid value`,
		})
		return
	}

	strTime := strconv.Itoa(int(time.Now().Unix()))
	jsonVal, err := json.Marshal(req)
	if err != nil {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `ESTIMATION`,
			Message: `invalid json encoding`,
		})
		return
	}

	producerMessage := &sarama.ProducerMessage{
		Topic: "test-topic-1",
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(string(jsonVal)),
	}

	_, _, err = Producer.SendMessage(producerMessage)
	if err != nil {
		client.conn.WriteJSON(SocketResponse{
			Status: false,
			Action: `ESTIMATION`,
			Message: `estimation was not saved`,
		})
		return
	}

	client.conn.WriteJSON(SocketResponse{
		Status: true,
		Action: `ESTIMATION`,
		Message: `your estimate was successfully saved`,
	})
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


