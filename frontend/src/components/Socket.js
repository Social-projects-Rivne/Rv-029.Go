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

  sendMessage = () => {
    const { socket } = this.state

    if (socket) {
      socket.send(this.state.message);

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
            onClick={this.subscribe}>
            Connect
          </Button>

          <Button
            raised
            onClick={this.sendMessage}>
            Send
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

