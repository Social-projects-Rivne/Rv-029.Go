import {
  HANDLE_PROJECT_SELECTED,
  HANDLE_PROJECTS_LOADED
} from '../constants/projects'

const initialState = {
  currentProject: null,
  currentProjects: [],
  currentProjectBoards: [],
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_PROJECTS_LOADED:
      return { ...state, currentProjects: action.payload }
    case HANDLE_PROJECT_SELECTED:
      return { ...state, currentProject: action.payload }
    default:
      return state
  }
}
