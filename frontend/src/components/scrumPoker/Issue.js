import React, { Component } from 'react'
import Typography from 'material-ui/Typography'
import Card, { CardContent } from 'material-ui/Card'
import {withStyles} from 'material-ui/styles'

class Issue extends Component {
  render() {
    const { classes } = this.props

    return (
      <Card className={classes.cardMargin}>
        <CardContent>
          <Typography type='headline' component="h2">
            Estimation
          </Typography>
          <Typography color="textSecondary">
            { this.props.name }
          </Typography>

          {(this.props.estimationResult) ? (

            <Typography type='headline' component="h2">
              Result: 10
            </Typography>

          ) : (null)}

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
