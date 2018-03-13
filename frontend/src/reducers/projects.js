import {
  HANDLE_PROJECT_SELECTED,
  HANDLE_PROJECTS_LOADED,
  HANDLE_PROJECT_USERS,
  CURRENT_PROJECT
} from '../constants/projects'

const initialState = {
  currentProject: null,
  currentProjects: [],
  currentProjectBoards: [],
  currentProjectUsers: [],
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_PROJECTS_LOADED:
      return { ...state, currentProjects: action.payload }
    case HANDLE_PROJECT_SELECTED:
      return { ...state, currentProject: action.payload }
    case HANDLE_PROJECT_USERS:
        return { ...state, currentProjectUsers: action.payload }
    default:
      return state
  }
}
