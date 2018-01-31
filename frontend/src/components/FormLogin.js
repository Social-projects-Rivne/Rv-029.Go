import React, { Component } from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import FormInput from './FormInput'
import SnackBar from './SnackBar'
import injectValidation from '../decorators/validate'
import injectHash from '../decorators/hash'
import Paper from 'material-ui/Paper'
import Grid from 'material-ui/Grid'
import Button from 'material-ui/Button'
import Typography from 'material-ui/Typography'
import { withStyles } from 'material-ui/styles'

const FormLogin = ({classes, form, action, ...decorator}) => {
  
  const sendUserData = (e) => {
    e.preventDefault()

    if (!checkValidation()) return

    axios.post('/auth/login', {
      email: form.email,
      password: decorator.MD5Encode(form.password)
    })
    .then((res) => {
      console.log(res)
      sessionStorage.setItem('token', res.token)
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

  const toggleFormToRestorePassword = () => {
    action.toggleFormType('restorePassword')
  }

  const toggleformToRegister = () => {
    action.toggleFormType('register')
  }

  return (
    <Paper
      className={classes.root}
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
          <Button
            color={'secondary'}
            onClick={toggleformToRegister}>
            Registration
          </Button>
      </Grid>

      <Grid
        container
        justify={'center'}>

        <Typography
          style={{marginTop: 15, cursor: 'pointer'}}
          type='caption'
          color='primary'
          component='h3'
          onClick={toggleFormToRestorePassword}>
          I have forgotten my password
        </Typography>
      </Grid>

      <SnackBar
        errorMessage={form.errorMessage}
        setErrorMessage={action.setErrorMessage}/>
      
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
  root: {
    padding: '4em 3em'
  },
  buttons: {
    marginTop: '1.5em',
  },
}

export default injectHash(
  injectValidation(
    withStyles(styles)(FormLogin)
  )
)

