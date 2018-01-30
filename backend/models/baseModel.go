package models

import (
	"reflect"
	"fmt"
	"github.com/relops/cqlr"
	"log"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)



type baseModel struct{
	Fields string
	Pointers string
}

func (b *baseModel) getFields(structure interface{}) {

	s := reflect.ValueOf(structure).Type()

	for i := 0; i < s.NumField(); i++ {
		if i == s.NumField()-1 {
			b.Fields += fmt.Sprintf("%s", s.Field(i).Name)
			b.Pointers += "?"
		} else {
			b.Fields += fmt.Sprintf("%s,", s.Field(i).Name)
			b.Pointers += "?,"
		}
	}
	fmt.Print(b.Fields)


}

func (b *baseModel) insert(table string,structure interface{}) {
	b.getFields(structure)
	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)",table,b.Fields,b.Pointers)

	bind := cqlr.Bind(query, structure)
	if err := bind.Exec(db.Session); err != nil {
		log.Fatal(err)
	}

}


