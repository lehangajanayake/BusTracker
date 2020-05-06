import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { createAppContainer } from 'react-navigation';
import { createStackNavigator } from 'react-navigation-stack'
import { Map } from './screens/Map';
import { Home } from './screens/Home';
import { Details } from './screens/Details'

const MainNavigator = createStackNavigator({
  Home: {screen: Home},
  Map: {screen: Map},
  Details: {screen: Details}
});

const App = createAppContainer(MainNavigator);
export default App;