import React, { Component } from 'react'
import openSocket from 'socket.io-client'

import { withStyles } from 'material-ui/styles'
import Button from 'material-ui/Button'
import Card, { CardActions, CardContent } from 'material-ui/Card'
import Typography from 'material-ui/Typography'
import TextField from 'material-ui/TextField'

class Socket extends Component {
  state = {
    message: '',
    socket: null,
    response: []
  }

  componentDidMount() {
    this.subscribe()
  }

  subscribe = () => {
    const socket = new WebSocket("ws://localhost:8080/socketserver");

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
          sprintID: 'e80e4de0-1ae5-11e8-a63a-00224d6aa6bb', // only for debugging
        })

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
        sprintID: 'e80e4de0-1ae5-11e8-a63a-00224d6aa6bb', // only for debugging
        userID: '8f3d2d45-25ec-11e8-aec9-00224d6aa6bb'
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
        sprintID: 'e80e4de0-1ae5-11e8-a63a-00224d6aa6bb', // only for debugging
        userID: '8f3d2d45-25ec-11e8-aec9-00224d6aa6bb',
        message: this.state.message
      })

      socket.send(msg)

      this.setState({
        message: ''
      })
    }
  }

  handleChange = (e) => {
    this.setState({
      message: e.target.value
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

          <TextField
            className={this.props.classes.textField}
            label="Message"
            value={this.state.message}
            onChange={this.handleChange}
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
