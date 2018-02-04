import React from 'react'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'

const FormContainer = ({ classes, form, formActions, children }) => {
  return (
    <Grid container
      className={classes.root}
      spacing={0}
      alignItems={'center'}
      justify={'center'} >
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

export default withStyles(styles)(FormContainer)
