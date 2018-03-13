import {
    HANDLE_PROJECT_SELECTED,
    HANDLE_PROJECTS_LOADED,
    HANDLE_PROJECT_USERS,
    CURRENT_PROJECT
} from '../constants/projects'

export const setProjects = (projects) => {
    return {
        type: HANDLE_PROJECTS_LOADED,
        payload: projects
    }
}
export const setCurrentProject = (project) => {
    return {
        type: HANDLE_PROJECT_SELECTED,
        payload: project
    }
}
export const setProjectUsers = (users) => {
    return {
        type: HANDLE_PROJECT_USERS,
        payload: users
    }
}
