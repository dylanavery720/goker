import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import {getRank} from "./Actions/requests"

class App extends Component {
  constructor() {
    super()
    this.state = {
      rank: null
    }
  }

  componentDidMount() {
    //TODO: Use Picky for Dropdown to select 5 cards. Submit to get rank.
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
           Rank: {this.state.rank}
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;
