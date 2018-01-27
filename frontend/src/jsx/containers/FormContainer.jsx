import React, { Component } from 'react';
import Grid from 'material-ui/Grid';
import { withStyles } from 'material-ui/styles';

function FormContainer(props) {
  const { classes } = props;

  return (
    <Grid container 
      className={classes.root}
      spacing={8}
      alignItems={'center'}
      justify={'center'} >
        {props.children}
    </Grid>
  )
}

const styles = {
  root: {
    height: '100vh',
    backgroundColor: '#f9f9f9'
  }
}

export default withStyles(styles)(FormContainer);
