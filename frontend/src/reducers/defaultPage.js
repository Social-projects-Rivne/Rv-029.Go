import {
  HANDLE_DRAWER_TOGGLE,
  HANDLE_ERROR_MESSAGE,
  HANDLE_NOTIFICATION_MESSAGE,
  HANDLE_PAGE_TITLE_CHANGE
} from '../constants/defaultPage'

const initialState = {
  isDrawerOpen: false,
  isUserToProjectOpen: false,
  pageTitle: null,
  errorMessage: null,
  notificationMessage: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_DRAWER_TOGGLE:
      return { ...state, isDrawerOpen: action.payload }
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
