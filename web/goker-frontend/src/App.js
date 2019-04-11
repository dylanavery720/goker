import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import {getRank} from "./Actions/requests"
import Select from "./Select/Select"
import {deck, suits} from "./Constants/constants"


class App extends Component {
  constructor() {
    super()
    this.state = {
      rank: "",
      card1: "",
      card1Suit: "",
      card2: "",
      card2Suit: "",
      card3: "",
      card3Suit: "",
      card4: "",
      card4Suit: "",
      card5: "",
      card5Suit: ""
    }
  }

  componentDidMount() {
    //TODO: Use Picky for Dropdown to select 5 cards. Submit to get rank.
  }

  handleClick(value, identifier) {
    this.setState({[identifier]: value })
  }

  submitCards() {
    const {card1, card1Suit, card2, card2Suit, card3, card3Suit, card4, card4Suit, card5, card5Suit} = this.state
    let hand = ""
    hand = `${card1}${card1Suit} ${card2}${card2Suit} ${card3}${card3Suit} ${card4}${card4Suit} ${card5}${card5Suit}`
    console.log(hand,'ranky')
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
        <p>Card 1</p>
        <Select handleClick={(value) => this.handleClick(value, "card1")} options={Object.keys(deck)}/>
        <Select handleClick={(value) => this.handleClick(value, "card1Suit")} options={Object.keys(suits)}/>
        <p>Card 2</p>
        <Select handleClick={(value) => this.handleClick(value, "card2")} options={Object.keys(deck)}/>
        <Select handleClick={(value) => this.handleClick(value, "card2Suit")} options={Object.keys(suits)}/>
        <p>Card 3</p>
        <Select handleClick={(value) => this.handleClick(value, "card3")} options={Object.keys(deck)}/>
        <Select handleClick={(value) => this.handleClick(value, "card3Suit")} options={Object.keys(suits)}/>
        <p>Card 4</p>
        <Select handleClick={(value) => this.handleClick(value, "card4")} options={Object.keys(deck)}/>
        <Select handleClick={(value) => this.handleClick(value, "card4Suit")} options={Object.keys(suits)}/>
        <p>Card 5</p>
        <Select handleClick={(value) => this.handleClick(value, "card5")} options={Object.keys(deck)}/>
        <Select handleClick={(value) => this.handleClick(value, "card5Suit")} options={Object.keys(suits)}/>
        <button onClick={() => this.submitCards()}></button>
      </div>
    );
  }
}

export default App;
