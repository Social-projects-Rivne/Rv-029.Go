import React, {Component} from 'react'
import {connect} from 'react-redux'
import axios from "axios";
import {bindActionCreators} from 'redux'
import * as scrumPokerActions from "../actions/ScrumPokerAction"
import * as defaultPageActions from "../actions/DefaultPageActions"
import { API_URL } from '../constants/global'
import messages from "../services/messages"
import SnackBar from '../components/scrumPoker/SnackBar'
import Users from '../components/scrumPoker/Users'
import Issue from '../components/scrumPoker/Issue'
import Estimation from '../components/scrumPoker/Estimation'
import {withStyles} from 'material-ui/styles'
import Grid from 'material-ui/Grid'

class EstimationRoom extends Component {

  state = {
    users: null,
    responseData: null,
    socket: null,
    issueName: '',
    issueDesc: '',
    issueStatus: ''
  }

  componentWillMount() {
    this.getIssue()
  }

  componentDidMount() {
    this.connect()
  }

  componentWillUnmount() {
    this.state.socket.close();
  }

  connect = () => {
    const socket = new WebSocket(`ws://localhost:8080/socketserver?token=${sessionStorage.getItem('token')}`)

    this.setState({
      socket: socket
    })

    socket.onopen = () => {
      this.createRoom()
      this.getUsers()
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
          {action, message, status, data} = res

    // set state for notification message
    this.setState({responseData: res})

    console.log(`actions: ${action}, message: ${message}`)

    switch (action) {
      case 'CREATE_ESTIMATION_ROOM':

        break
      case 'REGISTER_CLIENT':
        if (status) { this.props.scrumPokerActions.setStep(2) }
        break
      case 'ESTIMATION':
        if (status) { this.props.scrumPokerActions.setStep(3) }
        break
      case 'ESTIMATION_RESULTS':

        if (status) {
          this.props.scrumPokerActions.setEstResult(data.estimate)
          this.updateIssue(data.estimate)
        }

        break
      case 'GUEST':
        let roomUsers = [];
        for (let roomUser in data) {
            roomUsers.push(data[roomUser])
        }
        this.setState({users: roomUsers})
        break
      case 'NEW_USER_IN_ROOM':
          let newUsers = this.state.users.slice()

          if (!newUsers.includes(data)) {
              newUsers.push(data)
          }

          this.setState({users: newUsers})
        break
      case 'USER_DISCONNECT_FROM_ROOM':
          let users = this.state.users.slice()

          users.forEach(function(element, key) {
              if(element.UUID === data.UUID){
                  users.splice(key, 1)
              }

          });
          console.log(users)
          this.setState({users: users})
          break

    }
  }

  getIssue = () => {
    const { id } = this.props.ownProps.params

    axios.get(`${API_URL}project/board/issue/show/${ id }`)
      .then((res) => {
        this.setState({
          issueName: res.data.Name,
          issueDesc: res.data.Description,
          issueStatus: res.data.Status
        })
      })
      .catch((error) => {
        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occurred")
        }
      })
  }

  updateIssue = (result) => {
    const { id } = this.props.ownProps.params
    const { setErrorMessage, setNotificationMessage } = this.props.defaultPageActions

    axios.put(API_URL + `project/board/issue/update/${ id }`, {
      name: this.state.issueName,
      description: this.state.issueDesc,
      status: this.state.issueStatus,
      estimate: result
    })
      .then((res) => {
        setNotificationMessage(res.data.Message)
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

  createRoom = () => {
    const {socket} = this.state,
          {id} = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'CREATE_ESTIMATION_ROOM',
        issueID: id
      })

      socket.send(msg)
    }
  }

  getUsers = () => {
    const {socket} = this.state,
      {id} = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'GUEST',
        issueID: id
      })

      socket.send(msg)
    }
  }

  registerClient = () => {
    const {socket} = this.state,
      {id} = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'REGISTER_CLIENT',
        issueID: id,
      })

      socket.send(msg)
    }
  }

  sendEstimate = (est) => {
    const {socket} = this.state,
      {id} = this.props.ownProps.params

    if (socket) {
      let msg = JSON.stringify({
        action: 'ESTIMATION',
        issueID: id,
        estimate: est.toString()
      })

      socket.send(msg)
    }
  }

  render() {
    const {responseData} = this.state,
      {scrumPoker, scrumPokerActions, classes} = this.props

    return (
      <div className={classes.root}>
        <Grid
          container
          spacing={40}
          justify={'center'}
          className={classes.container}>

          <Grid item xs={4}>

            <Issue
              name={this.state.issueName}
              estimationResult={this.props.scrumPoker.estimationResult}/>

            <Users
              users={this.state.users}/>

          </Grid>
          <Grid item xs={8}>

            <Estimation
              actions={scrumPokerActions}
              activeStep={scrumPoker.activeStep}
              sendEstimate={this.sendEstimate}
              registerClient={this.registerClient}/>

          </Grid>
        </Grid>

        <SnackBar
          options={responseData}/>

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
    scrumPokerActions: bindActionCreators(scrumPokerActions, dispatch),
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch)
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(EstimationRoom)
)
