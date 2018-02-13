import React from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { Link, browserHistory } from 'react-router'
import PropTypes from 'prop-types'
import axios from 'axios'
import { API_URL } from '../constants/global'
import * as formActions from '../actions/FormActions'
import * as topBarActions from '../actions/TopBarActions';
import FormInput from './FormInput'
import SnackBar from './SnackBar'
import ModalNotification from './ModalNotification'
import injectValidation from '../decorators/validate'
import injectHash from '../decorators/hash'
import Paper from 'material-ui/Paper'
import Grid from 'material-ui/Grid'
import Button from 'material-ui/Button'
import Typography from 'material-ui/Typography'
import { withStyles } from 'material-ui/styles'

const FormLogin = ({ classes, form, topBar, action, topBarAction, ownProps, ...decorator }) => {

  const sendUserData = (e) => {
    e.preventDefault()

    if (!checkValidation()) return

    axios.post(API_URL + 'auth/login', {
      email: form.email,
      password: decorator.MD5Encode(form.password)
    })
    .then((response) => {
      //TODO: add global function for auth header
      if (response.data.Status) {
          sessionStorage.setItem('token', response.data.Token)
          axios.defaults.headers.common['Authorization'] = 'Bearer ' + response.data.Token;
          topBarAction.changePageTitle('Projects')
          browserHistory.push('/projects')
      }
    })
    .catch((err) => {
      action.setStatus(err.response.data.status)

      if (err.response.data.Message) {
        action.setErrorMessage(err.response.data.Message)
      } else {
        action.setErrorMessage("Server error occured")
      }
    })
  }
  
  const checkValidation = () => {
    let emailValidation = decorator.validateEmail(form.email)
    let passwordValidation = decorator.validatePassword(form.password)
    
    action.isValidEmail(!!emailValidation)
    action.isValidPassword(!!passwordValidation)

    return !!emailValidation && !!passwordValidation
  }
  
  const handleEmailInput = (e) => {
    action.handleEmail(e.target.value)
  }

  const handlePasswordInput = (e) => {
    action.handlePassword(e.target.value)
  }

  return (
    <Paper
      className={classes.paper}
      elevation={8}
      component='form'>

      <Typography
        type='headline'
        component='h3'>
        Log In
      </Typography>
      
      <FormInput
        type='text'
        name='Email'
        autoFocus={true}
        isValid={form.isValidEmail}
        onUserInput={handleEmailInput} />
      
      <FormInput
        type='password'
        name='Password'
        isValid={form.isValidPassword}
        onUserInput={handlePasswordInput} />
      
      <Grid
        container
        alignItems={'center'}
        justify={'space-around'}
        className={classes.buttons}>
          <Button
            type='submit'
            color='primary'
            onClick={sendUserData}>
            Submit
          </Button>
          <Link to={'/authorization/register'}
            className={classes.link}>
            <Button
              color={'secondary'}>
              Registration
            </Button>
          </Link>
      </Grid>

      <Grid
        container
        justify={'center'}>

        <Link to={'/authorization/restore-password'}>
          <Typography
            style={{marginTop: 15, cursor: 'pointer'}}
            type='caption'
            color='primary'
            component='h3'>
            I have forgotten my password
          </Typography>
        </Link>

      </Grid>

      <SnackBar
        errorMessage={form.errorMessage}
        setErrorMessage={action.setErrorMessage}/>

      <ModalNotification
        title='Notification'
        content={form.notificationMessage}
        setNotificationMessage={action.setNotificationMessage}/>

    </Paper>
  )
}

FormLogin.propTypes = {
  validateEmail: PropTypes.func.isRequired,
  validatePassword: PropTypes.func.isRequired,
  MD5Encode: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
  form: PropTypes.object.isRequired,
  action: PropTypes.object.isRequired
}

const styles = {
  paper: {
    padding: '4em 3em',
  },
  buttons: {
    marginTop: '1.5em',
  },
  link: {
    textDecoration: 'none'
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    form: state.form,
    topBar: state.topBar,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    action: bindActionCreators(formActions, dispatch),
    topBarAction: bindActionCreators(topBarActions, dispatch)
  }
}

export default injectHash(
  injectValidation(
    withStyles(styles)(
      connect(mapStateToProps, mapDispatchToProps)(FormLogin)
    )
  )
)

