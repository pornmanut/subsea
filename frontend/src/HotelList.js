import React from 'react';
import axios from 'axios';
import Filter from "./components/Filter"
import HotelCard from "./components/HotelCard"
import Container from '@material-ui/core/Container';

import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography'
import {
    Link
} from "react-router-dom"

class HotelList extends React.Component {
    //TODO: searchbar

    constructor(props) {
        super(props);
        this.readHotels();
        this.state = { hotels: [], name: '' };

        this.readHotels = this.readHotels.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    readHotels(name) {
        let url = window.global.api_location + '/hotels'
        console.log(name)
        if (name != undefined && name != '') {
            url = url + "?name=" + name
        }
        const self = this;
        axios.get(url).then(function (response) {
            console.log(response.data);
            self.setState({ hotels: response.data });
        }).catch(function (error) {
            console.log(error);
        });
    }
    getHotels() {
        let table = []

        for (let i = 0; i < this.state.hotels.length; i++) {
            table.push(
                <HotelCard hotel={this.state.hotels[i]} />
            );
        }

        return table
    }

    handleChange(event) {
        this.setState({ name: event.target.value })
    }

    handleSubmit(event) {
        console.log(this.state.name)
        this.readHotels(this.state.name)
        event.preventDefault();
    }


    render() {
        return (
            <Container maxWidth="md">
                <form onSubmit={this.handleSubmit}>
                    <label>
                        <Filter value={this.state.name} handleChange={this.handleChange} />
                    </label>
                    <input type="submit" value="Submit" />
                </form>
                {this.getHotels()}
            </Container>
        )
    }
}

export default HotelList