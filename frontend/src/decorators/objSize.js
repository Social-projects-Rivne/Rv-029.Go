import React, { Component } from 'react'

export default (OriginalComponent) => class InjectObjSize extends Component {
    size = (obj) => {
        var size = 0, key;
        for (key in obj) {
            if (obj.hasOwnProperty(key)) size++;
        }
    
        return size;
    };

  render() {
    return (
      <OriginalComponent {...this.props} objSize={this.size} />
    )
  }
}
