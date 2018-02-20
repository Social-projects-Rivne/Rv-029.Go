
// TODO: initial text when update
// TODO: update sprints list after single sprint update

import React, { Component } from 'react'
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as sprintsActions from "../actions/SprintsActions"
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import axios from "axios"
import Card, { CardHeader, CardActions, CardContent } from 'material-ui/Card'
import Button from 'material-ui/Button'
import { Link, browserHistory } from 'react-router'
import Typography from 'material-ui/Typography'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import Chip from 'material-ui/Chip';
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

class SprintCard extends Component {
  state = {
    anchorEl: null,
    updateSprintOpen: false
  }

  handleOpenUpdateSprintClick = () => {
    this.setState({
      updateSprintOpen: true
    })
    console.log(this.props.data)
    this.props.sprintsActions.setCurrentSprint(this.props.data)
  }

  handleClose = () => {
    this.setState({ updateSprintOpen: false });
  };

  updateSprint = () => {
    axios.put(API_URL + `project/board/sprint/update/${this.props.id}`, {
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
    const { classes } = this.props;
    const { open } = this.state;
    const { anchorEl } = this.state;

    return (

      <Card className={classes.root}>
        <CardHeader
          avatar={
            <Chip label={this.props.status} />
          }
          action={
            <Grid item>
              {/* FIXME: horizontal scroll cause of this btn WTF? */}
              <IconButton onClick={this.handleOpenUpdateSprintClick}>
                <Icon> edit_icon</Icon>
              </IconButton>
              <IconButton>
                <DeleteIcon />
              </IconButton>
            </Grid>
          }
          title={this.props.title}
          subheader={this.props.date} />
        <CardContent>
          <Typography>{this.props.desc}</Typography>
        </CardContent>
        <CardActions>
          <Link
            // to="view_board"
            className={this.props.classes.link}>
            <Button
              size="small"
              color={'secondary'}>
              View
            </Button>
          </Link>
        </CardActions>

        {/* #################### MODAL UPDATE SPRINT #################### */}
        <Dialog
          open={this.state.updateSprintOpen}
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
            <Button onClick={this.updateSprint} color="primary">
              Ok
            </Button>
          </DialogActions>
        </Dialog>

      </Card>
    )
  }
}

const styles = {
  root: {
    marginBottom: '1em'
  },
  link: {
    textDecoration: 'none'
  },
}


const mapStateToProps = (state, ownProps) => {
  return {
    defaultPage: state.defaultPage,
    sprints: state.sprints,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    sprintsActions: bindActionCreators(sprintsActions, dispatch),
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(SprintCard)
)


