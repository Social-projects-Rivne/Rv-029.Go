import React, { Component } from 'react'
import CryptoJs from 'crypto-js'

export default (OriginalComponent) => class InjectHash extends Component {
  MD5Encode = (string) => {
    return(CryptoJs.MD5(string).toString())
  }

  render() {
    return (
      <OriginalComponent {...this.props} MD5Encode={this.MD5Encode} />
    )
  }
}
