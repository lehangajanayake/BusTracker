import { Component } from 'react';
import {View, Text, TextInput, Button} from 'react-native';
import { Bus } from '../pkg/models';


interface props{
    navigation: any;
}
interface state{
    bus: Bus;
}
export class Home extends Component<props, state> {
    static navigationOptions = {
        title: 'Home',
    };
    private onPress(){
       this.props.navigation.navigate('Map', {bus: {
           LicenseNo:  this.state.bus.LicenseNo
       }});
    }
    render(){
        return(
            <View>
                <Text>Hi Enter the Bus LicenseNo</Text>
                <TextInput
                    autoFocus = {true}
                    onChangeText={(text) => this.setState({bus: {LicenseNo: text}})}
                />
                <Button 
                    title = "Submit"
                    onPress={() => this.onPress()} 
                />
            </View>
            
        );
    }
}

