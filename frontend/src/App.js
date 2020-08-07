import React from 'react';
import logo from './logo.svg';

import HotelList from "./HotelList"
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from  "react-router-dom";


function App() {
  return (
    <Router>
      <div className="App">
        <Switch>
            <Route path="/admin">
            </Route>
            <Route path="/">
              <HotelList/>
            </Route>
            <Route>
            </Route>
          </Switch>
      </div>
    </Router>
  );
}

export default App;
