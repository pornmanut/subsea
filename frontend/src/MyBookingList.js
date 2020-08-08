import React from 'react';
import axios from 'axios';
import Filter from "./components/Filter"

import {
    Link,
    Redirect
} from "react-router-dom"

class MyBookingList extends React.Component {
    //TODO: searchbar

    constructor(props) {
        super(props);

        const token = localStorage.getItem("token")
        this.state = {
            token: token,
            hotels: []
        };

        this.getBookingList = this.getBookingList.bind(this)
    }
    getBookingList(token) {
        // require token this request
        const self = this;
        axios.get(window.global.api_location + '/users/booking', {
            headers: {
                'Authorization': token
            }
        }).then(function (response) {
            self.setState({ hotels: response.data })
        }).catch(function (error) {
            console.log(error);
        });
    }


    componentDidMount() {
        this.getBookingList(this.state.token)
    }

    render() {
        // already login
        if (this.state.token) {
            console.log(this.state.hotels)
            return (
                <div>
                </div>
            )
        }


        return <Redirect to={"/"} />

    }
}

export default MyBookingList