// tmp

import React from 'react'
import Button from 'material-ui/Button'
import { browserHistory } from 'react-router'

const HomePage = () => {
  if (!sessionStorage.getItem('token')) {
    browserHistory.push("/authorization/login")
  }

  const logOut = () => {
    sessionStorage.removeItem('token')
    browserHistory.push("/authorization/login")
  }

  return (
    <Button
      onClick={logOut}
      color={'primary'}>
      Log Out
    </Button>
  )
}

export default HomePage