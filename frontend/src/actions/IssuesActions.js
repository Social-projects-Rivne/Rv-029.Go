import {
  HANDLE_ISSUES_LOADED
} from '../constants/issues'

export const setIssues = (issues) => {
  return {
    type: HANDLE_ISSUES_LOADED,
    payload: issues
  }
}
