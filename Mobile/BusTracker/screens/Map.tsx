import React, { Component } from 'react';
//import {} from '../pkg/api';
import {StyleSheet, Dimensions, Text} from 'react-native';
import MapView, {Marker} from 'react-native-maps';
import {getBusLocation, getBusAttributes} from '../pkg/api';
import {Bus} from '../pkg/models';

interface props{
    navigation: any;
}
interface state{
    bus: Bus;
}

export class Map extends Component<props, state>{
    static navigationOptions = {
        title: 'Map',
    };
    componentDidMount(){
        //this.bus = getBusLocation(this.props.bus.LicenseNo!)
        //var that = this;
        this.setState({bus: this.props.navigation.state});
        getBusLocation(this.state.bus.LicenseNo!).then((value) => {
            this.setState({bus:value!})
        });
        setInterval(() =>  getBusLocation(this.state.bus.LicenseNo!).then((value) => {
            this.setState({bus: value!})
        }), 1000) ;
        console.log(this.state.bus.Location!.Latitude);
    }

    render(){
        return(
           <MapView>
               <Marker 
               coordinate = {{latitude:this.state.bus.Location!.Latitude, longitude: this.state.bus.Location!.Longitude}}
               title = {this.state.bus.LicenseNo}
               //image = {require('../assets/Bus_icon.png')}
               />
           </MapView>
        );
    }
}

const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff',
      alignItems: 'center',
      justifyContent: 'center',
    },
  });