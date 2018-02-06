import React from 'react'
import { Route, IndexRoute, IndexRedirect } from 'react-router'

import { clearInputState } from './actions/FormActions'
import { store } from './store/configureStore'

import App from './App'
import FormContainer from './containers/FormContainer'
import FormLogin from './components/FormLogin'
import FormRegister from './components/FormRegister'
import FormRestorePassword from './components/FormRestorePassword'
import FormNewPassword from './components/FormNewPassword'

import HomePage from './containers/HomePage'

const reset = () => {
  store.dispatch(clearInputState())
}

export const routes = (
  <div>
    <Route path="/" component={App}>
      <IndexRedirect to ="authorization/login"/>
      <Route path="authorization" component={FormContainer} onChange={reset}>
        <IndexRedirect to="login"/>
        <Route path="login" component={FormLogin}/>
        <Route path="register" component={FormRegister}/>
        <Route path="restore-password" component={FormRestorePassword}/>
        <Route path="new-password/:token" component={FormNewPassword}/>
      </Route>
      <Route path="home-page" component={HomePage}/>
    </Route>
  </div>
)
