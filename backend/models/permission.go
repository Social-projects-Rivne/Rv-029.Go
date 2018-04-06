package models

const PERMISSION_CREATE_PROJECTS  = `project.create`
const PERMISSION_UPDATE_PROJECTS  = `project.update`
const PERMISSION_DELETE_PROJECTS  = `project.delete`
const PERMISSION_CREATE_BOARDS  = `board.create`
const PERMISSION_UPDATE_BOARDS  = `board.update`
const PERMISSION_DELETE_BOARDS  = `board.delete`
const PERMISSION_CREATE_ISSUES  = `issue.create`
const PERMISSION_UPDATE_ISSUES  = `issue.update`
const PERMISSION_DELETE_ISSUES  = `issue.delete`
const PERMISSION_CREATE_SPRINTS  = `sprint.create`
const PERMISSION_UPDATE_SPRINTS  = `sprint.update`
const PERMISSION_DELETE_SPRINTS  = `sprint.delete`
const PERMISSION_ADD_ISSUES_TO_SPRINTS  = `sprint.add.issue`
const PERMISSION_MANAGE_USER_PERMISSIONS  = `user.permissions.manage`


func GetPermissionsList() []string {
	return []string{
		PERMISSION_CREATE_PROJECTS,
		PERMISSION_UPDATE_PROJECTS,
		PERMISSION_DELETE_PROJECTS,
		PERMISSION_CREATE_BOARDS,
		PERMISSION_UPDATE_BOARDS,
		PERMISSION_DELETE_BOARDS,
		PERMISSION_CREATE_ISSUES,
		PERMISSION_UPDATE_ISSUES,
		PERMISSION_DELETE_ISSUES,
		PERMISSION_CREATE_SPRINTS,
		PERMISSION_UPDATE_SPRINTS,
		PERMISSION_DELETE_SPRINTS,
		PERMISSION_ADD_ISSUES_TO_SPRINTS,
		PERMISSION_MANAGE_USER_PERMISSIONS,
	}
}

func GetRolesList() []string {
	return []string{
		ROLE_ADMIN,
		ROLE_USER,
		ROLE_STAFF,
	}
}