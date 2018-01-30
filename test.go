package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gocql/gocql"
)

type User struct {
	id         gocql.UUID
	email      string
	first_name string
	last_name  string
	password   string
	salt       string
	role       string
	created_at time.Time
	updated_at time.Time
}

// type T struct {
// 	A int
// 	B string
// }

// if err := db.Session.Query(`INSERT INTO users (id,email,first_name,last_name,password,salt,role,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?);	`,
// 		gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Password, user.Salt, user.Role, user.CreatedAt, user.UpdatedAt).Exec(); err != nil {
// 		fmt.Println(err)
// 	}

func main() {

	t := User{gocql.TimeUUID(), "email1", "fname1", "lname1", "passwd", "salt", "role", time.Now(), time.Now()}
	s := reflect.ValueOf(&t).Elem()
	typeOfStruct := s.Type()

	str := "INSERT INTO Users ("

	for i := 0; i < s.NumField(); i++ {
		if i == s.NumField()-1 {
			str += fmt.Sprintf("%s) VALUES (", typeOfStruct.Field(i).Name)
		} else {
			str += fmt.Sprintf("%s,", typeOfStruct.Field(i).Name)
		}
	}

	v := reflect.ValueOf(t)
	fmt.Println(v)
	fmt.Println(t.created_at)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i)
		switch{
			case
		}
	}



	fmt.Println(values)
	fmt.Println(str)


	asd := fmt.Sprintf("%v",time.Now())
	fmt.Println(asd)

}
