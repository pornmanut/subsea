import React from 'react'
import axios from 'axios';
// booking: 1
// detail: "great view"
// height: 30.3
// id: "5f2ce2ea0365c2ae52c7d39e"
// max: 1
// name: "abc"
// price: 300
class Hotel extends React.Component {
    readData(name) {
        const self = this;
        axios.get(window.global.api_location+'/hotels/'+name).then(function(response) {
            console.log(response.data);

            self.setState({hotel: response.data,found: true});
        }).catch(function (error){
            console.log(error);
            self.setState({found: false})
        });
    }
    constructor(props) {
        super(props);
        this.state = {hotel: ""};
        this.readData(props.name)
        this.readData = this.readData.bind(this);
    }

    render() {
        if (!this.state.found){
            return <div>loading</div>
        }
        return (
            <div className="container">
                <p>{this.state.hotel.name}</p>
                <p>{this.state.hotel.price}</p>
                <p>{this.state.hotel.detail}</p>
            <p>{this.state.hotel.booking}/{this.state.hotel.max}</p>
            </div>
        )
    }
}

export default Hotel