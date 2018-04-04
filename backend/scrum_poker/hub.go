package scrum_poker

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"reflect"
)

const LIMIT = 0.6

type Hub struct {
	Clients           map[*Client]gocql.UUID
	Guests    		  map[gocql.UUID]*Client
	Register          chan *Client
	Unregister        chan *Client
	RegisterGuest     chan *Client
	UnregisterGuest   chan *Client
	Broadcast         chan *SocketResponse
	DirectedBroadcast chan *DirectedResponse
	BroadcastGusets   chan *SocketResponse
	Issue             models.Issue
	Summary           map[gocql.UUID]int
	Results           map[int]float32
}

func (h *Hub) Calculate() {
	var estimate int
	message := "estimation didn't get 60%"

	if len(h.Summary) > 0 {
		marks := make(map[int]float32, 0)
		for _, mark := range h.Summary {
			marks[mark]++
		}

		// count estimation percent of concrete mark
		for mark, countOfPeople := range marks {
			h.Results[mark] = countOfPeople / float32(len(h.Summary))

			if h.Results[mark] >= LIMIT {
				estimate = mark
				message = "estimation completed"
			}

		}

		if len(h.Summary) >= len(h.Clients) && len(h.Clients) > 0 {
			h.Broadcast <- &SocketResponse{
				Status:  true,
				Action:  `ESTIMATION_RESULTS`,
				Message: message,
				Data: struct {
					Summary  map[gocql.UUID]int `json:"summary"`
					Results  map[int]float32    `json:"results"`
					Estimate int                `json:"estimate,omitempty"`
				}{
					h.Summary,
					h.Results,
					estimate,
				},
			}
		}
	}
}

func newHub(issue models.Issue) Hub {
	return Hub{
		Clients:           make(map[*Client]gocql.UUID),
		Guests:    	 	   make(map[gocql.UUID]*Client),
		Register:          make(chan *Client),
		Unregister:        make(chan *Client),
		RegisterGuest:     make(chan *Client),
		UnregisterGuest:   make(chan *Client),
		Broadcast:         make(chan *SocketResponse),
		DirectedBroadcast: make(chan *DirectedResponse),
		BroadcastGusets:   make(chan *SocketResponse),
		Issue:             issue,
		Summary:           make(map[gocql.UUID]int, 0),
		Results:           make(map[int]float32, 0),
	}
}

func RegisterHub(req map[string]interface{}, client *Client) {
	issueUUID, err := gocql.ParseUUID(req["issueID"].(string))
	if err != nil {
		client.send(SocketResponse{
			Status:  false,
			Action:  `CREATE_ESTIMATION_ROOM`,
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
			Action:  `CREATE_ESTIMATION_ROOM`,
			Message: `issue not found`,
		})
		return
	}

	if _, ok := ActiveHubs[issueUUID]; ok {
		client.send(SocketResponse{
			Status:  false,
			Action:  `CREATE_ESTIMATION_ROOM`,
			Message: `room already exists`,
		})
		return
	}

	hub := newHub(issue)

	ActiveHubs[issueUUID] = &hub

	go hub.run()

	client.send(SocketResponse{
		Status:  true,
		Action:  `CREATE_ESTIMATION_ROOM`,
		Message: `room was successfully created`,
	})
	return
}

func (h *Hub) run() {
	for {
		select {


		case guest := <-h.UnregisterGuest:
			if _, ok := h.Guests[guest.user.UUID]; ok {
				delete(h.Guests, guest.user.UUID)
			}
		case client := <-h.Register:
			for onlineClient, userUUID := range ConnectedUsers {
				if client.user.UUID == userUUID {
					h.Clients[onlineClient] = userUUID

					onlineClient.send(SocketResponse{
						Status:  true,
						Action:  `REGISTER_CLIENT`,
						Message: `you was successfully connected to the estimation room`,
					})
				}
			}
		case guest := <-h.RegisterGuest:
			h.Guests[guest.user.UUID] = guest
		case msg := <-h.BroadcastGusets:
			if len(h.Guests) > 0 {
				for _, guest := range h.Guests {
					guest.send(msg)
				}
			}
		case client := <-h.Unregister:
			for hclient, _ := range h.Clients {
				if reflect.DeepEqual(hclient, client) {
					delete(h.Clients, client)

					if len(h.Clients) == 0 && len(h.Guests) == 0 {
						delete(ActiveHubs, h.Issue.UUID)
					}
				}
			}
		case msg := <-h.Broadcast:
			if len(h.Clients) > 0 {
				for client, _ := range h.Clients {
					client.send(msg)
				}
			}
		case msg := <-h.DirectedBroadcast:
			for client, userUUID := range h.Clients {
				if userUUID == msg.UserUUID {
					client.send(msg.SocketResponse)
				}
			}
		}

	}
}
