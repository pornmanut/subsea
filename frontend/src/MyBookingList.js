import React from 'react';
import axios from 'axios';
import HotelCard from "./components/HotelCard"

import {
    Grid,
    Button,
    Container,
    Typography,
    TextField,
    Backdrop,
    CircularProgress,
    Card,
    CardActionArea,
    CardActions,
    CardContent,
    CardMedia,

} from '@material-ui/core';
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
        this.getHotels = this.getHotels.bind(this)
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
            self.setState({ hotels: [] })
            console.log(error);
        });
    }


    componentDidMount() {
        this.getBookingList(this.state.token)
    }
    getHotels() {
        let table = []

        for (let i = 0; i < this.state.hotels.length; i++) {
            table.push(
                <HotelCard key={i} hotel={this.state.hotels[i]} />
            );
        }

        return table
    }
    render() {
        // already login
        if (this.state.token) {

            if (this.state.hotels.length === 0) {
                return (
                    <Container maxWidth="md">
                        <Typography>
                            No bookings
                        </Typography>
                    </Container>
                )
            }
            return (
                <Container maxWidth="md">
                    <Typography variant="h6" component="h4">
                        My Bookings
                    </Typography>
                    {this.getHotels()}
                </Container>
            )
        }


        return <Redirect to={"/login"} />

    }
}

export default MyBookingList