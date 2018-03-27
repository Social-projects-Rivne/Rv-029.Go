import React, { Component } from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import Users from '../components/scrumPocker/Users'
import Issue from '../components/scrumPocker/Issue'
import Estimation from '../components/scrumPocker/Estimation'
import {withStyles} from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import * as projectsActions from "../actions/ProjectsActions";
import * as defaultPageActions from "../actions/DefaultPageActions";
import * as sprintsActions from "../actions/SprintsActions";

class EstimationRoom extends Component {

  componentDidMount() {
    this.connect()
  }

  connect = () => {
    const socket = new WebSocket(`ws://localhost:8080/socketserver?token=${sessionStorage.getItem('token')}`);

    this.setState({
      socket: socket
    })

    socket.onopen = () => {
      this.createRoom()
      // this.getUsers()

    }

    socket.onmessage = (evt) => {
      // let currentMessages = this.state.response
      // currentMessages.unshift(evt.data)
      //
      // this.setState({
      //   response: currentMessages
      // })
        var messages = evt.data;
        console.log(messages)
    }

    socket.onclose = () => {
      console.log("connection close")
    }
  }

  createRoom = () => {
    const { socket } = this.state,
          { id } = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'CREATE_ESTIMATION_ROOM',
        issueID: id
      })

      socket.send(msg)
    }
  }

  getUsers = () => {
    const { socket } = this.state,
        { id } = this.props.ownProps.params

    if (socket) {
        let msg = JSON.stringify({
            action: 'GUEST',
            issueID: id
        })

        socket.send(msg)
    }
  }

  registerClient = () => {
    const { socket } = this.state,
          { id } = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'REGISTER_CLIENT',
        issueID: id,
      })

      socket.send(msg)
    }
  }

  sendEstimate = () => {
    const { socket } = this.state,
          { id } = this.props.ownProps.params

    if (socket) {

      let msg = JSON.stringify({
        action: 'ESTIMATION',
        issueID: id,
        estimate: this.state.estimate,
      })

      socket.send(msg)

      this.setState({
        estimate: ''
      })
    }
  }

  render() {
    const { classes } = this.props

    return (
      <div className={classes.root}>
        <Grid
          container
          spacing={40}
          justify={'center'}
          className={classes.container} >

          <Grid item xs={4}>

            <Issue
              issueID={this.props.ownProps.params.id} />

            <Users />

          </Grid>
          <Grid item xs={8}>

            <Estimation
              sendEstimate={this.sendEstimate}
              registerClient={this.registerClient} />

          </Grid>
        </Grid>
      </div>
    )
  }
}

const styles = {
  container: {
    padding: '2em 6em',
  },
  root: {
    width: '100%',
    overflow: 'hidden'
  },
  pos: {
    marginBottom: 12,
  },
}

const mapStateToProps = (state, ownProps) => {
  return {
    sprints: state.sprints,
    projects: state.projects,
    defaultPage: state.defaultPage,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    sprintsActions: bindActionCreators(sprintsActions, dispatch),
    projectsActions: bindActionCreators(projectsActions, dispatch),
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch)
  }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(EstimationRoom)
)
