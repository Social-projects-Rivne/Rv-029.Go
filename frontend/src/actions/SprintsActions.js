import { HANDLE_SPRINTS_LOADED } from '../constants/sprints'

export const setSprints= (projects) => {
  return {
    type: HANDLE_SPRINTS_LOADED,
    payload: projects
  }
}

