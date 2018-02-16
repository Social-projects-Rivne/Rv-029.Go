import { HANDLE_BOARDS_LOADED } from '../constants/boards'

const initialState = {
  currentBoard: null,
  currentBoards: [],
}

export default (state = initialState, action) => {
  switch (action.type) {
    case HANDLE_BOARDS_LOADED:
      return { ...state, currentBoards: action.payload }
    default:
      return state
  }
}
