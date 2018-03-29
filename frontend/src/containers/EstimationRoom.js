import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as scrumPokerActions from "../actions/ScrumPokerAction";
import SnackBar from '../components/scrumPoker/SnackBar'
import Users from '../components/scrumPoker/Users'
import Issue from '../components/scrumPoker/Issue'
import Estimation from '../components/scrumPoker/Estimation'
import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';

class EstimationRoom extends Component {

  state = {
    responseData: null,
  }

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
    }

    socket.onmessage = (evt) => {
      this.actionHandler(evt.data)
    }

    socket.onclose = () => {
      console.log("connection close")
    }
  }

  actionHandler = (jsonResponse) => {
    const res = JSON.parse(jsonResponse),
          { action, message, status } = res

    // set state for notification message
    this.setState({ responseData: res })

    switch (action) {
      case 'CREATE_ESTIMATION_ROOM':

        break
      case 'REGISTER_CLIENT':
        if (status) {
          this.props.scrumPokerActions.setStep(2)
        }
        break
      case 'ESTIMATION':
        if (status) {
          this.props.scrumPokerActions.setStep(3)
        }
        break
      case 'ESTIMATION_RESULT':
        break
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

  sendEstimate = (est) => {
    const { socket } = this.state,
          { id } = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'ESTIMATION',
        issueID: id,
        estimate: est
      })

      socket.send(msg)

    }
  }

  render() {
    const { responseData } = this.state,
          { scrumPoker, scrumPokerActions, classes } = this.props

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
              actions={scrumPokerActions}
              activeStep={scrumPoker.activeStep}
              sendEstimate={this.sendEstimate}
              registerClient={this.registerClient} />

          </Grid>
        </Grid>

        <SnackBar
          options={ responseData } />

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
    scrumPoker: state.scrumPoker,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    scrumPokerActions: bindActionCreators(scrumPokerActions, dispatch)
  }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(EstimationRoom)
)
