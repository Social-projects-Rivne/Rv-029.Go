package seeders

import (
	"fmt"
	"reflect"
)

type Seeder interface {
	Run()
}

func Call(class Seeder) {
	class.Run()
	fmt.Println(reflect.TypeOf(class).Name())
}
