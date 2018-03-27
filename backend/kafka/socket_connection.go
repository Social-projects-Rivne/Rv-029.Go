package kafka

import (
	//"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"fmt"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Estimation struct {
	Partition int32
	Sprint models.Sprint
	Users []models.User
	Producer *sarama.SyncProducer
}

var activeEstimations []Estimation

func SocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)

	response := make(map[string]interface{}, 0)

	_, msg, _ := conn.ReadMessage()

	json.Unmarshal(msg, &response)

	switch response["action"] {
	case "CREATE_ESTIMATION_ROOM":
		//todo: handle errors
		sprintUUID, _ := gocql.ParseUUID(response["sprintID"].(string))
		sprint := models.Sprint{
			ID: sprintUUID,
		}
		_ = models.SprintDB.FindByID(&sprint)
		userUUID, _ := gocql.ParseUUID(response["userID"].(string))
		user := models.User{
			UUID: userUUID,
		}
		_ = models.UserDB.FindByID(&user)

		activeEstimations = append(activeEstimations, Estimation{
			Partition: 11,
			Sprint: sprint,
			Users: []models.User{user},
			Producer: InitProducer(),
		})

		fmt.Println(activeEstimations)
	case "CONNECT_TO_ESTIMATION_ROOM":
		sprintUUID, _ := gocql.ParseUUID(response["sprintID"].(string))

		found := false

		for _, estimation := range activeEstimations {
			if estimation.Sprint.ID == sprintUUID {
				estimation.Users = append(estimation.Users, models.User{})//todo: find user
				found = true
			}
		}

		if found {
			//send success
		} else {
			//error
		}
		InitConsumer(conn, "test-topic-1", 0)
	case "ESTIMATION_MESSAGE":

	}


	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//go func() {
	//	InitConsumer(conn, "test-topic-1", 0)
	//}()
	//
	//go func() { RunProducer(conn, w) }()
}
