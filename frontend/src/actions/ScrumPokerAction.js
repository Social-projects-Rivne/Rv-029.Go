import {
  INCREASE_STEP,
  DECREASE_STEP,
  SET_STEP,
  SET_EST_RESULT
} from '../constants/scrumPoker'

export const increaseStep = () => {
  return {
    type: INCREASE_STEP
  }
}

export const decreaseStep = () => {
  return {
    type: DECREASE_STEP
  }
}

export const setStep = (step) => {
  return {
    type: SET_STEP,
    payload: step
  }
}

export const setEstResult = (result) => {
  return {
    type: SET_EST_RESULT,
    payload: result
  }
}
