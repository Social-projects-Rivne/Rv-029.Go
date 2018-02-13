import { HANDLE_DRAWER_TOGGLE,
    HANDLE_PROJECTS_LOADED,
    HANDLE_PAGE_TITLE_CHANGE
} from '../constants/top_bar'

const initialState = {
  isDrawerOpen: false,
  pageTitle: null,
  currentProject: null,
  currentProjects: [],
  currentProjectBoards: [],
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_DRAWER_TOGGLE:
      return { ...state, isDrawerOpen: action.payload }
    case HANDLE_PAGE_TITLE_CHANGE:
      return { ...state, pageTitle: action.payload }
    case HANDLE_PROJECTS_LOADED:
      return { ...state, currentProjects: action.payload }
    default:
      return state
  }
}
