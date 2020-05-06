"use strict";
exports.__esModule = true;
var react_native_1 = require("react-native");
var Map_1 = require("./screens/Map");
// function App() {
//   return (
//     <View style={styles.container}>
//       <Text>Open up App.tsx to start working on your app!</Text>
//     </View>
//   );
// }
var styles = react_native_1.StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center'
    }
});
exports["default"] = Map_1.Map;
