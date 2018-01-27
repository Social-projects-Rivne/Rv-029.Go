import React, { Component } from 'react';
import validate from 'validate.js';

export default (FormComponent) => class InjectValidation extends Component {

  validateEmail = (email) => {
    const constrants = {
      from: {
        email: true
      }
    };

    return validate({from: email}, constrants) === undefined; 
  };
  
  validateInput = (string) => {
    const constrants = {
      value: {
        format: {
          pattern: /^[\w-]{6,15}$/
        }
      }
    };
    
    return validate({value: string}, constrants) === undefined;
  };

  render() {
    return(
      <FormComponent {...this.props} 
        validateEmail={this.validateEmail}
        validateInput={this.validateInput} />
    )
  }
}


