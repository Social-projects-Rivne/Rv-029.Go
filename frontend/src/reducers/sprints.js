import {
  HANDLE_SPRINTS_LOADED,
} from '../constants/sprints'

const initialState = {
  currentSprints: [],
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_SPRINTS_LOADED:
      return { ...state, currentSprints: action.payload }
    default:
      return state
  }
}
