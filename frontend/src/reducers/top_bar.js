import { HANDLE_DRAWER_TOGGLE } from '../constants/top_bar'

const initialState = {
  isDrawerOpen: false,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case HANDLE_DRAWER_TOGGLE:
      return { ...state, isDrawerOpen: action.payload }
    default:
      return state
  }
}
