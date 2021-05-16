import { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import Home from './components/Pages/Home';
import Preferences from './components/Pages/Preferences';
import Login from './components/login/Login';
import useToken from './components/login/useToken';
import firebase from 'firebase/app';
import "firebase/auth";
//import { FirebaseAuthProvider, FirebaseAuthConsumer} from "@react-firebase/auth";

const App = ({ handleGoogleAuth }) => {
  //login token with custom hook:  /components/login/useToken
  const { token, setToken } = useToken();

  return (
    <>
      <Container maxWidth="xs">
        <BrowserRouter>
          <Switch>        
            <Route path="/login">
              <Login setToken={setToken} handleGoogleAuth={handleGoogleAuth} />
            </Route>
            <Route path="/home">
              <Home />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </BrowserRouter>
      </Container>
    </>
  );
}

export default App;
