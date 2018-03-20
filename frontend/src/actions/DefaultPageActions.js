import {
    HANDLE_DRAWER_TOGGLE,
    HANDLE_ERROR_MESSAGE,
    HANDLE_NOTIFICATION_MESSAGE,
    HANDLE_PAGE_TITLE_CHANGE,
    HANDLE_ADD_USER_TO_PROJECT_TOGGLE,
    HANDLE_PERMISSIONS_LOADED,
    HANDLE_ADD_USER_TO_PROJECT_WITH_PERMISSIONS_TOGGLE,
    HANDLE_IMPORT_USERS_TOGGLE,
    HANDLE_ROLES_LOADED,
    HANDLE_SET_USER,
} from '../constants/defaultPage'

export const setCurrentUser = (user) => {
    return {
        type: HANDLE_SET_USER,
        payload: user
    }
}
export const toggleDrawer = (state) => {
    return {
        type: HANDLE_DRAWER_TOGGLE,
        payload: state
    }
}
export const toggleAddUserToProject = (state) => {
    return {
        type: HANDLE_ADD_USER_TO_PROJECT_TOGGLE,
        payload: state
    }
}
export const changePageTitle = (title) => {
    return {
        type: HANDLE_PAGE_TITLE_CHANGE,
        payload: title
    }
}
export const setErrorMessage = (content) => {
    return {
        type: HANDLE_ERROR_MESSAGE,
        payload: content
    }
}
export const setNotificationMessage = (content) => {
    return {
        type: HANDLE_NOTIFICATION_MESSAGE,
        payload: content
    }
}
export const setPermissionsList = (permissions) => {
    return {
        type: HANDLE_PERMISSIONS_LOADED,
        payload: permissions
    }
}
export const setRolesList = (roles) => {
    return {
        type: HANDLE_ROLES_LOADED,
        payload: roles
    }
}
export const togglePermissionsDialog = (state, user = null) => {
    return {
        type: HANDLE_ADD_USER_TO_PROJECT_WITH_PERMISSIONS_TOGGLE,
        payload: state,
        user: state ? user : null
    }
}
export const toggleImportUsersDialog = (state) => {
    return {
        type: HANDLE_IMPORT_USERS_TOGGLE,
        payload: state,
    }
}