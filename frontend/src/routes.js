import React from 'react'
import { Route, IndexRoute, IndexRedirect } from 'react-router'
import axios from 'axios'
import { API_URL } from './constants/global'

import { clearInputState } from './actions/FormActions'
import { setNotificationMessage } from './actions/FormActions'
import { setErrorMessage } from './actions/FormActions'
import { store } from './store/configureStore'

import App from './App'
import FormContainer from './containers/FormContainer'
import FormLogin from './components/FormLogin'
import FormRegister from './components/FormRegister'
import FormRestorePassword from './components/FormRestorePassword'
import FormNewPassword from './components/FormNewPassword'
import ProjectCard from "./components/ProjectCard";

import HomePage from './containers/HomePage'
import DefaultPage from './containers/DefaultPage'
import ProjectsPage from "./containers/ProjectsPage";

import auth from './services/auth'
import ViewProjectPage from './components/ViewProjectPage'

// TODO move these out of here
// e.g. routeEvents.js
const reset = () => {
  store.dispatch(clearInputState())
}

const queryCheck = (nextState, replace, callback) => {
  if (nextState.location.query.token && nextState.location.query.uuid) {
    axios.post(API_URL + 'auth/confirm', {
      token: nextState.location.query.token,
      uuid: nextState.location.query.uuid
    })
    .then((res) => {
      store.dispatch(setNotificationMessage(res.data.Message))
      replace("/authorization/login")
      callback()
    })
    .catch((err) => {
      replace("/authorization/login")

      if (err.response.data.Message) {
        store.dispatch(setErrorMessage(err.response.data.Message))
      } else {
        store.dispatch(setErrorMessage('Server error occured'))
      }
      callback()
    })
  } else callback()
}

const authorizedMiddleware = (nextState, replace, callback) => {
    if (!auth.injectAuthHeader()) {
       replace('/authorization/login')
    }

    callback()
}

export const routes = (
  <div>
    <Route path="/" component={App}>
      <IndexRedirect to ="authorization/login"/>
      <Route path="authorization" component={FormContainer} onChange={reset}>
        <IndexRedirect to="login"/>
        <Route name="login" path="login" component={FormLogin} onEnter={queryCheck}/>
        <Route name="registration" path="register" component={FormRegister}/>
        <Route name="restore-password" path="restore-password" component={FormRestorePassword}/>
        <Route name="reset-password" path="new-password/:token" component={FormNewPassword}/>
      </Route>

      <Route component={DefaultPage} onEnter={authorizedMiddleware}>
        <Route name="projects_list" path="projects" component={ProjectsPage}/>
        <Route name="view_project" path="project/:id" component={ViewProjectPage}/>
      </Route>

      <Route path="home-page" component={HomePage}/>
    </Route>
  </div>
)
