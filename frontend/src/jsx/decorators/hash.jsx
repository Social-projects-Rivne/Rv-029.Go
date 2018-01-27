import React, { Component } from 'react';
import CryptoJs from 'crypto-js';

export default (OriginalComponent) => class InjectHash extends Component {
  MD5hash = (string) => {
    return(CryptoJs.MD5(string).toString());
  }

  render() {
    return (
      <OriginalComponent {...this.props} MD5hash={this.MD5hash} />
    )
  }
}
