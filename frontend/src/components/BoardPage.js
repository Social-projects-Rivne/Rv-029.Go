import React, { Component } from 'react'
import SprintCard from "./SprintCard"
import {API_URL} from "../constants/global"
import PropTypes from 'prop-types'
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as boardActions from "../actions/BoardsActions"
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import messages from "../services/messages"
import axios from "axios"
import Grid from 'material-ui/Grid'
import { withStyles } from 'material-ui/styles'
import ExpansionPanel, {
  ExpansionPanelSummary,
  ExpansionPanelDetails,
} from 'material-ui/ExpansionPanel'
import Typography from 'material-ui/Typography'
import ExpandMoreIcon from 'material-ui-icons/ExpandMore'
import Button from 'material-ui/Button'
import AddIcon from 'material-ui-icons/Add'
import IconButton from 'material-ui/IconButton';
import DeleteIcon from 'material-ui-icons/Delete';
import Icon from 'material-ui/Icon';
import TextField from 'material-ui/TextField';
import Dialog, {
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from 'material-ui/Dialog';

class BoardPage extends Component{
  state = {
    createIssueOpen: false,
    createSprintOpen: false,
    goal: "",
    name: "",
    desc: "",
    estimation: ""
  }

  componentWillMount() {
    axios.get(API_URL + `project/${this.props.projects.currentProject}/board/list`)
    .then((response) => {
      this.props.boardActions.setBoards(response.data)
    })
    .catch((error) => {
      if (error.response && error.response.data.Message) {
        messages(error.response.data.Message)
      } else {
        messages("Server error occured")
      }
    });
  }

  handleClickOpenCreateIssue = () => {
    this.setState({ createIssueOpen: true });
  };

  handleClickOpenCreateSprint = () => {
    this.setState({ createSprintOpen: true })
  }

  handleClose = () => {
    this.setState({
      createIssueOpen: false,
      createSprintOpen: false
    });
  };

  createIssue = () => {
    axios.post(API_URL + `project/board/${this.props.ownProps.params.id}/issue/create`, {
      name: this.props.name,
      desc: this.props.desc,
      estimate: this.props.estimation,
      user_id: 'userID'
    })
    .then((response) => {
      console.log(response.data.status)
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
      goal: this.state.goal,
      desc: this.state.desc
    })
    .then((response) => {
      console.log(response.data.status)
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

          <ExpansionPanel>
            <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
              <Typography>Issue 1</Typography>
            </ExpansionPanelSummary>
            <ExpansionPanelDetails>
              <Typography>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse malesuada lacus ex,
                sit amet blandit leo lobortis eget.
              </Typography>
            </ExpansionPanelDetails>
            <ExpansionPanelDetails>
              <Grid
                container
                justify={'flex-end'}>
                <Grid item>
                  <IconButton aria-label="Delete">
                    <DeleteIcon />
                  </IconButton>
                  <IconButton>
                    <Icon>edit_icon</Icon>
                  </IconButton>
                </Grid>
              </Grid>
            </ExpansionPanelDetails>
          </ExpansionPanel>

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
          {this.props.boards.currentBoards.map((item, i) => (
            <SprintCard
              key={i}
              title={item.name}
              date={item.created_at}
              desc={item.description} />
          ))}

        </Grid>

        {/* #################### MODAL ISSUE #################### */}
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
              onChange={(e) => { this.setState({ name: e.target.value}) }}
              fullWidth />
            <TextField
              margin="dense"
              id="desc"
              label="Description"
              type="text"
              onChange={(e) => { this.setState({ description: e.target.value}) }}
              fullWidth />
            <TextField
              margin="dense"
              id="est"
              label="Estimation"
              onChange={(e) => { this.setState({ estimation: e.target.value}) }}
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

        {/* #################### MODAL SPRINT #################### */}
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
              onChange={(e) => { this.setState({ goal: e.target.value }) }}
              fullWidth />
            <TextField
              margin="dense"
              id="desc"
              label="Description"
              type="text"
              onChange={(e) => { this.setState({ desc: e.target.value }) }}
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
    projects: state.projects,
    boards: state.boards,
    defaultPage: state.defaultPage,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    boardActions: bindActionCreators(boardActions, dispatch)
  }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(BoardPage)
)
