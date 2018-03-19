package validator

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"errors"
	"fmt"
)

type SetPermissionsRequestData struct {
	*baseValidator
	Permissions    []string
}

func (p *SetPermissionsRequestData) Validate(r *http.Request) error {
	for _, permission := range p.Permissions {
		if !stringInSlice(permission, models.GetPermissionsList()) {
			return errors.New(fmt.Sprintf("Invalid Permission: %s", permission))
		}
	}

	return nil
}
