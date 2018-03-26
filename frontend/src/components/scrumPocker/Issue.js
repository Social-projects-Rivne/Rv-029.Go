import React, { Component } from 'react'
import axios from "axios";
import { API_URL } from '../../constants/global'
import messages from "../../services/messages"
import Typography from 'material-ui/Typography'
import Card, { CardContent } from 'material-ui/Card'
import {withStyles} from 'material-ui/styles'

class Issue extends Component {

  state = {
    issueName: ''
  }

  componentWillMount() {
    this.getIssue()
  }

  getIssue = () => {
    const { issueID } = this.props

    axios.get(`${API_URL}project/board/issue/show/${issueID}`)
      .then((res) => {
        this.setState({
          issueName: res.data.Name
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

  render() {
    const { classes } = this.props

    return (
      <Card className={classes.cardMargin}>
        <CardContent>
          <Typography type='headline' component="h2">
            Estimation
          </Typography>
          <Typography color="textSecondary">
            { this.state.issueName }
          </Typography>
        </CardContent>
      </Card>
    )
  }
}

const styles = {
  cardMargin: {
    marginBottom: '1em'
  },
}

export default withStyles(styles)(Issue)
