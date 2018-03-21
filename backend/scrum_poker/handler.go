package scrum_poker

import (
	"encoding/json"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ActiveEstimations = make([]Hub, 0)

func SocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)

	go func() {
		for {
			req := make(map[string]interface{}, 0)
			_, msg, _ := conn.ReadMessage()
			json.Unmarshal(msg, &req)

			switch req["action"] {
			case "CREATE_ESTIMATION_ROOM":
				CreateRoom(req)
			case "REGISTER_CLIENT":
				RegisterClient(req, conn)
			case "ESTIMATION":
				fmt.Println(1)
			}
		}
	}()

}

func CreateRoom(req map[string]interface{}) {

	sprintUUID, _ := gocql.ParseUUID(req["sprintID"].(string))
	sprint := models.Sprint{
		ID: sprintUUID,
	}
	_ = models.SprintDB.FindByID(&sprint)

	// One producer & one consumer per room
	producer := InitProducer()
	consumer := InitConsumer("test-topic-1", 0)

	hub := newHub(producer, consumer, sprint)

	ActiveEstimations = append(ActiveEstimations, hub)

	go hub.run()

}
