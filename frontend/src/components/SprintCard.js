import React, { Component } from 'react'
import { Link } from 'react-router'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import axios from "axios"
import PropTypes from 'prop-types'
import { API_URL } from "../constants/global"
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as sprintsActions from "../actions/SprintsActions"
import { withStyles } from 'material-ui/styles'
import { FormControl } from 'material-ui/Form'
import { InputLabel } from 'material-ui/Input'
import { MenuItem } from 'material-ui/Menu'
import Button from 'material-ui/Button'
import Chip from 'material-ui/Chip'
import DeleteIcon from 'material-ui-icons/Delete'
import Grid from 'material-ui/Grid'
import IconButton from 'material-ui/IconButton'
import Select from 'material-ui/Select'
import TextField from 'material-ui/TextField'
import Typography from 'material-ui/Typography'
import EditIcon from 'material-ui-icons/ModeEdit';
import Card, {
  CardHeader,
  CardActions,
  CardContent
} from 'material-ui/Card'
import Dialog, {
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle
} from 'material-ui/Dialog'

class SprintCard extends Component {
  state = {
    updateSprintOpen: false,
  }

  static propTypes = {
    data: PropTypes.object.isRequired,
    onUpdate: PropTypes.func.isRequired,
    sprints: PropTypes.object.isRequired,
    classes: PropTypes.object.isRequired,
    sprintsActions: PropTypes.object.isRequired,
    defaultPageActions: PropTypes.object.isRequired
  }

  handleOpenUpdateSprintClick = () => {
    this.setState({ updateSprintOpen: true })
    this.props.sprintsActions.setCurrentSprint(this.props.data)
  }

  handleClose = () => {
    this.setState({ updateSprintOpen: false })
  }

  updateSprint = () => {
    const { id } = this.props.data
    const { onUpdate } = this.props
    const { sprintGoal, sprintDesc, sprintStatus } = this.props.sprints
    const { setNotificationMessage, setErrorMessage } = this.props.defaultPageActions

    axios.put(API_URL + `project/board/sprint/update/${ id }`, {
      goal: sprintGoal,
      desc: sprintDesc,
      status: sprintStatus
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

  deleteSprint = () => {
    const { id } = this.props.data
    const { onUpdate } = this.props
    const { setNotificationMessage, setErrorMessage } = this.props.defaultPageActions

    axios.delete(API_URL + `project/board/sprint/delete/${ id }`, {})
    .then((response) => {
      setNotificationMessage(response.data.Message)
      onUpdate()
    })
    .catch((error) => {
      if (error.response && error.response.data.Message) {
        setErrorMessage(error.response.data.Message)
      } else {
        setErrorMessage("Server error occured")
      }
    })
  }

  render() {
    const { classes } = this.props
    const { id, status, goal, created_at, description } = this.props.data
    const { sprintGoal, sprintDesc, sprintStatus } = this.props.sprints
    const {
      setGoalUpdateSprintInput,
      setDescUpdateSprintInput,
      setStatusUpdateSprintInput
    } = this.props.sprintsActions

    return (
      <Card className={classes.root}>
        <CardHeader
          className={classes.test}
          avatar={ <Chip label={status} /> }
          action={
            <Grid>
              {/* FIXME: horizontal scroll cause of this btn WTF? */}
              <IconButton onClick={this.handleOpenUpdateSprintClick}>
                <EditIcon />
              </IconButton>
              <IconButton onClick={this.deleteSprint}>
                <DeleteIcon />
              </IconButton>
            </Grid>
          }
          title={goal}
          subheader={created_at} />
        <CardContent>
          <Typography>{description}</Typography>
        </CardContent>
        <CardActions>
          <Link
            to={`${id}`}
            className={classes.link}>
            <Button
              size="small"
              color={'secondary'}>
              View
            </Button>
          </Link>
        </CardActions>

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
              value={sprintGoal}
              onChange={(e) => {setGoalUpdateSprintInput(e.target.value)}}
              fullWidth />
            <TextField
              margin="dense"
              id="desc"
              label="Description"
              type="text"
              value={sprintDesc}
              onChange={(e) => {setDescUpdateSprintInput(e.target.value)}}
              fullWidth />

            <FormControl className={classes.formControl}>
              <InputLabel htmlFor="age-simple">Status</InputLabel>
              <Select
                value={sprintStatus}
                onChange={(e) => {setStatusUpdateSprintInput(e.target.value)}}
                inputProps={{
                  name: 'status',
                  id: 'status-simple', }}>
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
  formControl: {
    width: '100%'
  },
  test: {
    maxWidth: '100%'
  }
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


