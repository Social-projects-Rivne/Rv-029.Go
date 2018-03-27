import {
  HANDLE_USER_PROFILE,
  SET_AVAILABLE_USERS,
  SET_ASSIGNED_USERS,
} from '../constants/users'
  
  const initialState = {
    userInfo: null,
    availableUsers: [],
    assignedUsers: []
  }
  
  export default function (state = initialState, action) {
    switch (action.type) {
      case HANDLE_USER_PROFILE:
        return { ...state, userInfo: action.payload }
      case SET_AVAILABLE_USERS:
        return { ...state, availableUsers: action.payload }
      case SET_ASSIGNED_USERS:
        return { ...state, assignedUsers: action.payload }
      default:
        return state
    }
  }
  