import React from 'react'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import TopBar from '../components/TopBar';

const DefaultPage = ({ classes, children }) => {
  return (
    <Grid
      className={classes.root} >
      <TopBar />
      { children }
    </Grid>
  )
}

const styles = {
  root: {
    minHeight: '100vh',
    backgroundColor: '#2B2D42',
  }
}

export default withStyles(styles)(DefaultPage)
