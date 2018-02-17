import {
    HANDLE_PROJECT_SELECTED,
    HANDLE_PROJECTS_LOADED
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
