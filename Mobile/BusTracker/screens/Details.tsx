import React, { Component } from 'react';
import {View, StyleSheet, Text} from 'react-native';
import { Bus } from '../BusTracker/pkg/models';
import { getBusAttributes, getBusAvailability } from '../BusTracker/pkg/api';

interface props{
    navigation: any;
}
interface state{
    bus: Bus
}

export class Details extends Component<props, state>{
    static ={
        navigationOptions: Details
    }

    componentDidMount(){
        this.setState({bus: this.props.navigation.state});
        getBusAttributes(this.state.bus.LicenseNo!).then(value => {
            this.setState({bus: value!})
        });
        setInterval(() => getBusAvailability(this.state.bus).then(value => {
            if (this.state.bus.Attributes?.Availability != value?.Attributes?.PathNo){
                this.setState({bus: value!});
            }
            return;
        }), (1000*60));
    
    }
    render(){
        return(
            <View>
                <Text>hI THIS SHOW THE DETAILS ABOUT THE BUS</Text>
                <Text>Bus PathNo: {this.state.bus.Attributes?.PathNo}</Text>
                <Text>Bus AC: {this.state.bus.Attributes?.AC}</Text>
                <Text>Bus Availability: {this.state.bus.Attributes?.Availability}</Text>
            </View>
        );
    }
}
