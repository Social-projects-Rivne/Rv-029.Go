package main

import (
	"fmt"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
)

type Test struct {
	Name string
	Email string
	Id int
}

func (t Test) GetClaims() (map[string]interface{})  {
	claimsMap := map[string]interface{} {
		"email": t.Email,
		"id": t.Id,
	}

	return claimsMap
}

func main() {
	t := Test{
		Name: "Roman",
		Email: "Email",
		Id: 1,
	}
	jwtToken, _ := jwt.GenerateToken(t)
	fmt.Printf("%+v\n", t.GetClaims())
	fmt.Printf("%+v\n", jwtToken)
}
