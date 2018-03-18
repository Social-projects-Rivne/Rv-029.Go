import {
  HANDLE_ISSUES_LOADED,
  HANDLE_CURRENT_ISSUE,
  HANDLE_NAME_UPDATE_ISSUE_INPUT,
  HANDLE_DESC_UPDATE_ISSUE_INPUT,
  HANDLE_STATUS_UPDATE_ISSUE_INPUT,
  HANDLE_ESTIMATE_UPDATE_ISSUE_INPUT,
  SET_ISSUES_HIERARCHY
} from '../constants/issues'

const initialState = {
  currentIssues: [],
  currentIssue: null,
  issueName: "",
  issueDesc: "",
  issueStatus: "",
  issueEstimate: "",
  hierarchy: null
}

export default (state = initialState, action) => {
  const { type, payload } = action

  switch (type) {
    case HANDLE_ISSUES_LOADED:
      return { ...state, currentIssues: payload }
    case HANDLE_CURRENT_ISSUE:
      return {
        ...state,
        currentIssue: payload,
        issueName: payload.Name,
        issueDesc: payload.Description,
        issueStatus: payload.Status,
        issueEstimate: payload.Estimate
      }
    case HANDLE_NAME_UPDATE_ISSUE_INPUT:
      return { ...state, issueName: payload }
    case HANDLE_DESC_UPDATE_ISSUE_INPUT:
      return { ...state, issueDesc: payload }
    case HANDLE_STATUS_UPDATE_ISSUE_INPUT:
      return { ...state, issueStatus: payload }
    case HANDLE_ESTIMATE_UPDATE_ISSUE_INPUT:
      return { ...state, issueEstimate: payload }
    case SET_ISSUES_HIERARCHY:
      return { ...state, hierarchy: payload}
    default:
      return state
  }
}
