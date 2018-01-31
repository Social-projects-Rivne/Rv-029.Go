package models

import (
	"reflect"
	"fmt"
	"github.com/relops/cqlr"
	"log"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
)



type BaseModel struct{
	Fields string
	Pointers string
	Condition string
}

func (b *BaseModel) GetFields(structure interface{}) {


	s := reflect.ValueOf(structure).Type()

	for i := 0; i < s.NumField(); i++ {
		if i == s.NumField()-1 {
			b.Fields += fmt.Sprintf("%s", s.Field(i).Tag.Get("cql"))
			b.Pointers += "?"
		} else {
			b.Fields += fmt.Sprintf("%s,", s.Field(i).Tag.Get("cql"))
			b.Pointers += "?,"
		}
	}


}

func (b *BaseModel) Insert(table string,structure interface{}) {
	b.GetFields(structure)
	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)",table,b.Fields,b.Pointers)

	bind := cqlr.Bind(query, structure)
	if err := bind.Exec(db.Session); err != nil {
		log.Fatal(err)
	}

}

func (b *BaseModel) UpdateHelper(structure interface{}) {


	s := reflect.ValueOf(structure).Type()

	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Offset == 0{
			continue
		}
		if i == s.NumField()-1 {
			b.Fields += fmt.Sprintf("%s = ? ", s.Field(i).Tag.Get("cql"))
		} else {
			b.Fields += fmt.Sprintf("%s = ? , ", s.Field(i).Tag.Get("cql"))
		}
	}


}


func (b *BaseModel) Update(table string,structure interface{}) {

	b.UpdateHelper(structure)

	query := fmt.Sprintf("UPDATE %v SET  ",table) + b.Fields + b.Condition
	fmt.Print(query)

	bind := cqlr.Bind(query, structure)
	if err := bind.Exec(db.Session); err != nil {
		log.Fatal(err)
	}
	b.Condition = ""

}

func (b *BaseModel) Where( column string, sign string ,value interface{} ) {

	b.Condition = "WHERE " + column + sign
	b.Condition += fmt.Sprintf("%v",value)


}

func (b *BaseModel) AndWhere( column string, sign string ,value interface{} ) {

	b.Condition = "AND " + column + sign
	b.Condition += fmt.Sprintf("%v",value)


}


