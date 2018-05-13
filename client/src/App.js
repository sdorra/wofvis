import React, { Component } from 'react';
import WebOfTrust from './WebOfTrust';

import "bulma/css/bulma.css";

class App extends Component {
  render() {
    return (
        <section className="section">
            <div className="container">
                <h1 className="title">Web Of Trust</h1>
                <p className="subtitle">Show the OpenPGP Web of Trust</p>
                <WebOfTrust />
            </div>
        </section>
    );
  }
}

export default App;
