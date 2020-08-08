import React from 'react'
import axios from 'axios';
import Button from '@material-ui/core/Button';
import { Redirect } from 'react-router-dom';

// booking: 1
// detail: "great view"
// height: 30.3
// id: "5f2ce2ea0365c2ae52c7d39e"
// max: 1
// name: "abc"
// price: 300

// TODO: booking
class Hotel extends React.Component {
    readData(name) {
        const self = this;
        axios.get(window.global.api_location + '/hotels/' + name).then(function (response) {
            console.log(response.data);

            self.setState({ hotel: response.data, found: true });
        }).catch(function (error) {
            console.log(error);
            self.setState({ found: false })
        });
    }


    booking(name, token) {
        // require token this request
        const self = this;
        axios.get(window.global.api_location + '/hotels/booking/' + name, {
            headers: {
                'Authorization': token
            }
        }).then(function (response) {
            //reload page
            window.location.reload(false);
            console.log(response.data);
        }).catch(function (error) {
            console.log(error);
        });
    }


    handleSubmit() {
        // get token from local

        const token = localStorage.getItem("token")

        if (token) {
            this.booking(this.props.name, token)
            return
        }
        // TODO: redirection to Login
        this.setState({ toLogin: true })
    }
    constructor(props) {
        super(props);
        this.state = { hotel: "" };
        this.readData(props.name)
        this.readData = this.readData.bind(this);
        this.booking = this.booking.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    render() {
        if (!this.state.found) {
            return <div>loading</div>
        }
        if (this.state.toLogin) {
            return <Redirect to="/login" />
        }
        return (
            <div className="container">
                <p>{this.state.hotel.name}</p>
                <p>{this.state.hotel.price}</p>
                <p>{this.state.hotel.detail}</p>
                <p>{this.state.hotel.booking}/{this.state.hotel.max}</p>
                <Button variant="contained" color="primary" onClick={this.handleSubmit} >Booking</Button>
                <button className="btn btn-primary" onClick={this.handleSubmit}>Booking</button>
            </div>
        )
    }
}

export default Hotel