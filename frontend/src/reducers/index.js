import { combineReducers } from 'redux'
import form from './form'
import topBar from './top_bar'
import { routerReducer } from 'react-router-redux'

export default combineReducers({
  routing: routerReducer,
  form,
  topBar
})

