import {
  HANDLE_USER_PROFILE,
  SET_AVAILABLE_USERS,
  SET_ASSIGNED_USERS,
  HANDLE_USERS_TOGGLE,
  HANDLE_SET_IMPORT_FILE,
} from '../constants/users'
  

export const setCurrentUser = (user) => {
  return {
    type: HANDLE_USER_PROFILE,
    payload: user
  }
}

export const setAvailableUsers = (users) => {
  return {
    type: SET_AVAILABLE_USERS,
    payload: users
  }
}

export const setAssignedUsers = (users) => {
  return {
    type: SET_ASSIGNED_USERS,
    payload: users
  }
}

export const toggleUsersDialog = (state) => {
  console.log('actions')
  return {
      type: HANDLE_USERS_TOGGLE,
      payload: state,
  }
}

export const setUsersPhotoToImport = (file) => {
  return {
      type: HANDLE_SET_IMPORT_FILE,
      payload: file,
  }
}
export const resetUsersPhotoToImport = () => {
  return {
      type: HANDLE_SET_IMPORT_FILE,
      payload: null,
  }
}

