import React from 'react'
import PropTypes from 'prop-types'
import Button from 'material-ui/Button'
import Dialog, {
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  withMobileDialog,
} from 'material-ui/Dialog'

class ResponsiveDialog extends React.Component {
  state = {
    open: false,
  }

  static propTypes = {
    fullScreen: PropTypes.bool.isRequired,
    setNotificationMessage: PropTypes.func.isRequired,
    title: PropTypes.string,
    content: PropTypes.string
  }

  componentWillReceiveProps(nextProps) {
    this.state.open = !!nextProps.content
  }

  handleClose = () => {
    this.setState({ open: false })

    this.props.setNotificationMessage(null)
  }

  render() {
    const { fullScreen } = this.props

    return (
      <div>
        <Dialog
          fullScreen={fullScreen}
          open={this.state.open}
          onClose={this.handleClose}
          aria-labelledby="responsive-dialog-title" >

          <DialogTitle id="responsive-dialog-title">{this.props.title}</DialogTitle>
          <DialogContent>
            <DialogContentText>
              {this.props.content}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary" autoFocus>
              Ok
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    )
  }
}

export default withMobileDialog()(ResponsiveDialog)