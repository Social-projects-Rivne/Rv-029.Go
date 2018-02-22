import {
  HANDLE_SPRINTS_LOADED,
  HANDLE_CURRENT_SPRINT,
  HANDLE_GOAL_UPDATE_SPRINT_INPUT,
  HANDLE_DESC_UPDATE_SPRINT_INPUT,
  HANDLE_STATUS_UPDATE_SPRINT_INPUT,
  HANDLE_SPRINT_OPEN,
  HANDLE_SPRINT_ISSUES_LOADED,
} from '../constants/sprints'

const initialState = {
  issues: [],
  activeSprint: null,
  currentSprint: null,
  currentSprints: [],
  sprintGoal: "",
  sprintDesc: "",
  sprintStatus: ""
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_SPRINT_OPEN:
        return { ...state, activeSprint: action.payload }
    case HANDLE_SPRINTS_LOADED:
      return { ...state, currentSprints: action.payload }
    case HANDLE_SPRINT_ISSUES_LOADED:
      return { ...state, issues: action.payload }
    case HANDLE_CURRENT_SPRINT:
      return {
        ...state,
        currentSprint: action.payload,
        sprintGoal: action.payload.goal,
        sprintDesc: action.payload.description,
        sprintStatus: action.payload.status
      }
    case HANDLE_GOAL_UPDATE_SPRINT_INPUT:
      return { ...state, sprintGoal: action.payload }
    case HANDLE_DESC_UPDATE_SPRINT_INPUT:
      return { ...state, sprintDesc: action.payload }
    case HANDLE_STATUS_UPDATE_SPRINT_INPUT:
      return { ...state, sprintStatus: action.payload }
    default:
      return state
  }
}
