import { combineReducers } from 'redux'
import form from './form'
import projects from './projects'
import defaultPage from './defaultPage'
import { routerReducer } from 'react-router-redux'

export default combineReducers({
  routing: routerReducer,
  form,
  defaultPage,
  projects
})

