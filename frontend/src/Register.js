import React from 'react'
import axios from 'axios';
import { Redirect } from 'react-router-dom';


import {
    Grid,
    Button,
    Container,
    Typography,
    TextField,
    Backdrop,
    CircularProgress

} from '@material-ui/core';

class Register extends React.Component {

    register(payload) {
        const self = this;
        console.log(payload)
        this.setState({ isLoading: true })

        axios.post(window.global.api_location + '/register', payload)
            .then(response => {
                console.log(response)
                this.setState({ isLoading: false, redirect: true })
            })
            .catch(err => {
                console.log(err)
                this.setState({ isLoading: false })
            })


    }

    handleChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
        console.log(this.state)
    };

    handleSubmit(event) {
        event.preventDefault()
        console.log(event)


        // TODO: verfiy field 


        const { email, firstname, lastname, username, password1, password2, birthdate } = this.state

        if (password1 !== password2) {
            // TODO: handle error password not match
            console.log("Password not match")
            return
        }
        // TODO: brithdate
        const payload = {
            email,
            username,
            password: password1,
            firstname,
            lastname,
            birthdate
        }

        this.register(payload)
        // eveything correct:
    }
    constructor(props) {
        super(props);

        this.state = {
            isLoading: false
        };

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.register = this.register.bind(this)

    }


    render() {
        const styles = {
            backdrop: {
                zIndex: 1,
                color: '#fff',
            },
        }

        if (this.state.redirect) {
            return <Redirect to="/login" />
        }
        return (
            <Container maxWidth="md">
                <Typography gutterBottom variant="h6" component="h4">
                    Register
                </Typography>
                <form onChange={this.handleChange} >
                    <Grid container>
                        <Grid item xs={6}>

                            <TextField id="outlined-basic" name="email" label="Email" />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField id="outlined-basic" name="username" label="Username" />

                        </Grid>
                    </Grid>
                    <Grid container>
                        <Grid item xs={6}>
                            <TextField id="outlined-basic" name="firstname" label="Firstname" />

                        </Grid>
                        <Grid item xs={6}>
                            <TextField id="outlined-basic" name="lastname" label="Lastname" />

                        </Grid>
                    </Grid>
                    <Grid container>
                        <Grid item xs={6}>
                            <TextField id="outlined-basic" name="password1" type="password" label="Password" />

                        </Grid>
                        <Grid item xs={6}>
                            <TextField id="outlined-basic" name="password2" type="password" label="Confirm Password" />

                        </Grid>
                    </Grid>
                    <Grid container>
                        <Grid item xs={6}>
                            <TextField
                                id="date"
                                name="birthdate"
                                label="Brithday"
                                type="date"
                                InputLabelProps={{
                                    shrink: true,
                                }} />
                        </Grid>
                        <Grid item xs={6}>
                            <Button variant="contained" color="primary" onClick={this.handleSubmit}>Submit</Button>

                        </Grid>
                    </Grid>

                </form>
                <Backdrop style={styles.backdrop} open={this.state.isLoading}>
                    <CircularProgress color="inherit" />
                </Backdrop>

            </Container >

        )
    }
}

export default Register