import React from 'react';
import axios from 'axios';
import Filter from "./components/Filter"

import {
    Link
} from "react-router-dom"

class HotelList extends React.Component {
    //TODO: searchbar

    constructor(props) {
        super(props);
        this.readHotels();
        this.state = {hotels: [],name:''};
    
        this.readHotels = this.readHotels.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    readHotels(name) {
        let url = window.global.api_location+'/hotels'
        console.log(name)
        if (name != undefined && name != ''){
            url = url + "?name="+name
        }
        const self = this;
        axios.get(url).then(function(response) {
            console.log(response.data);
            self.setState({hotels: response.data});
        }).catch(function (error){
            console.log(error);
        });
    }
    getHotels() {
        let table = []

        for (let i=0; i < this.state.hotels.length; i++) {
            let urlHotel = "/hotels/"+this.state.hotels[i].name
            table.push(
            <tr key={i}>
                <th scope="row">{i}</th>
                <td>{this.state.hotels[i].name}</td>
                <td>{this.state.hotels[i].price}</td>
                <td>{this.state.hotels[i].detail}</td>
                <td>{this.state.hotels[i].height}</td>
                <td><Link to={urlHotel}>HomeInfo</Link></td>
            </tr>
            );
        }

        return table
    }

    handleChange(event) {
        this.setState({name: event.target.value})
    }

    handleSubmit(event){
        console.log(this.state.name)
        this.readHotels(this.state.name) 
        event.preventDefault();
    }
  

    render() {
        return (
            <div className="container">
            <form onSubmit={this.handleSubmit}>
                <label>
                    <Filter value={this.state.name} handleChange={this.handleChange}/>
                </label>
                    <input type="submit" value="Submit" />
            </form>
           
            <table className="table">
                <thead className="thead-dark">
                    <tr>
                        <th scope="col"> 
                            #
                        </th>
                        <th scope="col"> 
                            Name
                        </th>
                        <th scope="col">
                            Price
                        </th>
                        <th scope="col">
                            Detail
                        </th>
                        <th scope="col">
                            Height
                        </th>
                        <th scope="col">
                            Link
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {this.getHotels()}
                </tbody>
            </table>
        </div>
        )
    }
}

export default HotelList