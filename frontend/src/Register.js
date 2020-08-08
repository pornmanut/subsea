import React from 'react'
import axios from 'axios';
import DatePicker from "react-datepicker";
import { Redirect } from 'react-router-dom';

import "react-datepicker/dist/react-datepicker.css";
class Register extends React.Component {

    register(payload) {
        const self = this;
        console.log(payload)

        axios.post(window.global.api_location + '/register', payload)
            .then(response => {
                console.log(response)
            })
            .catch(err => {
                console.log(err)
            })


    }

    handleChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
    };

    handleSubmit(event) {
        event.preventDefault()
        console.log(event)


        // TODO: verfiy field 


        const { email, firstname, lastname, username, password1, password2, currentDate } = this.state

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
            birthdate: currentDate
        }

        this.register(payload)
        // eveything correct:
    }
    constructor(props) {
        super(props);

        this.state = {
            currentDate: new Date()
        };

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.register = this.register.bind(this)

    }


    render() {
        return (
            <div>


                <form onChange={this.handleChange} onSubmit={this.handleSubmit}>
                    <div className="form-row">
                        <div className="col-md-4 mb-3">
                            <label>Email</label>
                            <input type="email" className="form-control" name="email" placeholder="Email" />
                        </div>
                        <div className="col-md-4 mb-3">
                            <label>Firstname</label>
                            <input type="text" className="form-control" name="firstname" placeholder="First name" />
                        </div>
                        <div className="col-md-4 mb-3">
                            <label>Lastname</label>
                            <input type="text" className="form-control" name="lastname" placeholder="Last name" />
                        </div>
                        <div className="col-md-4 mb-3">
                            <label>Username</label>
                            <input type="text" className="form-control" name="username" placeholder="username" />
                        </div>
                        <div className="col-md-4 mb-3">
                            <label>Password</label>
                            <input type="password" className="form-control" name="password1" placeholder="password" />
                        </div>
                        <div className="col-md-4 mb-3">
                            <label>Comfilm Password</label>
                            <input type="password" className="form-control" name="password2" placeholder="password" />
                        </div>
                        <div className="col-md-4 mb-3">
                            <label>Birthdate</label>
                            {/* <DatePicker
                                className="form-control"
                                selected={this.state.startDate}
                                onChange={this.handleChange}
                            /> */}
                        </div>
                    </div>

                    <div className="form-group">
                        <div className="form-check">
                            <input className="form-check-input" type="checkbox" value="" id="invalidCheck2" required />
                            <label className="form-check-label" >
                                Agree to terms and conditions
                </label>
                        </div>
                    </div>
                    <button className="btn btn-primary" type="submit">Submit form</button>
                </form>
            </div>
        )
    }
}

export default Register