import {
  HANDLE_BOARDS_LOADED,
  HANDLE_GOAL_INPUT,
  HANDLE_NAME_INPUT,
  HANDLE_DESC_INPUT,
  HANDLE_ESTIMATION,
  RESET
} from '../constants/boards'

const initialState = {
  currentBoards: [],
  goalInput: "",
  nameInput: "",
  descInput: "",
  estimation: ""
}

export default (state = initialState, action) => {
  switch (action.type) {
    case HANDLE_BOARDS_LOADED:
      return { ...state, currentBoards: action.payload }
    case HANDLE_GOAL_INPUT:
      return { ...state, goalInput: action.payload }
    case HANDLE_NAME_INPUT:
      return { ...state, nameInput: action.payload }
    case HANDLE_DESC_INPUT:
      return { ...state, descInput: action.payload }
    case HANDLE_ESTIMATION:
      return { ...state, estimation: action.payload }
    case RESET:
      return {
        ...state,
        goalInput: "",
        nameInput: "",
        descInput: "",
        estimation: ""
      }
    default:
      return state
  }
}
