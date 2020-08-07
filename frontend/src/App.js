import React from 'react';
import logo from './logo.svg';
import axios from 'axios';
import Hotel from './components/Hotel'
import Notfound from './components/NotFound'
import HotelList from "./HotelList"
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from  "react-router-dom";


function App() {
  return (
    <Router>
        <Switch>
            <Route exact path="/admin">
              <HotelList/>
            </Route>
            <Route exact path="/">
              <HotelList/>
            </Route>
            <Route path='/hotels/:name' render={(props) => {
                    return <Hotel name={props.match.params.name}/>

                }} />
            <Route>
              {/* TODO: 404 */}
              <Notfound/>
            </Route>
          </Switch>
    </Router>
  );
}
export default App;
