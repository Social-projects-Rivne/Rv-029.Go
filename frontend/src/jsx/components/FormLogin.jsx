import React, { Component } from 'react';
import PropTypes from 'prop-types';
import axios from 'axios';
import FormInput from './FormInput.jsx';
import injectValidation from '../decorators/validate.jsx';
import injectHash from '../decorators/hash.jsx';
import Paper from 'material-ui/Paper';
import Grid from 'material-ui/Grid';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';
import { SnackbarContent } from 'material-ui/Snackbar';
import { withStyles } from 'material-ui/styles';

class FormLogin extends Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
      isValidEmail: true,
      isValidPassword: true,
      errorMessage: null
    }
  }

  static propTypes = {
    validateEmail: PropTypes.func.isRequired,
    validateInput: PropTypes.func.isRequired,
    MD5hash: PropTypes.func.isRequired
  };
    
  sendUserData = (e) => {
    e.preventDefault();
    
    // TODO check validation onBlur
    if (!this.checkValidation()) return;

    const MD5 = this.props.MD5hash;

    axios.post('/auth/login', {
      email: this.state.email,
      password: MD5(this.state.password)
    })
    .then((res) => {
      console.log(res);
      // TODO token
      //localStorage.setItem('token', res.token);
    })
    .catch((err) => {
      switch (err.response.status) {
      case 401:
        this.setState({
          errorMessage: err.response.data.Message
        });
        break;
      default:
        // TODO handle server errors
        console.log(err);
      }
    })
  };
  
  checkValidation = () => {
    let emailValidation = this.props.validateEmail(this.state.email);
    let passwordValidation = this.props.validateInput(this.state.password);

    this.setState({
      isValidEmail: emailValidation,
      isValidPassword: passwordValidation
    });

    // false if at least one isn't correct
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
    const errorMessage = this.state.errorMessage;
    
    return (
      <Paper
        className={this.props.classes.root}
        elevation={4}
        component='form'>
        
        <Typography
          type='headline'
          component='h3'>
          Log In
        </Typography>
        
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
          <Button
            type='submit'
            raised color='primary'
            onClick={this.sendUserData}
            className={this.props.classes.button}>
            Submit
          </Button>
        </Grid>
  
        {/* TODO make banner more elegance */}
        {errorMessage ? (
          <Typography>
            {errorMessage}
          </Typography>
        ):(null)}
        
      </Paper>
    )
  }
}

const styles = {
  root: {
    padding: '4em 3em'
  },
  button: {
    marginTop: '1.5em',
  },
};

let FL = withStyles(styles)(FormLogin);
let F = injectValidation(FL);
export default injectHash(F);
