import React, { Component } from 'react'
import openSocket from 'socket.io-client'
import Grid from 'material-ui/Grid';
import { withStyles } from 'material-ui/styles'
import Button from 'material-ui/Button'
import Card, { CardActions, CardContent } from 'material-ui/Card'
import Typography from 'material-ui/Typography'
import TextField from 'material-ui/TextField'

class CreateRoom extends Component {
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

  createRoom = () => {
      const { socket } = this.state

      if (socket) {
        let msg = JSON.stringify({
          action: 'CREATE_ESTIMATION_ROOM',
          sprintID: this.props.sprintID,
        })

        socket.send(msg)
      }
  }

  render() {
    return (
        <Grid item xs={12} container>
          <CardActions>
            <Button
              raised
              color="secondary"
              onClick={this.createRoom}>
              Create Room
            </Button>
          </CardActions>
        </Grid>
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

export default withStyles(styles)(CreateRoom);






