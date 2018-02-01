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

const FormRegister = ({ classes, form, action, ...decorator}) => {

  const sendUserData = (e) => {
    e.preventDefault()

    if (!checkValidation()) return

    axios.post('/auth/register', {
      email: form.email,
      name: form.name,
      surname: form.surname,
      password: decorator.MD5Encode(form.password)
    })
    .then((res) => {
      // TODO action after registration
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
    let nameValidation = decorator.validateName(form.name)
    let surnameValidation = decorator.validateName(form.surname)
    let passwordValidation = decorator.validatePassword(form.password)

    action.isValidEmail(!!emailValidation)
    action.isValidName(!!nameValidation)
    action.isValidSurname(!!surnameValidation)
    action.isValidPassword(!!passwordValidation)

    return !!emailValidation &&
           !!nameValidation &&
           !!surnameValidation &&
           !!passwordValidation
  }

  const handleEmailInput = (e) => {
    action.handleEmail(e.target.value)
  }

  const handleNameInput = (e) => {
    action.handleName(e.target.value)
  }

  const handleSurnameInput = (e) => {
    action.handleSurname(e.target.value)
  }

  const handlePasswordInput = (e) => {
    action.handlePassword(e.target.value)
  }

  return (
    <Paper
      className={classes.root}
      elevation={8}
      component='form'>

      <Typography
        type='headline'
        component='h3'>
        Registration
      </Typography>

      <FormInput
        type='text'
        name='Email'
        autoFocus={true}
        isValid={form.isValidEmail}
        onUserInput={handleEmailInput}
      />

      <FormInput
        type='text'
        name='Name'
        isValid={form.isValidName}
        onUserInput={handleNameInput}  
      />

      <FormInput
        type='text'
        name='Surname'
        isValid={form.isValidSurname}
        onUserInput={handleSurnameInput}
      />

      <FormInput
        type='password'
        name='Password'
        isValid={form.isValidPassword}
        onUserInput={handlePasswordInput}
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
            Log In
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
    padding: '4em 4em'
  },
  buttons: {
    marginTop: '1.5em',
  },
  link: {
    textDecoration: 'none'
  }
}

FormRegister.propTypes = {
  validateName: PropTypes.func.isRequired,
  validateEmail: PropTypes.func.isRequired,
  validatePassword: PropTypes.func.isRequired,
  MD5Encode: PropTypes.func.isRequired,
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
      connect(mapStateToProps, mapDispatchToProps)(FormRegister)
    )
  )
)
