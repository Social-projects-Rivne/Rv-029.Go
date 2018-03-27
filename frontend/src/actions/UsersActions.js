import {
  HANDLE_USER_PROFILE,
  SET_AVAILABLE_USERS,
  SET_ASSIGNED_USERS,
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