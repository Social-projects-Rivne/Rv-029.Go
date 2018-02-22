import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import axios from "axios"
import { API_URL } from "../constants/global"
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as issuesActions from "../actions/IssuesActions"
import { withStyles } from 'material-ui/styles'
import Button from 'material-ui/Button'
import Chip from 'material-ui/Chip'
import DeleteIcon from 'material-ui-icons/Delete'
import ExpandMoreIcon from 'material-ui-icons/ExpandMore'
import { FormControl } from 'material-ui/Form'
import Grid from 'material-ui/Grid'
import IconButton from 'material-ui/IconButton'
import EditIcon from 'material-ui-icons/ModeEdit';
import { InputLabel } from 'material-ui/Input'
import { MenuItem } from 'material-ui/Menu'
import Select from 'material-ui/Select'
import TextField from 'material-ui/TextField'
import Typography from 'material-ui/Typography'
import Dialog, {
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle
} from 'material-ui/Dialog'
import ExpansionPanel, {
  ExpansionPanelSummary,
  ExpansionPanelDetails,
} from 'material-ui/ExpansionPanel'

class IssueCard extends Component  {
  state = {
    updateIssueOpen: false
  }

  static propTypes = {
    data: PropTypes.object.isRequired,
    onUpdate: PropTypes.func.isRequired,
    issues: PropTypes.object.isRequired,
    classes: PropTypes.object.isRequired,
    issuesActions: PropTypes.object.isRequired,
    defaultPageActions: PropTypes.object.isRequired
  }

  handleOpenUpdateIssueClick = () => {
    this.setState({ updateIssueOpen: true })
    this.props.issuesActions.setCurrentIssue(this.props.data)
  }

  handleClose = () => {
    this.setState({ updateIssueOpen: false })
  }

  updateIssue = () => {
    const { id } = this.props.data
    const { onUpdate } = this.props
    const { setErrorMessage, setNotificationMessage } = this.props.defaultPageActions
    const { issueName, issueDesc, issueEstimate, issueStatus } = this.props.issues

    axios.put(API_URL + `project/board/issue/update/${ id }`, {
      name: issueName,
      description: issueDesc,
      estimate: +issueEstimate,
      status: issueStatus
    })
    .then((res) => {
      setNotificationMessage(res.data.Message)
      onUpdate()
      this.handleClose()
    })
    .catch((err) => {
      if (err.response && err.response.data.Message) {
        setErrorMessage(err.response.data.Message)
      } else {
        setErrorMessage("Server error occured")
      }
      this.handleClose()
    })
  }

  deleteIssue = () => {
    const { id } = this.props.data
    const { onUpdate } = this.props
    const { setErrorMessage, setNotificationMessage } = this.props.defaultPageActions

    axios.delete(API_URL + `project/board/issue/delete/${ id }`, {})
    .then((res) => {
      setNotificationMessage(res.data.Message)
      onUpdate()
    })
    .catch((err) => {
      if (err.response && err.response.data.Message) {
        setErrorMessage(err.response.data.Message)
      } else {
        setErrorMessage("Server error occured")
      }
      this.handleClose()
    })
  }

  render() {
    const { classes } = this.props
    const { name, description, status, estimate } = this.props.data
    const { issueName, issueDesc, issueEstimate, issueStatus  } = this.props.issues
    const {
      setNameUpdateIssueInput,
      setDescUpdateIssueInput,
      setEstimateUpdateIssueInput,
      setStatusUpdateIssueInput
    } = this.props.issuesActions

    return (
      <ExpansionPanel>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <Grid
            container
            alignItems={'center'}>
            <Grid style={{ marginRight: '1em' }}>
              <Chip label={status} />
            </Grid>
            <Grid>
              <Typography>{name}</Typography>
            </Grid>
          </Grid>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <Grid container >
            <Grid item xs={12}>
              <Typography type={'title'}>{ `Estimate: ${estimate}` }</Typography>
              <Typography> {description} </Typography>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>
        <ExpansionPanelDetails>
          <Grid
            container
            justify={'flex-end'}>
            <Grid item>
              <IconButton onClick={this.handleOpenUpdateIssueClick}>
                <EditIcon />
              </IconButton>
              <IconButton
                aria-label="Delete"
                onClick={this.deleteIssue}>
                <DeleteIcon />
              </IconButton>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>

        <Dialog
          open={this.state.updateIssueOpen}
          onClose={this.handleClose}
          aria-labelledby="form-dialog-title" >
          <DialogTitle id="form-dialog-title">Update Sprint</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Please, fill required fields.
            </DialogContentText>
            <TextField
              autoFocus
              margin="dense"
              id="name"
              label="Name"
              type="text"
              value={issueName}
              onChange={(e) => {setNameUpdateIssueInput(e.target.value)}}
              fullWidth />
            <TextField
              margin="dense"
              id="desc"
              label="Description"
              type="text"
              value={issueDesc}
              onChange={(e) => {setDescUpdateIssueInput(e.target.value)}}
              fullWidth />
            <TextField
              margin="dense"
              id="estimate"
              label="Estimate"
              type="text"
              value={issueEstimate}
              onChange={(e) => {setEstimateUpdateIssueInput(e.target.value)}}
              fullWidth />

            <FormControl className={classes.formControl}>
              <InputLabel htmlFor="age-simple">Status</InputLabel>
              <Select
                value={issueStatus}
                onChange={(e) => {setStatusUpdateIssueInput(e.target.value)}}
                inputProps={{
                  name: 'status',
                  id: 'status-simple',
                }} >
                <MenuItem value={"Todo"}>Todo</MenuItem>
                <MenuItem value={"In process"}>In process</MenuItem>
                <MenuItem value={"Done"}>Done</MenuItem>
              </Select>
            </FormControl>

          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Cancel
            </Button>
            <Button onClick={this.updateIssue} color="primary">
              Ok
            </Button>
          </DialogActions>
        </Dialog>

      </ExpansionPanel>
    )
  }
}

const styles = {
  formControl: {
    width: '100%'
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    defaultPage: state.defaultPage,
    issues: state.issues,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    issuesActions: bindActionCreators(issuesActions, dispatch)
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(IssueCard)
)
