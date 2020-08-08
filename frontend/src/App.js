import React from 'react';
import Hotel from './components/HotelInfo'
import Notfound from './components/NotFound'
import HotelList from "./HotelList"
import Register from "./Register"
import MyBookingList from "./MyBookingList"
import Login from "./Login"
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";


function App() {
  return (
    <div className="container">
      <Router>
        <Switch>
          <Route exact path="/hotels">
            <HotelList />
          </Route>
          <Route exact path="/">
            <HotelList />
          </Route>
          <Route path='/hotels/:name' render={(props) => {
            return <Hotel name={props.match.params.name} />
          }} />
          <Route exact path='/login'>
            <Login />
          </Route>
          <Route exact path='/register'>
            <Register />
          </Route>
          <Route exact path='/mybookings'>
            <MyBookingList />
          </Route>
          <Route>
            <Notfound />
          </Route>

        </Switch>
      </Router>
    </div >

  );
}
export default App;
