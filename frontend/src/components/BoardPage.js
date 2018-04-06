import React, { Component } from 'react'
import SprintCard from "./SprintCard"
import {API_URL} from "../constants/global"
import IssueCard from "./IssueCard"
import InjectTransformIssues from '../decorators/transformIssues'
import PropTypes from 'prop-types'
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as boardsActions from "../actions/BoardsActions"
import * as sprintsActions from "../actions/SprintsActions"
import * as issuesActions from "../actions/IssuesActions"
import * as projectsActions from "../actions/ProjectsActions"
import * as usersActions from '../actions/UsersActions'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import messages from "../services/messages"
import axios from "axios"
import Grid from 'material-ui/Grid'
import { withStyles } from 'material-ui/styles'
import Typography from 'material-ui/Typography'
import Button from 'material-ui/Button'
import AddIcon from 'material-ui-icons/Add'
import TextField from 'material-ui/TextField';
import Avatar from 'material-ui/Avatar';
import PersonIcon from 'material-ui-icons/Person';
import DeleteIcon from 'material-ui-icons/Delete';
import List, { ListItem, ListItemAvatar, ListItemText } from 'material-ui/List';
import Chip from 'material-ui/Chip';
import FaceIcon from 'material-ui-icons/Face';
import { Link, browserHistory} from 'react-router'
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
    addUserOpen: false,
    selectedEmail : null,
  }

  handleClickOpenAddUser = () => {
      this.setState({ addUserOpen: true })
  };

  handleClickOpenCreateIssue = () => {
    this.props.boardsActions.resetInput()
    this.setState({ createIssueOpen: true })
  }

  handleClickOpenCreateSprint = () => {
    this.props.boardsActions.resetInput()
    this.setState({ createSprintOpen: true })
  }

  handleClose = () => {
    this.setState({
      createIssueOpen: false,
      createSprintOpen: false,
      addUserOpen: false
    });
  };

  componentDidMount() {
    this.props.sprintsActions.setCurrentSprint(null)
    this.getSprintsList()
    this.getIssuesList()
    this.getCurrentBoard()
  }

  getSprintsList = () => {
    axios.get(API_URL + `project/board/${this.props.ownProps.params.id}/sprint/list`)
      .then((response) => {{this.props.issues.currentIssues.map((item, i) => (
                  <IssueCard
                      key={i}
                      data={item}
                      onUpdate={this.getIssuesList}
                  />
              ))}
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
    const { issuesActions, transformIssues } = this.props

    axios.get(API_URL + `project/board/${this.props.ownProps.params.id}/issue/list`)
      .then((response) => {
        console.log(response.data.Data)
        issuesActions.setIssues(transformIssues(response.data.Data))
      })
      .catch((error) => {
        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occured")
        }
      })
  }

  getCurrentBoard = () => {
      axios.get(API_URL + `project/board/select/${this.props.ownProps.params.id}`)
          .then((response) => {
              this.props.boardsActions.setCurrentBoard(response.data.Data)

              this.getProjectUsers()
          })
          .catch((error) => {
              if (error.response && error.response.data.Message) {
                  messages(error.response.data.Message)
              } else {
                  messages("Server error occured")
              }
          });
  }

  findAvailableUsers = () => {

    let { currentProjectUsers } = this.props.projects,
        { Users } = this.props.boards.currentBoard,
        assignedUsers = []

    for (let key in Users) {
      assignedUsers.push(key)
    }

    return currentProjectUsers.filter((item) => {
      return !assignedUsers.includes(item.UUID)
    })
  }

  findAssignedUsers = () => {

    const { Users } = this.props.boards.currentBoard || {},
          { currentProjectUsers } = this.props.projects

    let assignedUsers = []

    if (Users) {
      currentProjectUsers.forEach((item) => {
        if (Users[item.UUID]) {
          assignedUsers.push(item)
        }
      })
    }

    return assignedUsers
  }

  getProjectUsers = () => {
      axios.get(API_URL + `project/${this.props.boards.currentBoard.ProjectID}/users`)
          .then((response) => {
              this.props.projectsActions.setProjectUsers(response.data.Data)

            this.props.usersActions.setAvailableUsers(this.findAvailableUsers())
            this.props.usersActions.setAssignedUsers(this.findAssignedUsers())
          })
          .catch((error) => {
              if (error.response && error.response.data.Message) {
                  messages(error.response.data.Message)
              } else {
                  messages("Server error occured")
              }
          });
  }

   addUserToBoard = (user) => {
        axios.post(API_URL + `project/board/assign-user/${this.props.ownProps.params.id}`,{
            email: user.Email,
            user_id: user.UUID,
        })
            .then(() => {
                // TODO append to redux
                this.getCurrentBoard()
                this.handleClose()
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            });
    }

    deleteUserFromBoard = (userId) => () => {
        axios.delete(API_URL + `project/board/${this.props.ownProps.params.id}/user/${userId}`)
            .then(() => {
                // TODO append to redux
                this.getCurrentBoard()
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            });
    }

  createIssue = () => {
    axios.post(API_URL + `project/board/${this.props.ownProps.params.id}/issue/create`, {
      name: this.props.boards.nameInput,
      description: this.props.boards.descInput,
      estimate: +this.props.boards.estimation,
      status: 'TODO'
    })
    .then((response) => {
        // TODO append to redux
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
      // TODO append to redux
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
              item xs={3}
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
                          onClick={this.handleClickOpenAddUser}
                          className={classes.button}>
                          <AddIcon />
                      </Button>
                  </Grid>

                  <Grid>
                      <Typography type={"headline"} className={classes.pageTitle}>Users</Typography>
                  </Grid>

              </Grid>

          {(this.props.users.assignedUsers) ? (
            this.props.users.assignedUsers.map((item, i) => (

              <Chip
                // todo: link to user page
                onClick={this.handleClose}
                key={i}
                label={ <Link className={classes.link} to={'/profile/'+item.UUID}>{item.FirstName} {item.LastName}</Link>}
                onDelete={this.deleteUserFromBoard(item.UUID)}
                className={classes.chip}
                avatar={
                  <Avatar>
                    <FaceIcon />
                  </Avatar>
                } />

            ))
          ) : (null)}

          </Grid>

        <Grid
          item xs={3}
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

          { (this.props.issues.currentIssues) ? (

            this.props.issues.currentIssues.map((item, i) => (
              <IssueCard
                key={i}
                data={item}
                onUpdate={this.getIssuesList} />
              ))

            ) : (null)
           }

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

        {/*/!* #################### MODAL USER #################### *!/*/}
        <Dialog
            open={this.state.addUserOpen}
            onClose={this.handleClose}
            aria-labelledby="form-dialog-title" >
            <DialogTitle id="form-dialog-title">Add user</DialogTitle>
            <DialogContent>
                <List>
                    {(this.props.users.availableUsers) ?
                      (this.props.users.availableUsers.map((item, i) => (
                        <ListItem button onClick={() => this.addUserToBoard(item)} key={i}>
                            <ListItemAvatar>
                                <Avatar className={classes.avatar}>
                                    <PersonIcon />
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText primary={ `${item.FirstName} ${item.LastName}` }  />
                        </ListItem>
                    ))) : (null)}
                </List>
            </DialogContent>
            <DialogActions>
                <Button onClick={this.handleClose} color="primary">
                    Cancel
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
    textDecoration: 'none',
  },
  pageTitle: {
    color: '#fff',
  },
  titleGrid: {
    margin: '1em 0'
  },
  button: {
    marginRight: '1em'
  },
  chip: {
    background: '#fff',
    marginBottom: '.5em',
    marginRight: '.5em',
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    defaultPage: state.defaultPage,
    sprints: state.sprints,
    boards: state.boards,
    issues: state.issues,
    projects: state.projects,
    users: state.users,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    sprintsActions: bindActionCreators(sprintsActions, dispatch),
    issuesActions: bindActionCreators(issuesActions, dispatch),
    boardsActions: bindActionCreators(boardsActions, dispatch),
    projectsActions: bindActionCreators(projectsActions, dispatch),
    usersActions: bindActionCreators(usersActions, dispatch)
  }
}

export default InjectTransformIssues(
  withStyles(styles)( connect(mapStateToProps, mapDispatchToProps)(BoardPage) )
)
