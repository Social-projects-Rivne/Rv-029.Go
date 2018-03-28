import React, { Component } from 'react'
import openSocket from 'socket.io-client'

import { withStyles } from 'material-ui/styles'
import Button from 'material-ui/Button'
import Card, { CardActions, CardContent } from 'material-ui/Card'
import Typography from 'material-ui/Typography'
import TextField from 'material-ui/TextField'

class Socket extends Component {
  state = {
    sprintID: '',
    userID: '',
    issueID: '',
    estimate: '',
    socket: null,
    response: []
  }

  componentDidMount() {
    this.subscribe()
  }

  subscribe = () => {
    const socket = new WebSocket("ws://localhost:8080/socketserver?token=" + sessionStorage.getItem("token"));

    this.setState({
      socket: socket
    })

    socket.onopen = () => {
      console.log("Connection is open")
    }

    socket.onmessage = (evt) => {
      let currentMessages = this.state.response
      currentMessages.unshift(evt.data)

      this.setState({
        response: currentMessages
      })
    }

    socket.onclose = () => {
      console.log("connection close")
    }
  }

  // sendMessage = () => {
  //   const { socket } = this.state
  //
  //   let json = {action: 'CREATE_ESTIMATION_ROOM', sprintID: '12345', user: '123456'}
  //
  //   let jsonStr = JSON.stringify(json)
  //
  //   if (socket) {
  //     socket.send(jsonStr);
  //
  //     this.setState({
  //       message: ''
  //     })
  //   }
  // }

  createRoom = () => {
      const { socket } = this.state

      if (socket) {
        let msg = JSON.stringify({
          action: 'CREATE_ESTIMATION_ROOM',
          issueID: this.state.issueID, // only for debugging
        })
        console.log(msg);
        socket.send(msg)
      }
  }

  registerClient = () => {
    console.log('test');
    const { socket } = this.state
    console.log(socket);
    if (socket) {
      let msg = JSON.stringify({
        action: 'REGISTER_CLIENT',
        issueID: this.state.issueID, // only for debugging
      })
      console.log(msg);
      socket.send(msg)
    }
  }

  sendMessage = () => {
    const { socket } = this.state

    if (socket) {
      let msg = JSON.stringify({
        action: 'ESTIMATION',
        issueID: this.state.issueID,
        estimate: this.state.estimate
      })

      socket.send(msg)

      this.setState({
        estimate: ''
      })
    }
  }

  handleIssueChange = (e) => {
    this.setState({
      issueID: e.target.value
    })
  }

  handleEstimateChange = (e) => {
    this.setState({
      estimate: e.target.value
    })
  }

  render() {
    return (
      <Card className={this.props.classes.card}>
        <CardContent>
          <Typography
            type='subheading' >
            Socket Server
          </Typography>

          <TextField fullWidth
            className={this.props.classes.textField}
            label="Issue ID"
            value={this.state.issueID}
            onChange={this.handleIssueChange}
            margin="normal" />
          <TextField fullWidth
            className={this.props.classes.textField}
            label="Estimate"
            value={this.state.estimate}
            onChange={this.handleEstimateChange}
            margin="normal" />
        </CardContent>

        <CardActions>
          <Button
            raised
            color="secondary"
            onClick={this.createRoom}>
            Create Room
          </Button>

          <Button
            raised
            onClick={this.registerClient}>
            Register
          </Button>

          <Button
            raised
            onClick={this.sendMessage}>
            Send Estimation
          </Button>
        </CardActions>
        Issues:
          <p>910cbb11-25ec-11e8-b8d7-00224d6aa6bb</p>
          <p>910ce356-25ec-11e8-b8d8-00224d6aa6bb</p>
          <p>910d12af-25ec-11e8-b8d9-00224d6aa6bb</p>
          <p>910d2dd1-25ec-11e8-b8da-00224d6aa6bb</p>
          <p>910d539a-25ec-11e8-b8db-00224d6aa6bb</p>
          <p>910d6e83-25ec-11e8-b8dc-00224d6aa6bb</p>
          <p>910d8e0e-25ec-11e8-b8dd-00224d6aa6bb</p>
          <p>910dade5-25ec-11e8-b8de-00224d6aa6bb</p>
          <p>910ad1fb-25ec-11e8-b8cf-00224d6aa6bb</p>
          <p>910aeeea-25ec-11e8-b8d0-00224d6aa6bb</p>

        <CardContent>
            {
              this.state.response.map((item, i) => (
                <Typography
                  key={i}
                  type='body2'>
                  {item}
                </Typography>
              ))
            }
        </CardContent>
      </Card>
    )
  }
}

const styles = {
  card: {
    minWidth: 275,
    display: 'inline-block',
    minHeight: '100vh'
  }
}

export default withStyles(styles)(Socket);





// subscribe = () => {
//   const socket = openSocket.connect("http://localhost:8080", {
//     extraHeaders: {
//       Connection: "upgrade"
//     }
//   })
//
//   socket.on('open', () => {
//       ws.send("ping");
//       console.log("firs message sent")
//   })
//
//   socket.on('message', () => {
//       console.log(evt.data)
//       if(evt.data === "pong") {
//         setTimeout(function(){ws.send("ping");}, 2000);
//       }
//   })
//
//   socket.on('close', () => {
//       console.log("connection close")
//   })
// }
