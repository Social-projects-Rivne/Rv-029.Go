package controllers

import (
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/gorilla/mux"
)

func AddUserPermission(w http.ResponseWriter, r *http.Request) {
	var addPermissionRequestData validator.PermissionRequestData
	err := decodeAndValidate(r, &addPermissionRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	role := models.Role{Name: vars["role_name"]}
	err = models.RoleDB.FindByName(&role)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	role.AddPermission(addPermissionRequestData.Permission)

	err = models.RoleDB.Update(&role)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "User permission has been updated", Data: role}
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

	vars := mux.Vars(r)
	role := models.Role{Name: vars["role_name"]}
	err = models.RoleDB.FindByName(&role)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}
	role.RemovePermission(removePermissionRequestData.Permission)

	err = models.RoleDB.Update(&role)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Role updated", Data: role}
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
	
	vars := mux.Vars(r)
	role := models.Role{Name: vars["role_name"]}
	err = models.RoleDB.FindByName(&role)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}
	role.SetPermissions(setPermissionsRequestData.Permissions)

	err = models.RoleDB.Update(&role)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Role updated", Data: role}
	response.Success(w)
}

