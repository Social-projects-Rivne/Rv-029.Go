import React, {Component} from 'react';
import axios from 'axios';

class Form extends Component {
  constructor(props) {
    super(props);

    this.state = {
      buttonHasClicked: 'false'
    }
  }

  getMessage = (e) => {
    this.setState({ buttonHasClicked: 'true' });
    console.log('Form state has been changed');

    axios.get('/data')
      .then(function (response) {
        var p = document.createElement('p');
        p.textContent = response.data;
        document.body.appendChild(p);
      })
      .catch(function (error) {
        console.log(error);
      });

    e.preventDefault();
  }

  render() {
    return (
      <form className="form">
        <button onClick={this.getMessage}>
          Click
        </button>
      </form>
    )
  }
}

export default Form;
