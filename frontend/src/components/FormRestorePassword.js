import React, { Component } from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import FormInput from './FormInput'
import SnackBar from './SnackBar'
import injectValidation from '../decorators/validate'
import injectHash from '../decorators/hash'
import Paper from 'material-ui/Paper'
import Button from 'material-ui/Button'
import Grid from 'material-ui/Grid'
import Typography from 'material-ui/Typography'
import { withStyles } from 'material-ui/styles'

const FormRestorePassword = ({ classes, form, action, ...decorator }) => {

  const sendUserData = (e) => {
    e.preventDefault()

    if (!checkValidation()) return

    axios.post('/auth/restore', {
      email: form.email,
    })
    .then((res) => {
      //TODO restore password
      console.log(res)
    })
    .catch((err) => {

      action.setStatus(err.response.data.status)

      if (err.response.message) {
        action.setErrorMessage(err.response.data.message)
      } else {
        action.setErrorMessage("Server error occured")
      }
    })
  }

  const checkValidation = () => {
    let emailValidation = decorator.validateEmail(form.email)
    action.isValidEmail(!!emailValidation)
    return !!emailValidation
  }

  const handleEmailInput = (e) => {
    action.handleEmail(e.target.value)
  }

  const toggleFormToLogin = () => {
    action.toggleFormType('login')
  }

  return (
    <Paper
      className={classes.root}
      elevation={8}
      component='form'>

      <Typography
        type='headline'
        component='h3'>
        Restore password
      </Typography>

      <FormInput
        type='text'
        name='Email'
        autoFocus={true}
        isValid={form.isValidEmail}
        onUserInput={handleEmailInput}
      />

      <Grid
        container
        alignItems={'center'}
        justify={'space-around'}
        className={classes.buttons}>

        <Button
          type='submit'
          color='primary'
          onClick={sendUserData}
          className={classes.button}>
          Submit
        </Button>
        <Button
          color={'secondary'}
          onClick={toggleFormToLogin}>
          Cancel
        </Button>
      </Grid>

      <SnackBar
        errorMessage={form.errorMessage}
        setErrorMessage={action.setErrorMessage}/>

    </Paper>
  )
}

const styles = {
  root: {
    padding: '4em 3em'
  },
  buttons: {
    marginTop: '1.5em',
  },
}

FormRestorePassword.propTypes = {
  validatePassword: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
  form: PropTypes.object.isRequired,
  action: PropTypes.object.isRequired
}

export default injectHash(
  injectValidation(
    withStyles(styles)(FormRestorePassword)
  )
)
