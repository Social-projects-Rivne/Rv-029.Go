import React, { Component } from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { Link } from 'react-router'
import * as formActions from '../actions/FormActions'
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

    axios.post('/auth/forget-password', {
      email: form.email,
    })
    .then((res) => {
      //TODO restore password
      console.log(res)
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
        <Link to={'/authorization/login'}
          className={classes.link}>
          <Button
            color={'secondary'}>
            Cancel
          </Button>
        </Link>
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
  link: {
    textDecoration: 'none'
  }
}

FormRestorePassword.propTypes = {
  validatePassword: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
  form: PropTypes.object.isRequired,
  action: PropTypes.object.isRequired
}

const mapStateToProps = (state) => {
  return {
    form: state.form
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    action: bindActionCreators(formActions, dispatch)
  }
}

export default injectHash(
  injectValidation(
    withStyles(styles)(
      connect(mapStateToProps, mapDispatchToProps)(FormRestorePassword)
    )
  )
)
