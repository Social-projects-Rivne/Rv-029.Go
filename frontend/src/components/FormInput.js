import React, { Component } from 'react'
import PropTypes from 'prop-types'
import TextField from 'material-ui/TextField'
import Grid from 'material-ui/Grid'

const FormInput = (props) => {

  const handleChange = (e) => {
    props.onUserInput(e)
  }

  const message = props.isValid ? "" :
    `${props.name} is not valid`

  return(
    <Grid item xs={12}>
      <TextField
        margin="normal"
        autoFocus={props.autoFocus || false}
        error={!props.isValid}
        helperText={message}
        label={props.name}
        type={props.type}
        onChange={handleChange}/>
    </Grid>
  )
}

FormInput.propTypes = {
  type: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  autoFocus: PropTypes.bool,
  isValid: PropTypes.bool.isRequired,
  onUserInput: PropTypes.func.isRequired,
}

export default FormInput

