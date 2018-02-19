import {
  HANDLE_BOARDS_LOADED,
  HANDLE_GOAL_INPUT,
  HANDLE_NAME_INPUT,
  HANDLE_DESC_INPUT,
  HANDLE_ESTIMATION,
  RESET
} from '../constants/boards'

export const setBoards = (boards) => {
  return {
    type: HANDLE_BOARDS_LOADED,
    payload: boards
  }
}

export const setGoal = (goal) => {
  return {
    type: HANDLE_GOAL_INPUT,
    payload: goal
  }
}

export const setName = (name) => {
  return {
    type: HANDLE_NAME_INPUT,
    payload: name
  }
}

export const setDesc = (desc) => {
  return {
    type: HANDLE_DESC_INPUT,
    payload: desc
  }
}

export const setEstimation = (estimation) => {
  return {
    type: HANDLE_ESTIMATION,
    payload: estimation
  }
}

export const resetInput = () => {
  return { type: RESET }
}

