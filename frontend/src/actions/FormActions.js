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
  CLEAR_STATE
} from '../constants/Form'

export const handleEmail = (email) => {
  return {
    type: HANDLE_EMAIL_INPUT,
    payload: email
  }
}

export const handlePassword = (password) => {
  return {
    type: HANDLE_PASSWORD_INPUT,
    payload: password
  }
}

export const handleName = (name) => {
  return {
    type: HANDLE_NAME_INPUT,
    payload: name
  }
}

export const handleSurname = (surname) => {
  return {
    type: HANDLE_SURNAME_INPUT,
    payload: surname
  }
}

export const isValidEmail = (email) => {
  return {
    type: IS_VALID_EMAIL,
    payload: email
  }
}

export const isValidPassword = (password) => {
  return {
    type: IS_VALID_PASSWORD,
    payload: password
  }
}

export const isValidName = (name) => {
  return {
    type: IS_VALID_NAME,
    payload: name
  }
}

export const isValidSurname = (surname) => {
  return {
    type: IS_VALID_SURNAME,
    payload: surname
  }
}

export const setStatus = (message) => {
  return {
    type: STATUS,
    payload: message
  }
}

export const setErrorMessage = (message) => {
  return {
    type: ERROR_MESSAGE,
    payload: message
  }
}

export const clearState = () => {
  return {
    type: CLEAR_STATE
  }
}

