import {
  HANDLE_CURRENT_SPRINT,
  HANDLE_SPRINTS_LOADED,
  HANDLE_GOAL_UPDATE_SPRINT_INPUT,
  HANDLE_DESC_UPDATE_SPRINT_INPUT,
  HANDLE_STATUS_UPDATE_SPRINT_INPUT,
} from '../constants/sprints'

export const setSprints = (projects) => {
  return {
    type: HANDLE_SPRINTS_LOADED,
    payload: projects
  }
}

export const setCurrentSprint = (sprint) => {
  return {
    type: HANDLE_CURRENT_SPRINT,
    payload: sprint
  }
}

export const setGoalUpdateSprintInput = (goal) => {
  return {
    type: HANDLE_GOAL_UPDATE_SPRINT_INPUT,
    payload: goal
  }
}

export const setDescUpdateSprintInput = (desc) => {
  return {
    type: HANDLE_DESC_UPDATE_SPRINT_INPUT,
    payload: desc
  }
}

export const setStatusUpdateSprintInput = (status) => {
  return {
    type: HANDLE_STATUS_UPDATE_SPRINT_INPUT,
    payload: status
  }
}
