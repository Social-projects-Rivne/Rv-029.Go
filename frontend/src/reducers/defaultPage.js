import {
  HANDLE_DRAWER_TOGGLE,
  HANDLE_ERROR_MESSAGE,
  HANDLE_NOTIFICATION_MESSAGE,
  HANDLE_PAGE_TITLE_CHANGE,
  HANDLE_ADD_USER_TO_PROJECT_TOGGLE,
  HANDLE_SET_USER
} from '../constants/defaultPage'

const initialState = {
  isDrawerOpen: false,
  isUserToProjectOpen: false,
  pageTitle: null,
  user: null,
  errorMessage: null,
  notificationMessage: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_DRAWER_TOGGLE:
      return { ...state, isDrawerOpen: action.payload }
    case HANDLE_SET_USER:
      return { ...state, user: action.payload }
    case HANDLE_ADD_USER_TO_PROJECT_TOGGLE:
      return { ...state, isUserToProjectOpen: action.payload }
    case HANDLE_PAGE_TITLE_CHANGE:
      return { ...state, pageTitle: action.payload }
    case HANDLE_ERROR_MESSAGE:
      return { ...state, errorMessage: action.payload }
    case HANDLE_NOTIFICATION_MESSAGE:
      return { ...state, notificationMessage: action.payload }
    default:
      return state
  }
}
