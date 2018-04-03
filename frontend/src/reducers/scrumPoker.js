import {
  INCREASE_STEP,
  DECREASE_STEP,
  SET_STEP,
  SET_EST_RESULT
} from '../constants/scrumPoker'

const initialState = {
  activeStep: 0,
  estimationResult: null
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
    case SET_EST_RESULT:
      return { ...state, estimationResult: payload }
    default:
      return state
  }
}
