import { combineReducers } from 'redux'
import form from './form'
import projects from './projects'
import boards from './boards'
import sprints from './sprints'
import issues from './issues'
import user from './user'
import defaultPage from './defaultPage'
import { routerReducer } from 'react-router-redux'

export default combineReducers({
  routing: routerReducer,
  form,
  defaultPage,
  projects,
  boards,
  sprints,
  issues,
  user
})

