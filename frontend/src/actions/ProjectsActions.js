import {
    HANDLE_PROJECTS_LOADED,
    CURRENT_PROJECT
} from '../constants/projects'

export const setProjects = (projects) => {
    return {
        type: HANDLE_PROJECTS_LOADED,
        payload: projects
    }
}

export const setCurrentProject = (projectID) => {
    return {
        type: CURRENT_PROJECT,
        payload: projectID
    }
}
