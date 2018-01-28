import React, { Component } from 'react';
import PropTypes from 'prop-types';
import TextField from 'material-ui/TextField';
import Grid from 'material-ui/Grid';

export default class FormInput extends Component {
  constructor(props) {
    super(props)
  }

  static propTypes = {
    isValid: PropTypes.bool.isRequired,
    type: PropTypes.string.isRequired,
    name: PropTypes.string.isRequired,
    onUserInput: PropTypes.func.isRequired,
  };

  handleChange = (e) => {
    this.props.onUserInput(e)
  };

  render() {

    // TODO make inputs more flexible
    // implement this.props.message

    const message = this.props.isValid ? "" :
      `${this.props.name} is not valid`;
    
    return(
      <Grid item xs={12}>
        <TextField
          autoFocus={this.props.autoFocus || false}
          error={!this.props.isValid}
          helperText={message}
          margin="normal"
          label={this.props.name}
          type={this.props.type} 
          onChange={this.handleChange}/>
      </Grid>
    )
  }
}
