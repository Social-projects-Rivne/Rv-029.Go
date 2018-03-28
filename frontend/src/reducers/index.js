import { combineReducers } from 'redux'
import form from './form'
import projects from './projects'
import boards from './boards'
import sprints from './sprints'
import issues from './issues'
import users from './users'
import defaultPage from './defaultPage'
import scrumPoker from './scrumPoker'
import { routerReducer } from 'react-router-redux'

export default combineReducers({
  routing: routerReducer,
  form,
  defaultPage,
  projects,
  boards,
  sprints,
  issues,
  users,
  scrumPoker
})

