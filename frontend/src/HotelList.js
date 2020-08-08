import React from 'react';
import axios from 'axios';
import Filter from "./components/Filter"
import HotelCard from "./components/HotelCard"


import {
    Grid,
    Card,
    CardActionArea,
    CardActions,
    CardContent,
    CardMedia,
    Button,
    Container,
    Typography,
    TextField

} from '@material-ui/core';



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
    readHotels(name, detail, lt, gt) {
        let url = window.global.api_location + '/hotels?'
        if (name != undefined && name != '') {
            url = url + "&name=" + name
        }

        // DETAIL SEARCH WITH REGEXP. 
        // WILL MAKE MONGODB KABBOOOM
        if (detail != undefined && detail != '') {
            url = url + "&detail=" + detail
        }

        if (lt != undefined && lt != '') {
            url = url + "&lt=" + lt
        }

        if (gt != undefined && gt != '') {
            url = url + "&gt=" + gt
        }
        const self = this;
        axios.get(url).then(function (response) {
            console.log(response.data);
            let res = []
            if (response.data) {
                res = response.data
            }
            self.setState({ hotels: res });
        }).catch(function (error) {
            self.setState({ hotels: [] });

            console.log(error);
        });
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

    handleChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;
        console.log(this.state)
        this.setState({
            [name]: value
        })
    }

    handleSubmit(event) {
        const data = this.state
        console.log(this.state.name)
        this.readHotels(data.name, data.detail, data.lt, data.gt)
        event.preventDefault();
    }


    render() {
        const styles = {
            input: {
                width: "80%",
                padding: 5
            },
        }
        return (

            <Container maxWidth="md" >
                <Typography gutterBottom variant="h6" component="h4">
                    Search Options
                </Typography>
                <Grid container>
                    <Grid item xs={10}>
                        <Grid container>
                            <Grid item xs={6}>
                                <TextField style={styles.input} type="search" id="outlined-basic" name="name" label="Name" onChange={this.handleChange} />

                            </Grid>
                            <Grid item xs={6}>
                                <TextField style={styles.input} type="search" id="outlined-basic" name="detail" label="Detail" onChange={this.handleChange} />

                            </Grid>
                        </Grid>

                        <Grid container>
                            <Grid item xs={6}>
                                <TextField style={styles.input} type="search" id="outlined-basic" name="lt" label="Less than" onChange={this.handleChange} />

                            </Grid>
                            <Grid item xs={6}>
                                <TextField style={styles.input} type="search" id="outlined-basic" name="gt" label="Greater than" onChange={this.handleChange} />

                            </Grid>
                        </Grid>
                    </Grid>
                    <Grid item xs={2}>
                        <Button size="large" variant="contained" color="primary" onClick={this.handleSubmit}>Search</Button>
                    </Grid>
                </Grid>
                {this.getHotels()}
            </Container >
        )
    }
}

export default HotelList