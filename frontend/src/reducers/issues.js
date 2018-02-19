import {
  HANDLE_ISSUES_LOADED
} from '../constants/issues'

const initialState = {
  currentIssues: [],
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_ISSUES_LOADED:
      return { ...state, currentIssues: action.payload }
    default:
      return state
  }
}
