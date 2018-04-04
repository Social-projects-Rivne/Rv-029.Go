import {
  HANDLE_USER_PROFILE,
  SET_AVAILABLE_USERS,
  SET_ASSIGNED_USERS,
  HANDLE_USERS_TOGGLE,
  HANDLE_SET_IMPORT_FILE,
} from '../constants/users'
  
  const initialState = {
    isUsersOpened: false,
    userInfo: null,
    availableUsers: [],
    assignedUsers: []
  }
  
  export default function (state = initialState, action) {
    console.log(action)
    switch (action.type) {
      case HANDLE_USER_PROFILE:    
        return { ...state, userInfo: action.payload }
      case SET_AVAILABLE_USERS:
        return { ...state, availableUsers: action.payload }
      case SET_ASSIGNED_USERS:
        return { ...state, assignedUsers: action.payload }
      case HANDLE_USERS_TOGGLE:
      console.log('reduser')
        return { ...state, isUsersOpened: action.payload }
      default:
        return state
    }
  }
  