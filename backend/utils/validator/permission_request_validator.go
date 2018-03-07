package validator

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"errors"
)

type PermissionRequestData struct {
	*baseValidator
	Permission    string
}

func (p *PermissionRequestData) Validate(r *http.Request) error {
	if stringInSlice(p.Permission, models.GetPermissionsList()) {
		return errors.New("Invalid Permission")
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}