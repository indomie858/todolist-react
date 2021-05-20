//import { useState } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import Home from './components/Pages/Home';
//import Preferences from './components/Pages/Preferences';
import Login from './components/login/Login';
import useToken from './components/login/useToken';
//import firebase from 'firebase/app';
// import firebase from "firebase";
//import { FirebaseAuthProvider, FirebaseAuthConsumer} from "@react-firebase/auth";




const App = ({ handleGoogleAuth }) => {
  //login token with custom hook:  /components/login/useToken
  const { setToken } = useToken();

  // comment out if connected to firebase auth
  // sessionStorage.setItem('token', 'cool token');
  // sessionStorage.setItem('email', 'test@gmail.com');

  return (
    <>
      <Container maxWidth="xs">
        <Router>
          <Switch>
            <Route
              exact
              path="/login"
              render={() => (<Login setToken={setToken} handleGoogleAuth={handleGoogleAuth} />)}
            />
            <Route
              exact
              path="/home"
              render={() => (<Home />)}
            />
            <Route
              exact
              path="/"
              render={() => (<Home />)}
            />
          </Switch>
        </Router>
      </Container>
    </>
  );
}

export default App;
