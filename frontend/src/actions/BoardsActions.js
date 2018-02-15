import { HANDLE_BOARDS_LOADED } from '../constants/boards'

export const setBoards = (boards) => {
  return {
    type: HANDLE_BOARDS_LOADED,
    payload: boards
  }
}
