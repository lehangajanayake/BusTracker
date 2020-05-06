import React from 'react';
import { View, Button, TextInput } from 'react-native';
import Map from './screens/Map';


class Main extends React.Component{
  render(){
    return(
      <View>
        <TextInput
          name = 'PathNo'
        />
        <Button 
          title = "Hi I'm a button"
        />
      </View>
    )
  }
}

export default Map;

