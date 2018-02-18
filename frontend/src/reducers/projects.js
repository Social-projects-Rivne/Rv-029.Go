import {
  HANDLE_PROJECTS_LOADED,
  CURRENT_PROJECT
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
    case CURRENT_PROJECT:
      return { ...state, currentProject: action.payload }
    default:
      return state
  }
}
