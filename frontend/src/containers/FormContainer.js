import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import FormLogin from '../components/FormLogin'
import FormRegister from  '../components/FormRegister'
import FormRestorePassword from '../components/FormRestorePassword'
import * as formActions from '../actions/FormActions'

const FormContainer = ({ classes, form, formActions }) => {
  let formType = null

  if (form.type === 'login') {
    formType = <FormLogin form={form} action={formActions}/>
  } else if (form.type === 'register') {
    formType = <FormRegister form={form} action={formActions} />
  }
    else if (form.type === 'restorePassword') {
    formType = <FormRestorePassword form={form} action={formActions}/>
  }

  return (
    <Grid container
      className={classes.root}
      spacing={0}
      alignItems={'center'}
      justify={'center'} >
      {formType}
    </Grid>
  )
}

const mapStateToProps = (state) => {
  return {
    form: state.form
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    formActions: bindActionCreators(formActions, dispatch)
  }
}

const styles = {
  root: {
    height: '100vh',
    backgroundColor: '#2B2D42'
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(FormContainer)
)
