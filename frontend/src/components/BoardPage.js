import React, { Component } from 'react'
import SprintCard from "./SprintCard"
import {API_URL} from "../constants/global"
import IssueCard from "./IssueCard"
import PropTypes from 'prop-types'
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as boardsActions from "../actions/BoardsActions"
import * as sprintsActions from "../actions/SprintsActions"
import * as issuesActions from "../actions/IssuesActions"
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import messages from "../services/messages"
import axios from "axios"
import Grid from 'material-ui/Grid'
import { withStyles } from 'material-ui/styles'
import Typography from 'material-ui/Typography'
import Button from 'material-ui/Button'
import AddIcon from 'material-ui-icons/Add'
import TextField from 'material-ui/TextField'
import { FormGroup, FormControlLabel } from 'material-ui/Form'
import Checkbox from 'material-ui/Checkbox'
import Dialog, {
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from 'material-ui/Dialog'

class BoardPage extends Component{
  state = {
    createIssueOpen: false,
    createSprintOpen: false,
    setSubTaskOpen: false,
    checkedSubTasks: [],
    parent: null,
  }

  handleClickOpenCreateIssue = () => {
    this.props.boardsActions.resetInput()
    this.setState({ createIssueOpen: true })
  }

  handleClickOpenCreateSprint = () => {
    this.props.boardsActions.resetInput()
    this.setState({ createSprintOpen: true })
  }

  handleSetSubTaskClick = parent => () => {
    this.setState({
      setSubTaskOpen: true,
      parent
    })
  }

  handleCheckBoxChange = id => e => {
    let { checkedSubTasks } = this.state

    if (e.target.checked) {
      checkedSubTasks.push(id)
      this.setState({ checkedSubTasks })
    } else {
      checkedSubTasks.splice(checkedSubTasks.indexOf(id), 1)
      this.setState({ checkedSubTasks })
    }
  }


  setSubTasks = () => {
    axios.put(API_URL + `project/board/issue/set_parent`, {
      issues: this.state.checkedSubTasks,
      parent: this.state.parent
    })
      .then((res) => {
        this.getIssuesList()
      })
      .catch((error) => {
        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occured")
        }
      })

    this.handleClose()
  }

  handleClose = () => {
    this.setState({
      createIssueOpen: false,
      createSprintOpen: false,
      setSubTaskOpen: false,
      checkedSubTasks: [],
      parent: null
    })
  }
  
  componentDidMount() {
    this.props.sprintsActions.setCurrentSprint(null)
    this.getSprintsList()
    this.getIssuesList()
  }

  getSprintsList = () => {
    axios.get(API_URL + `project/board/${this.props.ownProps.params.id}/sprint/list`)
      .then((response) => {
        this.props.sprintsActions.setSprints(response.data.Data)
      })
      .catch((error) => {
        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occured")
        }
      })
  }

  getIssuesList = () => {
    axios.get(API_URL + `project/board/${this.props.ownProps.params.id}/issue/list`)
      .then((response) => {
        this.props.issuesActions.setIssues(response.data.Data)
      })
      .catch((error) => {
        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occured")
        }
      })
  }

  createIssue = () => {
    axios.post(API_URL + `project/board/${this.props.ownProps.params.id}/issue/create`, {
      name: this.props.boards.nameInput,
      description: this.props.boards.descInput,
      user_id:'9646324a-0aa2-11e8-ba34-b06ebf83499f', // debug
      estimate: +this.props.boards.estimation,
      status: 'Todo'
    })
    .then((response) => {
      this.getIssuesList()
      this.handleClose()
    })
    .catch((error) => {
      if (error.response && error.response.data.Message) {
        this.props.defaultPageActions.setErrorMessage(error.response.data.Message)
      } else {
        this.props.defaultPageActions.setErrorMessage("Server error occured")
      }
      this.handleClose()
    })
  }

  createSprint = () => {
    axios.post(API_URL + `project/board/${this.props.ownProps.params.id}/sprint/create`, {
      goal: this.props.boards.goalInput,
      desc: this.props.boards.descInput
    })
    .then((response) => {
      this.props.defaultPageActions.setNotificationMessage(response.data.Message)
      this.getSprintsList()
      this.handleClose()
    })
    .catch((error) => {
      if (error.response && error.response.data.Message) {
        this.props.defaultPageActions.setErrorMessage(error.response.data.Message)
      } else {
        this.props.defaultPageActions.setErrorMessage("Server error occured")
      }
      this.handleClose()
    })
  }

  render() {
    const { classes } = this.props
    return (
      <Grid
        container
        spacing={0}>

        <Grid
          item xs={6}
          className={classes.left}>

          <Grid
            container
            spacing={0}
            alignItems={'flex-end'}
            className={classes.titleGrid}>

            <Grid>
              <Button
                fab
                raised={true}
                mini={true}
                onClick={this.handleClickOpenCreateIssue}
                className={classes.button}>
                <AddIcon />
              </Button>
            </Grid>

            <Grid>
              <Typography type={"headline"} className={classes.pageTitle}>Backlog</Typography>
            </Grid>

          </Grid>

          {this.props.issues.currentIssues.map((item, i) => (
            <IssueCard
              key={i}
              data={item}
              onUpdate={this.getIssuesList}
              setSubTaskClick={this.handleSetSubTaskClick} />
          ))}

        </Grid>

        <Grid item xs={6} className={classes.right}>

          <Grid
            container
            spacing={0}
            alignItems={'flex-end'}
            className={classes.titleGrid}>

            <Grid>
              <Button
                fab
                raised={true}
                mini={true}
                onClick={this.handleClickOpenCreateSprint}
                className={classes.button}>
                <AddIcon />
              </Button>
            </Grid>

            <Grid>
              <Typography className={classes.pageTitle} type={"headline"}>Sprints</Typography>
            </Grid>

          </Grid>

          {this.props.sprints.currentSprints.map((item, i) => (
            <SprintCard
              key={i}
              data={item}
              onUpdate={this.getSprintsList} />
          ))}

        </Grid>

        {/* ### Modal issue ### */}
        <Dialog
          open={this.state.createIssueOpen}
          onClose={this.handleClose}
          aria-labelledby="form-dialog-title" >
          <DialogTitle id="form-dialog-title">Create Issue</DialogTitle>
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
              onChange={(e) => {this.props.boardsActions.setName(e.target.value)}}
              fullWidth />
            <TextField
              margin="dense"
              id="desc"
              label="Description"
              type="text"
              onChange={(e) => {this.props.boardsActions.setDesc(e.target.value)}}
              fullWidth />
            <TextField
              margin="dense"
              id="est"
              label="Estimation"
              onChange={(e) => {this.props.boardsActions.setEstimation(e.target.value)}}
              type="text"
              fullWidth />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Cancel
            </Button>
            <Button onClick={this.createIssue} color="primary">
              Ok
            </Button>
          </DialogActions>
        </Dialog>

        {/* ### Modal create sprint ### */}
        <Dialog
          open={this.state.createSprintOpen}
          onClose={this.handleClose}
          aria-labelledby="form-dialog-title" >
          <DialogTitle id="form-dialog-title">Create Sprint</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Please, fill required fields.
            </DialogContentText>
            <TextField
              autoFocus
              margin="dense"
              id="goal"
              label="Goal"
              type="text"
              onChange={(e) => {this.props.boardsActions.setGoal(e.target.value)}}
              fullWidth />
            <TextField
              margin="dense"
              id="desc"
              label="Description"
              type="text"
              onChange={(e) => {this.props.boardsActions.setDesc(e.target.value)}}
              fullWidth />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Cancel
            </Button>
            <Button onClick={this.createSprint} color="primary">
              Ok
            </Button>
          </DialogActions>
        </Dialog>

        { /* TODO remove */ }
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
                if (item.Parent === "00000000-0000-0000-0000-000000000000") {
                  return (

                    <FormControlLabel
                      key={i}
                      label={item.Name}
                      control={
                      <Checkbox
                        onChange={this.handleCheckBoxChange(item.UUID)} />
                      }/>

                    )
                  } else {

                  return (
                    <FormControlLabel
                      key={i}
                      label={item.Name}
                      control={
                        <Checkbox
                          defaultChecked={true}
                          onChange={this.handleCheckBoxChange(item.UUID)} />
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

      </Grid>
    )
  }
}

const styles = {
  left: {
    padding: '0 1em 0 2em'
  },
  right: {
    padding: '0 2em 0 1em'
  },
  link: {
    textDecoration: 'none'
  },
  pageTitle: {
    color: '#fff',
  },
  titleGrid: {
    margin: '1em 0'
  },
  button: {
    marginRight: '1em'
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    defaultPage: state.defaultPage,
    sprints: state.sprints,
    boards: state.boards,
    issues: state.issues,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    sprintsActions: bindActionCreators(sprintsActions, dispatch),
    issuesActions: bindActionCreators(issuesActions, dispatch),
    boardsActions: bindActionCreators(boardsActions, dispatch)
  }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(BoardPage)
)
