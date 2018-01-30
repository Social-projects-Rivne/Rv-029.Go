import { HANDLE_EMAIL_INPUT } from "../constants/Form"
import { HANDLE_PASSWORD_INPUT } from "../constants/Form"
import { IS_VALID_EMAIL } from "../constants/Form"
import { IS_VALID_PASSWORD } from "../constants/Form"
import { IS_VALID_NAME } from "../constants/Form"
import { IS_VALID_SURNAME } from "../constants/Form"
import { STATUS } from "../constants/Form"
import { ERROR_MESSAGE } from "../constants/Form"
import { HANDLE_NAME_INPUT } from "../constants/Form"
import { HANDLE_SURNAME_INPUT } from "../constants/Form"
import { FORM_TYPE } from "../constants/Form"

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
  type: 'login'
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
    case FORM_TYPE:
      return { ...initialState, type: action.payload }
    default:
      return state
  }
}
