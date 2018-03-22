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
import Menu, { MenuItem } from 'material-ui/Menu';
import Select from 'material-ui/Select'
import TextField from 'material-ui/TextField'
import Typography from 'material-ui/Typography'
import SettingsIcon from 'material-ui-icons/Settings';
import Tooltip from 'material-ui/Tooltip';
import { FormGroup, FormControlLabel } from 'material-ui/Form'
import Checkbox from 'material-ui/Checkbox'
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
import Avatar from 'material-ui/Avatar';
import FaceIcon from 'material-ui-icons/Face';
import messages from "../services/messages";

import List, { ListItem, ListItemText } from 'material-ui/List';
import ImageIcon from 'material-ui-icons/Image';
import WorkIcon from 'material-ui-icons/Work';
import BeachAccessIcon from 'material-ui-icons/BeachAccess';
import Divider from 'material-ui/Divider';



class IssueCard extends Component  {
  state = {
    updateIssueOpen: false,
    isIssueOpen: false,
    anchorEl: null,
    setSubTaskOpen: false,
    parent: null,
    checkedIssues: [],
    logInput: ''
  }

  static propTypes = {
    data: PropTypes.object.isRequired,
    onUpdate: PropTypes.func.isRequired,
    issues: PropTypes.object.isRequired,
    classes: PropTypes.object.isRequired,
    issuesActions: PropTypes.object.isRequired,
    defaultPageActions: PropTypes.object.isRequired
  }

  handleLogInput = (e) => {
    this.setState({
      logInput: e.target.value
    })
  }

