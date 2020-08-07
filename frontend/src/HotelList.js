import React from 'react';
import axios from 'axios';

class HotelList extends React.Component {
    readData() {
        const self = this;
        axios.get(window.global.api_location+'/hotels').then(function(response) {
            console.log(response.data);

            self.setState({hotels: response.data});
        }).catch(function (error){
            console.log(error);
        });
    }
    getHotels() {
        let table = []

        for (let i=0; i < this.state.hotels.length; i++) {

            table.push(
            <tr key={i}>
                <th scope="row">{i}</th>
                <td>{this.state.hotels[i].name}</td>
                <td>{this.state.hotels[i].price}</td>
                <td>{this.state.hotels[i].detail}</td>
            </tr>
            );
        }

        return table
    }
    constructor(props) {
        super(props);
        this.readData();
        this.state = {hotels: []};
    
        this.readData = this.readData.bind(this);
    }

    render() {
        return (
            <div className="container">
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