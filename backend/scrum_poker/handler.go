package scrum_poker

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/dgrijalva/jwt-go"
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
	Status  bool        `json:"status"`
	Action  string      `json:"action,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var ActiveHubs = make(map[gocql.UUID]*Hub, 0)
var ConnectedUsers = make(map[gocql.UUID]*Client, 0)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	var client *Client
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
				Status:  false,
				Action:  `CONNECTION`,
				Message: `invalid token`,
			})
		}

		user := models.User{
			UUID: userUUID,
		}
		err = models.UserDB.FindByID(&user)
		if err != nil {
			conn.WriteJSON(SocketResponse{
				Status:  false,
				Action:  `CONNECTION`,
				Message: `user not found`,
			})
		}

		client = &Client{
			conn: conn,
			//send: make(chan []byte, 256),
			user: &user,
		}

		//go client.WriteWorker()

		ConnectedUsers[user.UUID] = client

		conn.WriteJSON(SocketResponse{
			Status:  true,
			Action:  `CONNECTION`,
			Message: `you was successfully connected to the socket server`,
		})

	}

	// on disconnect remove from hub clients
	defer func() {
		for _, hub := range ActiveHubs {
			if len(hub.Clients) > 0 {
				for userID, _ := range hub.Clients {
					if reflect.DeepEqual(userID, client.user.UUID) {
						hub.Calculate()
						hub.Unregister <- client

						if len(hub.Guests) > 0 {
							hub.BroadcastGusets <- &SocketResponse{
								Status:  true,
								Action:  `USER_DISCONNECT_FROM_ROOM`,
								Message: `user disconnected from the room`,
								Data: client.user,
							}
						}
					}
				}
			}
			if len(hub.Guests) > 0 {
				for userID, _ := range hub.Guests {
					if reflect.DeepEqual(userID, client.user.UUID) {
						hub.Calculate()
						hub.UnregisterGuest <- client
					}
				}
			}
		}

		// on disconnect remove from ConnectedUsers
		delete(ConnectedUsers, client.user.UUID)
	}()

	for {
		_, msg, _ := conn.ReadMessage()

		req := make(map[string]interface{}, 0)
		json.Unmarshal(msg, &req)

		switch req["action"] {
		case "CREATE_ESTIMATION_ROOM":
			RegisterHub(req, client)
		case "REGISTER_CLIENT":
			RegisterClient(req, client)
		case "ESTIMATION":
			SendEstimation(req, client)
		case "GUEST":
			GetClients(req, client)
		}
	}
}
