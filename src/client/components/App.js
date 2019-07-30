import React, { Component } from 'react';
import logo from '../images/logo.svg';
import '../style/App.css';
import { CarClient } from "../api/v1/motor_pb_service";
import { render } from 'react-dom';

export default class App extends Component {

    constructor(props) {
        super(props);

        let client = new CarClient("http://" + window.location.hostname + ":8551");

        this.state = {
            client: client
        }
    }


    onComponentDidMount() {

    }

    render(){
        return (
            <div className="App">
                <header className="App-header">
                    Welcome to the car!
                </header>
                <svg width="600" height="600" className="arena" >
                    <rect width="25" height="40" x="300" y="300" transform = "rotate(30, 300, 300)" className="car">
                        <rect width="3" height="5" className="stripe"></rect>
                    </rect>
                </svg>
            </div>
        );
    }
}