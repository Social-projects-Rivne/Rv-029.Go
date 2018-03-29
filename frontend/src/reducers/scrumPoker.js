import {
  INCREASE_STEP,
  DECREASE_STEP,
  SET_STEP
} from '../constants/scrumPoker'

const initialState = {
  activeStep: 0
}

export default (state = initialState, action) => {
  const { type, payload } = action

  switch (type) {
    case INCREASE_STEP:
      return { ...state, activeStep: state.activeStep + 1 }
    case DECREASE_STEP:
      return { ...state, activeStep: state.activeStep - 1 }
    case SET_STEP:
      return { ...state, activeStep: payload }
    default:
      return state
  }
}
