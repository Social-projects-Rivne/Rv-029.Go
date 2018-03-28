import {
  INCREASE_STEP,
  DECREASE_STEP
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
