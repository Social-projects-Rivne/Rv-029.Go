import {
    HANDLE_USER_PROFILE,
  } from '../constants/user'
  
  const initialState = {
    userInfo: null,
  }
  
  export default function (state = initialState, action) {
    switch (action.type) {
      case HANDLE_USER_PROFILE:
        return { ...state, userInfo: action.payload }
      default:
        return state
    }
  }
  