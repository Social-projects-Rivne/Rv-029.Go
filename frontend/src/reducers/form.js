import {
  HANDLE_EMAIL_INPUT,
  HANDLE_PASSWORD_INPUT,
  IS_VALID_EMAIL,
  IS_VALID_NAME,
  IS_VALID_SURNAME,
  IS_VALID_PASSWORD,
  STATUS,
  ERROR_MESSAGE,
  HANDLE_NAME_INPUT,
  HANDLE_SURNAME_INPUT,
  CLEAR_INPUT_STATE, NOTIFICATION_MESSAGE
} from '../constants/form'

const initialState = {
  email: '',
  password: '',
  name: '',
  surname: '',
  isValidEmail: true,
  isValidPassword: true,
  isValidName: true,
  isValidSurname: true,
  status: null,
  errorMessage: null,
  notificationMessage: null
}

export default function login(state = initialState, action) {
  switch (action.type) {
    case HANDLE_EMAIL_INPUT:
      return { ...state, email: action.payload }
    case HANDLE_PASSWORD_INPUT:
      return { ...state, password: action.payload }
    case IS_VALID_EMAIL:
      return { ...state, isValidEmail: action.payload }
    case IS_VALID_PASSWORD:
      return { ...state, isValidPassword: action.payload }
    case IS_VALID_NAME:
      return { ...state, isValidName: action.payload }
    case IS_VALID_SURNAME:
      return { ...state, isValidSurname: action.payload }
    case STATUS:
      return { ...state, status: action.payload }
    case HANDLE_NAME_INPUT:
      return { ...state, name: action.payload }
    case HANDLE_SURNAME_INPUT:
      return { ...state, surname: action.payload }
    case ERROR_MESSAGE:
      return { ...state, errorMessage: action.payload }
    case NOTIFICATION_MESSAGE:
      return { ...state, notificationMessage: action.payload }
    case CLEAR_INPUT_STATE:
      return { ...state,
        isValidEmail: true,
        isValidPassword: true,
        isValidName: true,
        isValidSurname: true,
      }
    default:
      return state
  }
}
