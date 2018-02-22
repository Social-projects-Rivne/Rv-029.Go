import {
  HANDLE_ISSUES_LOADED,
  HANDLE_CURRENT_ISSUE,
  HANDLE_NAME_UPDATE_ISSUE_INPUT,
  HANDLE_DESC_UPDATE_ISSUE_INPUT,
  HANDLE_STATUS_UPDATE_ISSUE_INPUT,
  HANDLE_ESTIMATE_UPDATE_ISSUE_INPUT
} from '../constants/issues'

const initialState = {
  currentIssues: [],
  currentIssue: null,
  issueName: "",
  issueDesc: "",
  issueStatus: "",
  issueEstimate: "",
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_ISSUES_LOADED:
      return { ...state, currentIssues: action.payload }
    case HANDLE_CURRENT_ISSUE:
      return {
        ...state,
        currentIssue: action.payload,
        issueName: action.payload.Name,
        issueDesc: action.payload.Description,
        issueStatus: action.payload.Status,
        issueEstimate: action.payload.Estimate
      }
    case HANDLE_NAME_UPDATE_ISSUE_INPUT:
      return { ...state, issueName: action.payload }
    case HANDLE_DESC_UPDATE_ISSUE_INPUT:
      return { ...state, issueDesc: action.payload }
    case HANDLE_STATUS_UPDATE_ISSUE_INPUT:
      return { ...state, issueStatus: action.payload }
    case HANDLE_ESTIMATE_UPDATE_ISSUE_INPUT:
      return { ...state, issueEstimate: action.payload }
    default:
      return state
  }
}
