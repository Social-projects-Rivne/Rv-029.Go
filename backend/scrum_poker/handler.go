package scrum_poker

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketResponse struct {
	Status bool `json:"status"`
	Action string `json:"action,omitempty"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

var ActiveHubs = make(map[gocql.UUID]*Hub, 0)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	defer func() {
		for _, hub := range ActiveHubs {
			for _, client := range hub.Clients {
				if reflect.DeepEqual(&client.conn, &conn) {
					hub.Unregister <- client
				}
			}
		}
	}()

	for {
		_, msg, _ := conn.ReadMessage()

		req := make(map[string]interface{}, 0)
		json.Unmarshal(msg, &req)

		switch req["action"] {
		case "CREATE_ESTIMATION_ROOM":
			RegisterHub(req, conn)
		case "REGISTER_CLIENT":
			RegisterClient(req, conn)
		case "ESTIMATION":
			SendEstimation(req, conn)
		}
	}
}
