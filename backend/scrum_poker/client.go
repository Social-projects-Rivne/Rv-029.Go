package scrum_poker

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	conn *websocket.Conn
	//send chan []byte
	user *models.User
	mu   sync.Mutex
}

func (c *Client) send(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteJSON(v)
}

func (c *Client) disconnect() error {
	return c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

func RegisterClient(req map[string]interface{}, client *Client) {

	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.send(SocketResponse{
			Status:  false,
			Action:  `REGISTER_CLIENT`,
			Message: `invalid issue id`,
		})
		return
	}

	if hub, ok := ActiveHubs[issueUUID]; ok {
		hub = ActiveHubs[issueUUID]
		hub.Register <- client

		hub.Broadcast <- &SocketResponse{
			Status:  true,
			Action:  `NEW_USER_IN_ROOM`,
			Message: `new user connected to the room`,
			Data: client.user,
		}
	} else {
		client.send(SocketResponse{
			Status:  false,
			Action:  `REGISTER_CLIENT`,
			Message: `room not found`,
		})
		return
	}

	client.send(SocketResponse{
		Status:  true,
		Action:  `REGISTER_CLIENT`,
		Message: `you was successfully connected to the estimation room`,
	})
	return
}

func SendEstimation(req map[string]interface{}, client *Client) {
	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.send(SocketResponse{
			Status:  false,
			Action:  `ESTIMATION`,
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
			Action:  `ESTIMATION`,
			Message: `issue not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID]; !ok {
		client.send(SocketResponse{
			Status:  false,
			Action:  `ESTIMATION`,
			Message: `room not found`,
		})
		return
	}


	if _, ok := ActiveHubs[issueUUID].Clients[client]; !ok {
		client.send(SocketResponse{
			Status:  false,
			Action:  `ESTIMATION`,
			Message: `user is not connected to the room`,
		})
		return
	}

	var estimate int
	if value, ok := req["estimate"]; !ok {
		client.send(SocketResponse{
			Status:  false,
			Action:  `ESTIMATION`,
			Message: `estimation is not set`,
		})
		return
	} else {
		estimate, err = strconv.Atoi(value.(string))
		if err != nil || estimate < 0 || estimate > 10 {
			client.send(SocketResponse{
				Status:  false,
				Action:  `ESTIMATION`,
				Message: `estimation has invalid value`,
			})
			return
		}
	}

	strTime := strconv.Itoa(int(time.Now().Unix()))
	jsonVal, err := json.Marshal(struct {
		Action    string     `json:"action"`
		UserUUID  gocql.UUID `json:"user_uuid"`
		IssueUUID gocql.UUID `json:"issue_uuid"`
		Estimate  int        `json:"estimate"`
		CreatedAt string     `json:"created_at"`
	}{
		Action:    req["action"].(string),
		UserUUID:  client.user.UUID,
		IssueUUID: issueUUID,
		Estimate:  estimate,
		CreatedAt: time.Now().Format("2006-01-02 03:04:05"),
	})

	if err != nil {
		client.send(SocketResponse{
			Status:  false,
			Action:  `ESTIMATION`,
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
		client.send(SocketResponse{
			Status:  false,
			Action:  `ESTIMATION`,
			Message: `estimation was not saved`,
		})
		return
	}

	client.send(SocketResponse{
		Status:  true,
		Action:  `ESTIMATION`,
		Message: `your estimate was successfully saved`,
	})
}

//
//func (c *Client) WriteWorker() {
//	const (
//		writeWait = 10 * time.Second
//	)
//
//	defer func() {
//		c.conn.Close()
//	}()
//
//	for {
//		select {
//		case message, ok := <-c.send:
//			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
//			if !ok {
//				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
//				return
//			}
//
//			w, err := c.conn.NextWriter(websocket.TextMessage)
//			if err != nil {
//				return
//			}
//
//			w.Write(message)
//
//			n := len(c.send)
//			for i := 0; i < n; i++ {
//				w.Write([]byte{'\n'})
//				w.Write(<-c.send)
//			}
//
//			err = w.Close()
//			if err != nil {
//				return
//			}
//		}
//	}
//}
