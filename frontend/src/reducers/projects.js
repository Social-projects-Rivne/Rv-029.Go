import {
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
    default:
      return state
  }
}
