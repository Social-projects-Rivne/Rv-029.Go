package models

var BoardRequest Request

func InitBoardRequest(req Request) {
	BoardRequest = req
}

type BoardCreateRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
