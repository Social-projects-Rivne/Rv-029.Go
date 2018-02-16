import {
    HANDLE_PROJECTS_LOADED
} from '../constants/projects'

export const setProjects = (projects) => {
    return {
        type: HANDLE_PROJECTS_LOADED,
        payload: projects
    }
}
