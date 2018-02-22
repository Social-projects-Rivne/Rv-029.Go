package models

import (
	"github.com/gocql/gocql"
	"log"
	//"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)

type IStore interface {
	Delete(board *Board) error
}

type Store struct {
	db *gocql.Session
}

var CStore IStore

func (s *Store) Delete(board *Board) error {
	err := s.db.Query(`DELETE FROM boards where id = ?;`, board.ID).Exec()

	if err != nil {
		log.Printf("Error in method Delete inside models/board.go: %s\n", err.Error())
		return err
	}

	return nil
}

func InitStore(s IStore)  {
	CStore = s
}