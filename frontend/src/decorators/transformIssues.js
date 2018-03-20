import React, { Component } from 'react'
import { EMPTY_ID } from '../constants/global'

export default (OriginalComponent) => class InjectTransformIssues extends Component {

  buildTree = (issues, result = []) => {
    issues.forEach((item) => {
      if (item.Parent === EMPTY_ID) {
        result.push(item)
      }
    })

    result.forEach((item) => {
      item.Children = this.setChildren(issues, item.UUID)
    })

    this.setNesting(result)

    return this.flatTransform(result)
  }

  flatTransform = (hierarchy, result = []) => {
    for (let i = 0; i < hierarchy.length; i++) {

      // improve nesting output
      let nesting = hierarchy[i].Nesting.split('>')

      nesting.splice(-1, 1)

      if (nesting.length) {
        nesting.push('')
      }

      hierarchy[i].Nesting = nesting.join('>')


      result.push(hierarchy[i])

      if (hierarchy[i].Children.length) {
        this.flatTransform(hierarchy[i].Children, result)
      }
    }

    return result
  }

  setNesting = (arr, prefix = null) => {
    if (arr.length > 0) {
      for (let i = 0; i < arr.length; i++) {
        if (!prefix) {
          arr[i].Nesting = arr[i].Name
        } else {
          arr[i].Nesting = prefix + ' > ' + arr[i].Name
        }

        if (arr[i].Children.length > 0) {
          this.setNesting(arr[i].Children, arr[i].Nesting)
        }
      }
    }
  }

  setChildren = (arr, parent) => {
    let out = []

    for (let i in arr) {
      if (arr[i].Parent === parent) {
        let children = this.setChildren(arr, arr[i].UUID)

        children.length ?
          arr[i].Children = children :
          arr[i].Children = []

        out.push(arr[i])
      }
    }

    return out
  }

  render() {
    return (
      <OriginalComponent {...this.props} transformIssues={this.buildTree} />
    )
  }
}
