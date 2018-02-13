import { HANDLE_DRAWER_TOGGLE,
    HANDLE_PAGE_TITLE_CHANGE
} from '../constants/top_bar'

const initialState = {
  isDrawerOpen: false,
  pageTitle: null,
  currentProject: null,
  currentBoardProjects: [],
  currentProjectIssues: [],
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_DRAWER_TOGGLE:
      return { ...state, isDrawerOpen: action.payload }
    case HANDLE_PAGE_TITLE_CHANGE:
      return { ...state, pageTitle: action.payload }
    default:
      return state
  }
}
