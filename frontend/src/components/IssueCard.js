import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import axios from "axios"
import { API_URL } from "../constants/global"
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as issuesActions from "../actions/IssuesActions"
import * as sprintsActions from "../actions/SprintsActions"
import { withStyles } from 'material-ui/styles'
import Button from 'material-ui/Button'
import Chip from 'material-ui/Chip'
import DeleteIcon from 'material-ui-icons/Delete'
import PlayArrow from 'material-ui-icons/PlayArrow'
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
import messages from "../services/messages";

class IssueCard extends Component  {
  state = {
    updateIssueOpen: false,
    isIssueOpen: false,
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
    const { UUID } = this.props.data
    const { onUpdate } = this.props
    const { setErrorMessage, setNotificationMessage } = this.props.defaultPageActions
    const { issueName, issueDesc, issueEstimate, issueStatus } = this.props.issues

    axios.put(API_URL + `project/board/issue/update/${ UUID }`, {
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
    const { UUID } = this.props.data
    const { onUpdate } = this.props
    const { setErrorMessage, setNotificationMessage } = this.props.defaultPageActions
    this.setState({ isIssueOpen: false })

    axios.delete(API_URL + `project/board/issue/delete/${ UUID }`, {})
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

  addToSprint = () => {
    const todoSprints = this.props.sprints.currentSprints.filter((sprint, index) => {
      return sprint.Status == "TODO"
    })

    if (todoSprints.length == 0) {
      messages("There is no appropriate sprint to add the issue in. Sprint should have \"TODO\" status", false)
    } else {
      const {UUID} = this.props.data
      const { onUpdate } = this.props
      const sprint = todoSprints[todoSprints.length-1]

      this.setState({ isIssueOpen: false })

      axios.put(API_URL + `project/board/sprint/${sprint.ID}/add/issue/${ UUID }`, {})
        .then((response) => {
          messages(response.data.Message, true)
          onUpdate()
        })
        .catch((err) => {
          if (response.response && response.response.data.Message) {
              messages(response.response.data.Message, false)
          } else {
              messages("Server error occured", false)
          }
          this.handleClose()
        })
    }
  }

  toggleOpened = () => {
    this.setState({ isIssueOpen: !this.state.isIssueOpen })
  }

  render() {
    const { classes } = this.props
    const { Name, Description, Status, Estimate, SprintID } = this.props.data
    const { issueName, issueDesc, issueEstimate, issueStatus  } = this.props.issues

    const {
      setNameUpdateIssueInput,
      setDescUpdateIssueInput,
      setEstimateUpdateIssueInput,
      setStatusUpdateIssueInput
    } = this.props.issuesActions

    return (
      <ExpansionPanel expanded={this.state.isIssueOpen}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />} onClick={this.toggleOpened}>
          <Grid
            container
            alignItems={'center'}>
            <Grid style={{ marginRight: '1em' }}>
              <Chip label={Status} />
            </Grid>
            <Grid>
              <Typography>{Name}</Typography>
            </Grid>
          </Grid>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <Grid container >
            <Grid item xs={12}>
              <Typography type={'title'}>{ `Estimate: ${Estimate}` }</Typography>
              <Typography> {Description} </Typography>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>
        <ExpansionPanelDetails>
          <Grid
            container
            justify={'flex-end'}>
            <Grid item>
              {/*You can not update issue if it's already in finished sprint*/}
              {(this.props.sprints.currentSprint == null || (this.props.sprints.currentSprint != null && this.props.sprints.currentSprint.Status != "Done")) ? (
                  <IconButton onClick={this.handleOpenUpdateIssueClick}>
                      <EditIcon />
                  </IconButton>
              ) : (
                  ""
              )}
              {/*You can not remove issue if it's in sprint*/}
              {(SprintID != "00000000-0000-0000-0000-000000000000") ? (
                  ""
              ) : (
                  <IconButton
                      aria-label="Delete"
                      onClick={this.deleteIssue}>
                      <DeleteIcon />
                  </IconButton>
              )}
              {/*You can not remove issue if it's in sprint*/}
              {(SprintID != "00000000-0000-0000-0000-000000000000" || this.props.sprints.currentSprint != null) ? (
                  ""
              ) : (
                  <IconButton
                      aria-label="Add To Sprint"
                      onClick={this.addToSprint}>
                      <PlayArrow />
                  </IconButton>
              )}
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
                <MenuItem value={"TODO"}>TODO</MenuItem>
                <MenuItem value={"In Progress"}>In Progress</MenuItem>
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
    sprints: state.sprints,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    sprintsActions: bindActionCreators(sprintsActions, dispatch),
    issuesActions: bindActionCreators(issuesActions, dispatch)
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(IssueCard)
)
