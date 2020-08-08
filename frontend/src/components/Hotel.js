import React from 'react'
import axios from 'axios';
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
        }
        // TODO: redirection to Login

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
        return (
            <div className="container">
                <p>{this.state.hotel.name}</p>
                <p>{this.state.hotel.price}</p>
                <p>{this.state.hotel.detail}</p>
                <p>{this.state.hotel.booking}/{this.state.hotel.max}</p>
                <button className="btn btn-primary" onClick={this.handleSubmit}>Booking</button>
            </div>
        )
    }
}

export default Hotel