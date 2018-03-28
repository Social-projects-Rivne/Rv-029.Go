import React, { Component } from 'react';
import Snackbar, { SnackbarContent } from 'material-ui/Snackbar';
import { withStyles } from 'material-ui/styles'

class SnackBar extends Component {
  state = {
    open: false,
    message: '',
    status: null
  };

  componentWillReceiveProps(next) {

    if (next.options) {

      this.setState({
        open: !!next.options,
        message: next.options.message,
        status: next.options.status
      })
    }
  }

  handleClose = () => {
    this.setState({
      open: false,
    });
  };

  render() {

    const { status } = this.state,
          { classes } = this.props

    return (

      <Snackbar
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'left',
        }}
        open={this.state.open}
        onClose={this.handleClose}
        SnackbarContentProps={{
          'aria-describedby': 'message-id',
          classes: {
            root: status ? classes.rootSuccess : classes.rootFailed
          }
        }}
        message={<span id="message-id">{ this.state.message }</span>} />

    )
  }
}

const styles = {
  rootSuccess: {
    background: '#abe9ff',
    color: 'black'
  },
  rootFailed: {
    background: '#f799a0',
    color: 'black'
  }
}

export default withStyles(styles)(SnackBar)