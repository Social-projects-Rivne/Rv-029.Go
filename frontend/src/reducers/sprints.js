import {
  HANDLE_SPRINTS_LOADED,
  HANDLE_CURRENT_SPRINT
} from '../constants/sprints'

const initialState = {
  currentSprint: null,
  currentSprints: [],
  sprintGoal: "", // todo: more fields
}

// todo: handles for sprintGoal etc

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_SPRINTS_LOADED:
      return { ...state, currentSprints: action.payload }
    case HANDLE_CURRENT_SPRINT:
      return {
        ...state,
        currentSprint: action.payload,
        sprintGoal: action.payload.goal
      }
    default:
      return state
  }
}
