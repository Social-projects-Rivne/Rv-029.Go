package models

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/db"
	"github.com/relops/cqlr"
)

const PRIMERY = "primery"

type BaseModel struct {
	Fields    string
	Pointers  string
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

func (b *BaseModel) Insert(table string, structure interface{}) {
	b.GetFields(structure)
	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", table, b.Fields, b.Pointers)

	bind := cqlr.Bind(query, structure)
	if err := bind.Exec(db.GetInstance().Session); err != nil {
		log.Fatal(err)
	}

}

func (b *BaseModel) UpdateHelper(structure interface{}) {

	s := reflect.ValueOf(structure)
	typeOfS := s.Type()

	var emptyTime time.Time

	var fields []string

	for i := 0; i < s.NumField(); i++ {

		if strings.ToLower(typeOfS.Field(i).Tag.Get("key")) == PRIMERY || s.Field(i).Interface() == emptyTime || s.Field(i).Interface() == 0  || s.Field(i).Interface() == ""{
			continue
		}

		fields = append(fields, fmt.Sprintf("%s = ? ", typeOfS.Field(i).Tag.Get("cql")))
		// if i == s.NumField() - 1 {
			// b.Fields += fmt.Sprintf("%s = ? ", typeOfS.Field(i).Tag.Get("cql"))
		// } else {
			// b.Fields += fmt.Sprintf("%s = ? , ", typeOfS.Field(i).Tag.Get("cql"))
		// }
	}

	b.Fields += strings.Join(fields, ", ")
}

func (b *BaseModel) Update(table string, structure interface{}) {

	b.UpdateHelper(structure)

	query := fmt.Sprintf("UPDATE %v SET  ", table) + b.Fields + b.Condition
	fmt.Println(query)
	bind := cqlr.Bind(query, structure)

	fmt.Println(b.Fields)
	fmt.Println(b.Condition)
	fmt.Println(structure)

	if err := bind.Exec(db.Session); err != nil {
		log.Fatal(err)
		fmt.Println("message")
	}
	b.Condition = ""

}

func (b *BaseModel) Where(column string, sign string, value interface{}) {

	b.Condition = " WHERE " + column + sign
	b.Condition += fmt.Sprintf("%v",value)


}

func (b *BaseModel) AndWhere(column string, sign string, value interface{}) {

	b.Condition += " AND " + column + sign
	b.Condition += fmt.Sprintf("%v",value)

	user := &User{}

	for c.Scan(&user) {
		log.Println(user.UUID)
		fmt.Println(user)
	}
}

//FindIssue finds issue by any field
func (b *BaseModel) FindIssue() *Issue {

	query := fmt.Sprintf("SELECT * FROM issues %v ", b.Condition)

	q := db.GetInstance().Session.Query(query)
	c := cqlr.BindQuery(q)

	issue := &Issue{}

	for c.Scan(&issue) {
		log.Println(issue.UUID)
		return issue
	}
	return nil
}
