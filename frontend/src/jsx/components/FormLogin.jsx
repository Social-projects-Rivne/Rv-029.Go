import React, { Component } from 'react';
import PropTypes from 'prop-types';
import axios from 'axios';
import FormInput from './FormInput.jsx';
import injectValidation from '../decorators/validate.jsx';
import injectHash from '../decorators/hash.jsx';
import Paper from 'material-ui/Paper';
import Grid from 'material-ui/Grid';
import Button from 'material-ui/Button';
import { withStyles } from 'material-ui/styles';

class FormLogin extends Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
      isValidEmail: true,
      isValidPassword: true
    }
  }

  static propTypes = {
    validateEmail: PropTypes.func.isRequired,
    validateInput: PropTypes.func.isRequired,
    MD5hash: PropTypes.func.isRequired
  };
    
  sendUserData = () => {
    if (!this.checkValidation()) return;

    const MD5 = this.props.MD5hash;

    axios.post('/auth/login', {
      email: this.state.email,
      password: MD5(this.state.password)
    })
    .then((res) => {
      console.log(res)
    })
    .catch((err) => {
      console.log(err)
    })
  };
  
  checkValidation = () => {
    let emailValidation = this.props.validateEmail(this.state.email);
    let passwordValidation = this.props.validateInput(this.state.password);

    this.setState({
      isValidEmail: emailValidation,
      isValidPassword: passwordValidation
    })

    return !!emailValidation && !!passwordValidation;
  };

  handleEmailInput = (e) => {
    this.setState({
      email: e.target.value
    })
  };

  handlePasswordInput = (e) => {
    this.setState({
      password: e.target.value
    })
  };

  render() {
    return (
      <Paper
        className={this.props.classes.root}
        elevation={4}
        component='form'>

        <FormInput
          isValid={this.state.isValidEmail}
          type='text'
          name='Email' 
          onUserInput={this.handleEmailInput} />

        <FormInput
          isValid={this.state.isValidPassword}
          type='password'
          name='Password' 
          onUserInput={this.handlePasswordInput} />

        <Grid item xs={12}>
          <Button raised color="primary"
            onClick={this.sendUserData}>
            Login
          </Button>
        </Grid>

      </Paper>
    )
  }
}

const styles = {
  root: {
    padding: '3em 2em'
  }
};

// TODO looks like shit
let FL = withStyles(styles)(FormLogin);
let F = injectValidation(FL);
export default injectHash(F);
