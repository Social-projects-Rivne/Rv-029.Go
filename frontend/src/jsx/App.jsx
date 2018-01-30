import React, {Component} from 'react';
import { Route, Router, BrowserRouter } from 'react-router-dom';
import FormLogin from './components/FormLogin.jsx';
import FormRegister from './components/FormRegister.jsx';
import FormContainer from './containers/FormContainer.jsx';
//import PropTypes from 'prop-types';

var API_URL = "http://localhost:3000";

export default class App extends Component {
  render() {
    return (
      <FormContainer>
        <FormLogin />
      </FormContainer>
    )
  }
}
