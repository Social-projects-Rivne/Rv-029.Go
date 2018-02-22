package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"net/http"
	"github.com/gocql/gocql"
	"fmt"
)

// decodeAndValidate - entry point for deserialization and validation
// of the submission input
func decodeAndValidate(r *http.Request, v validator.InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	// perform validation on the InputValidation implementation
	return v.Validate(r)
}

type CRUD interface {
	Insert() error
	Update() error
	Delete() error
	FindByID() error
	List(gocql.UUID) ([]map[string]interface{}, error)
}

type NewModel struct {
	Model CRUD
}

func Delete(c CRUD) error {
	fmt.Println("3#######################")
	return c.Delete()
}

func (n *NewModel) Insert() error {
	return n.Model.Insert()
}

func (n *NewModel) Update() error {
	return n.Model.Update()
}

func (n *NewModel) Delete() error {
	fmt.Println("123")
	return n.Model.Delete()
}

func (n *NewModel) FindByID() error {
	return n.Model.FindByID()
}

func (n *NewModel) List(g gocql.UUID) ([]map[string]interface{}, error) {
	return n.Model.List(g)
}
