package controllers

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
)

func AddUserPermission(w http.ResponseWriter, r *http.Request) {
	var addPermissionRequestData validator.PermissionRequestData
	err := decodeAndValidate(r, &addPermissionRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)
	user.AddPermission(addPermissionRequestData.Permission)

	err = models.UserDB.Update(&user)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "User permission has been updated", Data: user}
	response.Success(w)
}

func RemoveUserPermissions(w http.ResponseWriter, r *http.Request) {
	var removePermissionRequestData validator.PermissionRequestData
	err := decodeAndValidate(r, &removePermissionRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	user := r.Context().Value("user").(models.User)
	user.RemovePermission(removePermissionRequestData.Permission)

	err = models.UserDB.Update(&user)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has updated", Data: user}
	response.Success(w)
}

func SetUserPermissions(w http.ResponseWriter, r *http.Request) {
	var setPermissionsRequestData validator.SetPermissionsRequestData
	err := decodeAndValidate(r, &setPermissionsRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}


	user := r.Context().Value("user").(models.User)
	user.RemovePermission(setPermissionsRequestData.Permission)

	err = models.UserDB.Update(&user)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has updated", Data: user}
	response.Success(w)
}

