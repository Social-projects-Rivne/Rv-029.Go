package scrum_poker

import (
	"encoding/json"
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

var ActiveHubs = make(map[gocql.UUID]*Hub, 0)

func SocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)

	go func() {
		for {
			_, msg, _ := conn.ReadMessage()

			req := make(map[string]interface{}, 0)
			json.Unmarshal(msg, &req)

			switch req["action"] {
			case "CREATE_ESTIMATION_ROOM":
				RegisterHub(req)
			case "REGISTER_CLIENT":
				RegisterClient(req, conn)
			case "ESTIMATION":
				SendEstimation(string(msg))
			}
		}
	}()
}
