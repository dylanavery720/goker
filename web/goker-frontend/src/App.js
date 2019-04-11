import React, { Component } from 'react';
import './App.css';
import {getRank} from "./Actions/requests"
import Select from "./Select/Select"
import {deck, suits} from "./Constants/constants"


class App extends Component {
  constructor() {
    super()
    this.state = {
      rank: "",
      showRank: false,
      hand: "",
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

  handleClick(value, identifier) {
    this.setState({[identifier]: value })
  }

  async submitCards() {
    const {card1, card1Suit, card2, card2Suit, card3, card3Suit, card4, card4Suit, card5, card5Suit} = this.state
    let hand = ""
    hand = `${card1}${card1Suit} ${card2}${card2Suit} ${card3}${card3Suit} ${card4}${card4Suit} ${card5}${card5Suit}`
    const rank = await getRank(hand)
    this.setState({rank, hand, showRank: true})
  }

  render() {
    const {rank, hand} = this.state
    return (
      <div className="App">
        {!this.state.showRank && 
        <div>
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
        <button onClick={() => this.submitCards()}>SUBMIT HAND FOR REVIEW</button>
        </div>
        }
         <footer className="App-footer">
          <p>
           Rank: {rank}
          </p>
          <p>
           Hand: {hand}
          </p>
        </footer>
      </div>
    );
  }
}

export default App;