  addLog = () => {
    const { UUID } = this.props.data
    const { onUpdate } = this.props
    const { setErrorMessage, setNotificationMessage } = this.props.defaultPageActions

    axios.put(API_URL + `project/board/issue/add_issue_log`, {
      issueID: this.props.data.UUID,
      userID: this.props.data.UUID, // TODO
      log: this.state.logInput
    })
      .then((res) => {
        this.props.onUpdate()
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

  handleOpenUpdateIssueClick = () => {
    this.setState({ updateIssueOpen: true })
    this.props.issuesActions.setCurrentIssue(this.props.data)
  }

  handleClose = () => {
    this.setState({
      updateIssueOpen: false,
      anchorEl: null,
      setSubTaskOpen: false,
      checkedSubTasks: [],
      logInput: ''
    })
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

    axios.delete(API_URL + `project/board/issue/delete/${ UUID }`, {})
    .then((res) => {
      setNotificationMessage(res.data.Message)
      this.setState({ isIssueOpen: false })
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

  handleCheckBoxChange = (child, parent) => e => {
    const emptyID = "00000000-0000-0000-0000-000000000000"

    let { checkedIssues } = this.state

    let toChange = {}

    if (e.target.checked) {
      toChange.parent = parent
      toChange.child = child
    } else {
      toChange.parent = emptyID
      toChange.child = child
    }

    // if checkbox was clicked twice
    for (let i = 0; i < checkedIssues.length; i++) {
      if (checkedIssues[i].child === toChange.child) {
        checkedIssues[i].parent = toChange.parent

        this.setState({ checkedIssues })
        return
      }
    }

    checkedIssues.push(toChange)
    this.setState({ checkedIssues })
  }


  setSubTasks = () => {
    axios({
      method: 'put',
      url: API_URL + `project/board/issue/set_parent`,
      data: this.state.checkedIssues
    })
      .then((res) => {
        this.props.onUpdate()

        this.setState({ checkedIssues: [] })
      })
      .catch((error) => {
        this.setState({ checkedIssues: [] })

        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occured")
        }
      })

    this.handleClose()
  }

  toggleOpened = () => {
    this.setState({ isIssueOpen: !this.state.isIssueOpen })
  }

  handleSettingsClick = e => {
    this.setState({ anchorEl: e.target })
  }

  handleSubTaskClick = () => {
    this.setState({
      anchorEl: null,
      setSubTaskOpen: true,
    })
  }

  render() {
    const { classes } = this.props
    const { Name, Description, Status, Estimate, SprintID, UUID, Nesting, Logs } = this.props.data
    const { issueName, issueDesc, issueEstimate, issueStatus } = this.props.issues

    const {
      setNameUpdateIssueInput,
      setDescUpdateIssueInput,
      setEstimateUpdateIssueInput,
      setStatusUpdateIssueInput
    } = this.props.issuesActions

    return (
      <ExpansionPanel expanded={this.state.isIssueOpen}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />} onClick={this.toggleOpened}>
          <Tooltip title={Nesting === '' ? '' : `Parents: ${Nesting}`}>
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
          </Tooltip>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <Grid container >
            <Grid item xs={12}>
              <Typography type={'body2'}>{Nesting === '' ? null : `Nesting: ${Nesting} ${Name}`}</Typography>
              <Typography type={'body2'}>{ `Estimate: ${Estimate}` }</Typography>
              <Typography> {Description} </Typography>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>

        <ExpansionPanelDetails>
          <Grid
            container
            spacing={0}
            justify={'flex-end'}>
            <Grid item>

              {(SprintID === "00000000-0000-0000-0000-000000000000") ? (
                <div className={classes.settings}>
                  <IconButton onClick={this.handleSettingsClick}
                    aria-haspopup="true"
                    aria-owns={this.state.anchorEl ? 'simple-menu' : null} >

                    <SettingsIcon />
                  </IconButton>

                  <Menu
                    id="simple-menu"
                    open={Boolean(this.state.anchorEl)}
                    anchorEl={this.state.anchorEl}
                    onClose={this.handleClose} >

                    <MenuItem onClick={ this.handleSubTaskClick }>sub tasks</MenuItem>
                  </Menu>
                </div>
              ) : ( "" )}

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

            <Grid container spacing={8}>
              <Grid item xs={12}>
                <TextField
                  id="logInput"
                  label="Log progress"
                  helperText="Any job you made working on this task"
                  onChange={this.handleLogInput}
                  value={this.state.logInput}
                  fullWidth
                  margin="normal" />

                <Button
                  onClick={this.addLog}
                  fullWidth
                  color={'primary'}>
                  send
                </Button>
              </Grid>

              {(this.props.sprints.issues) ? (

                <List>
                  <ListItem>
                    <Avatar>
                      <FaceIcon />
                    </Avatar>
                    <ListItemText primary="Photos" secondary="Jan 9, 2014" />
                  </ListItem>
                </List>

              ) : (null)}


            </Grid>
          </Grid>

        {/*</ExpansionPanelDetails>*/}
        {/*<ExpansionPanelDetails>*/}




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

        {/* ### Modal set sub tasks ### */}
        <Dialog
          open={this.state.setSubTaskOpen}
          onClose={this.handleClose}
          aria-labelledby="form-dialog-title" >
          <DialogTitle id="form-dialog-title">Set sub tasks</DialogTitle>
          <DialogContent>
            <FormGroup>

              {
                this.props.issues.currentIssues.map((item, i) => {

                  if (item.UUID === UUID) return null

                  else if (item.Parent === "00000000-0000-0000-0000-000000000000") {
                    return (

                      <FormControlLabel
                        key={i}
                        label={item.Name}
                        control={
                          <Checkbox
                            onChange={this.handleCheckBoxChange(item.UUID, UUID)} />
                        }/>

                    )
                  } else if (item.Parent === UUID) {

                    return (
                      <FormControlLabel
                        key={i}
                        label={item.Name}
                        control={
                          <Checkbox
                            defaultChecked // fixme: bug: requires double click
                            onChange={this.handleCheckBoxChange(item.UUID, UUID)} />
                        }/>
                    )
                  }
                })
              }

            </FormGroup>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Cancel
            </Button>
            <Button onClick={this.setSubTasks} color="primary">
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
  },
  settings: {
    display: 'inline-block'
  },
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
