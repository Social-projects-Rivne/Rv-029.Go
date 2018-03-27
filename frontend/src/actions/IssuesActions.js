import {
  HANDLE_ISSUES_LOADED,
  HANDLE_CURRENT_ISSUE,
  HANDLE_NAME_UPDATE_ISSUE_INPUT,
  HANDLE_DESC_UPDATE_ISSUE_INPUT,
  HANDLE_STATUS_UPDATE_ISSUE_INPUT,
  HANDLE_ESTIMATE_UPDATE_ISSUE_INPUT,
  SET_ISSUES_HIERARCHY
} from '../constants/issues'

export const setIssues = (issues) => {
  return {
    type: HANDLE_ISSUES_LOADED,
    payload: issues
  }
}

export const setCurrentIssue = (issue) => {
  return {
    type: HANDLE_CURRENT_ISSUE,
    payload: issue
  }
}

export const setNameUpdateIssueInput = (name) => {
  return {
    type: HANDLE_NAME_UPDATE_ISSUE_INPUT,
    payload: name
  }
}

export const setStatusUpdateIssueInput = (status) => {
  return {
    type: HANDLE_STATUS_UPDATE_ISSUE_INPUT,
    payload: status
  }
}

export const setDescUpdateIssueInput = (desc) => {
  return {
    type: HANDLE_DESC_UPDATE_ISSUE_INPUT,
    payload: desc
  }
}

export const setEstimateUpdateIssueInput = (estimate) => {
  return {
    type: HANDLE_ESTIMATE_UPDATE_ISSUE_INPUT,
    payload: estimate
  }
}

export const setIssuesHierarchy = (obj) => {
  return {
    type: SET_ISSUES_HIERARCHY,
    payload: obj
  }
}