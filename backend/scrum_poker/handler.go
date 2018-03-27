package scrum_poker

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"github.com/dgrijalva/jwt-go"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
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
var ConnectedUsers = make(map[gocql.UUID]*Client, 0)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	jwtToken := r.URL.Query().Get("token")

	var keyFunc jwt.Keyfunc
	keyFunc = func(token *jwt.Token) (interface{}, error) {
		return "SomeSecretKey", nil
	}

	token, _ := jwt.Parse(jwtToken, keyFunc)
	claims := token.Claims

	if userID, ok := claims.(jwt.MapClaims)["UUID"]; ok {
		userUUID, err := gocql.ParseUUID(userID.(string))
		if err != nil {
			conn.WriteJSON(SocketResponse{
				Status: false,
				Action: `CONNECTION`,
				Message: `invalid token`,
			})
		}

		user := models.User{
			UUID: userUUID,
		}
		err = models.UserDB.FindByID(&user)
		if err != nil {
			conn.WriteJSON(SocketResponse{
				Status: false,
				Action: `CONNECTION`,
				Message: `user not found`,
			})
		}

		client := Client{}
		client.conn = conn
		client.send = make(chan []byte, 256)
		client.user = &user
		go client.WriteWorker()

		ConnectedUsers[user.UUID] = &client

		conn.WriteJSON(SocketResponse{
			Status: true,
			Action: `CONNECTION`,
			Message: `you was successfully connected to the socket server`,
		})

		defer func() {
			delete(ConnectedUsers, user.UUID)
		}()
	}

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
