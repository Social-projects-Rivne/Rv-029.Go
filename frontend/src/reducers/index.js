import { combineReducers } from 'redux'
import form from './form'
import { routerReducer } from 'react-router-redux'

export default combineReducers({
  routing: routerReducer,
  form
})

