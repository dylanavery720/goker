import React, { Component } from 'react';
import Picky from 'react-picky';
import 'react-picky/dist/picky.css';
import {deck} from "../Constants/constants"

class Select extends Component {
  constructor(props) {
    super(props);
    this.state = {
      value: null,
    };
  }

  selectOption(value) {
    // const retailerOffers = [];
    // for (let i = 0; i < value.length; i++) {
    //   retailerOffers.push(
    //     this.props.retailers.filter(retailer => value[i] === retailer.name),
    //   );
    // }

    this.props.handleClick(value);
    this.setState({ value: value });
  }

  render() {
    return (
      <Picky
        options={this.props.options}
        value={this.state.value}
        multiple={false}
        includeFilter={true}
        onChange={values => this.selectOption(values)}
        dropdownHeight={600}
      />
    );
  }
}

export default Select;