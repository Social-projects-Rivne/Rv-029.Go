package models

import (
	"encoding/json"
	"net/http"
)

//go:generate mockgen -destination=../mocks/mock_request.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/models Request

func init() {
	InitBoardRequest(&Decoder{})
}

type Request interface {
	Decode(*http.Request, interface{}) (interface{}, error)
}

type Decoder struct{}

func (d *Decoder) Decode(r *http.Request, i interface{}) (interface{}, error) {
	err := json.NewDecoder(r.Body).Decode(i)

	defer r.Body.Close()

	if err != nil {
		return nil, err
	}

	return i, nil
}

var BoardRequest Request

func InitBoardRequest(req Request) {
	BoardRequest = req
}
