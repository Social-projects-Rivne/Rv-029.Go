import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import Snackbar from 'material-ui/Snackbar'
import IconButton from 'material-ui/IconButton'
import CloseIcon from 'material-ui-icons/Close'

class SimpleSnackbar extends React.Component {
  constructor(props) {
    super(props)
    
    this.state = {
      isOpen: false
    }
  }
  
  static propTypes = {
    classes: PropTypes.object.isRequired,
    setErrorMessage: PropTypes.func.isRequired,
    errorMessage: PropTypes.string
  }
  
  handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return
    }
    
    this.setState({
      isOpen: false
    })
  }
  
  clearMessage = () => {
    this.props.setErrorMessage(null)
  }
  
  componentWillReceiveProps(nextProps) {
    const res = nextProps.errorMessage

    this.setState({
      isOpen: !!res
    })
  }
  
  render() {
    const { classes } = this.props
    
    return (
      <div>
        <Snackbar
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'left',
          }}
          open={this.state.isOpen}
          autoHideDuration={6000}
          onClose={this.handleClose}
          SnackbarContentProps={{
            'aria-describedby': 'message-id',
          }}
          message={this.props.errorMessage}
          onExited={this.clearMessage}
          action={[
            <IconButton
              key="close"
              aria-label="Close"
              color="inherit"
              className={classes.close}
              onClick={this.handleClose} >
              <CloseIcon />
            </IconButton>,
          ]}
        />
      </div>
    )
  }
}

const styles = theme => ({
  close: {
    width: theme.spacing.unit * 4,
    height: theme.spacing.unit * 4,
  },
})

export default withStyles(styles)(SimpleSnackbar)