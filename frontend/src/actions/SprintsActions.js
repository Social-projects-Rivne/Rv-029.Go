import {
  HANDLE_CURRENT_SPRINT,
  HANDLE_SPRINTS_LOADED
} from '../constants/sprints'

export const setSprints= (projects) => {
  return {
    type: HANDLE_SPRINTS_LOADED,
    payload: projects
  }
}

export const setCurrentSprint= (sprint) => {
  return {
    type: HANDLE_CURRENT_SPRINT,
    payload: sprint
  }
}
