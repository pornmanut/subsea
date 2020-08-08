import React from 'react'
import axios from 'axios';
import './Login.css';

import {
    Grid,
    Button,
    Container,
    Typography,
    TextField,
    Backdrop,
    CircularProgress

} from '@material-ui/core';
import { Redirect } from 'react-router-dom';


class Login extends React.Component {

    async login() {
        const self = this;
        let token = '';

        const payload = {
            username: this.state.username,
            password: this.state.password
        }
        console.log(payload)

        self.setState({ isLoading: true })
        try {
            const response = await axios.post(window.global.api_location + '/login', payload)
            token = response.data.token
            self.setState({ isLoading: false, redirect: true, login: true })
        } catch (err) {
            self.setState({ isLoading: false, login: false })

        }
        return token

    }

    constructor(props) {
        super(props);
        this.state = { isLoading: false, username: "", redirect: null }
        this.login = this.login.bind(this);

        this.handleSubmit = this.handleSubmit.bind(this)
        this.handleChange = this.handleChange.bind(this)
    }

    handleSubmit(event) {
        event.preventDefault();

        this.login().then(token => {
            localStorage.setItem('token', token);
        })
    }
    handleChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
    }

    render() {
        const styles = {
            backdrop: {
                zIndex: 1,
                color: '#fff',
            },
            input: {
                width: "100%",
                padding: 10,
                marginTop: 20
            }
        }
        if (this.state.redirect) {
            return <Redirect to={"/"} />
        }
        return (
            <Container maxWidth="md" >
                <form className="form-signin" onChange={this.handleChange} >
                    <Typography variant="h5" component="h3">
                        Please Sign in
                    </Typography>
                    <TextField id="outlined-basic" style={styles.input} name="username" label="Username" />
                    <TextField id="outlined-basic" style={styles.input} name="password" type="password" label="Password" />
                    <Button variant="contained" color="primary" onClick={this.handleSubmit}>Sign in</Button>
                </form>
                <Backdrop style={styles.backdrop} open={this.state.isLoading}>
                    <CircularProgress color="inherit" />
                </Backdrop>
            </Container>

        )
    }
}

export default Login