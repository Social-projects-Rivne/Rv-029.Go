import {
  HANDLE_PROJECT_SELECTED,
  HANDLE_PROJECTS_LOADED,
  HANDLE_PROJECT_USERS,
  HANDLE_USERS_LIST
} from '../constants/projects'

const initialState = {
  users: [],
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
    case HANDLE_USERS_LIST:
        return { ...state, users: action.payload }
    default:
      return state
  }
}
