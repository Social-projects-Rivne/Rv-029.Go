import React, { Component } from 'react'
import validate from 'validate.js'

export default (FormComponent) => class InjectValidation extends Component {

  validateEmail = (email) => {
    const constrants = {
      from: {
        email: true
      }
    }

    return validate({from: email}, constrants) === undefined
  }

  validateName = (name) => {
    const constrants = {
      value: {
        format: {
          pattern:  /^[\w-]{2,15}$/
        }
      }
    }

    return validate({value: name}, constrants) === undefined
  }
  
  validatePassword = (password) => {
    const constrants = {
      value: {
        format: {
          pattern: /^[\w-]{6,15}$/
        }
      }
    }
    
    return validate({value: password}, constrants) === undefined
  }



  render() {
    return(
      <FormComponent {...this.props}
                     validateEmail={this.validateEmail}
                     validateName={this.validateName}
                     validatePassword={this.validatePassword} />
    )
  }
}


