import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Redirect } from 'react-router-dom';

import {
    Grid,
    Card,
    CardActionArea,
    CardActions,
    CardContent,
    CardMedia,
    Button,
    Typography

} from '@material-ui/core';


class HotelCard extends React.Component {

    constructor(props) {
        super(props)
        this.state = { redirect: false }

        this.handleClick = this.handleClick.bind(this)
    }
    handleClick() {
        this.setState({ redirect: true })
    }
    render() {
        const styles = {
            root: {
                margin: 15,
            },

            media: {
                height: 250,
            }
        }
        const hotel = this.props.hotel
        if (this.state.redirect) {
            return <Redirect to={"/hotels/" + hotel.name} />
        }
        return (

            //TODO: hotel show
            // image optional
            // location
            // detail list bullet
            // telephone
            //price
            <Card style={styles.root} >
                <CardActionArea onClick={this.handleClick}>
                    <CardMedia
                        image="https://via.placeholder.com/500x500"
                        title="Contemplative Reptile"
                        style={styles.media}
                    />
                    <CardContent>

                        <Grid container>

                            <Grid item xs={10}>
                                <Typography gutterBottom variant="h5" component="h2">
                                    {hotel.name}
                                </Typography>
                                <Typography variant="body2" color="textSecondary" component="p">
                                    {hotel.detail}
                                </Typography>
                            </Grid>
                            <Grid item xs={2}>
                                <Typography variant="h5" align="right" color="textPrimary" component="p">
                                    {hotel.price} $
                                 </Typography>
                            </Grid>
                        </Grid>


                    </CardContent>
                </CardActionArea>
            </Card >
        );
    }
}

export default HotelCard