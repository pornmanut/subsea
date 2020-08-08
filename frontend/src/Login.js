import React from 'react'
import axios from 'axios';
import './Login.css';

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

        self.setState({ isLogining: true })
        try {
            const response = await axios.post(window.global.api_location + '/login', payload)
            token = response.data.token
            self.setState({ isLogining: false, redirect: true, login: true })
        } catch (err) {
            self.setState({ isLogining: false, redirect: true, login: false })

        }

        return token

    }

    constructor(props) {
        super(props);
        this.state = { isLogining: false, username: "", redirect: null }
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
        console.log(this.state.isLogining)
        if (this.state.isLogining === true) {
            return (
                <div>
                    Login
                </div>
            )
        }
        if (this.state.redirect) {
            if (this.state.login) {
                return <Redirect to={"/"} />
            }
        }
        return (
            <div className="container">
                <form className="form-signin" onSubmit={this.handleSubmit} >
                    <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
                    <label className="sr-only">Username</label>
                    <input
                        type="username"
                        name="username"
                        className="form-control"
                        placeholder="Username"
                        onChange={this.handleChange}
                    />
                    <label htmlFor="inputPassword" className="sr-only">Password</label>
                    <input
                        type="password"
                        name="password"
                        className="form-control"
                        placeholder="Password"
                        onChange={this.handleChange}
                    />
                    <button className="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
                </form>
            </div>
        )
    }
}

export default Login