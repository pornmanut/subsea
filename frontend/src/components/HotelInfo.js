import React from 'react'
import axios from 'axios';
import { Redirect } from 'react-router-dom';
import {
    Grid,
    Card,
    CardActionArea,
    CardActions,
    CardContent,
    CardMedia,
    Button,
    Typography,
    Container,
    Paper,
    Divider

} from '@material-ui/core';
import NotFound from './NotFound';
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
        const styles = {
            indent: {
                textIndent: 30,
                padding: 20
            },
            image: {
                width: "100%",
                height: "auto"
            },
            textRight: {
                textAlign: "right",
                paddingRight: 20
            },
            booking: {
                marginTop: 100
            }

        }
        const hotel = this.state.hotel
        if (!this.state.found) {
            return <NotFound />
        }
        if (this.state.toLogin) {
            return <Redirect to="/login" />
        }
        return (
            <Container maxWidth="md">
                <Paper>
                    <Grid container spacing={2}>
                        <Grid item xs={6}>
                            <img src="https://via.placeholder.com/500x500" style={styles.image} />
                        </Grid>
                        <Grid item md xs={6}>
                            <Typography gutterBottom variant="h4" component="h4">
                                {hotel.name}
                            </Typography>
                            <Divider />
                            <Typography gutterBottom variant="h6" component="p" style={styles.indent}>
                                {hotel.detail}
                            </Typography>
                            <Typography gutterBottom component="p" style={styles.textRight}>
                                {hotel.price} $
                            </Typography>
                            <Grid container style={styles.booking}>
                                <Grid item xs={6}>
                                    <Button variant="contained" color="primary" onClick={this.handleSubmit} >Booking Now</Button>
                                </Grid>
                                <Grid item xs={6}>
                                    <Typography gutterBottom component="p" style={styles.textRight} color="primary">
                                        Available {this.state.hotel.booking} from {this.state.hotel.max} Rooms
                                    </Typography>
                                </Grid>
                            </Grid>
                        </Grid>
                    </Grid>
                </Paper>
            </Container >

        )
    }
}

export default Hotel