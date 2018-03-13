import {
    HANDLE_USER_PROFILE,
  } from '../constants/user'
  

export const setCurrentUser = (user) => {
    return {
      type: HANDLE_USER_PROFILE,
      payload: user
    }
  }