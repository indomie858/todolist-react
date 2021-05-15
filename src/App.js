import TestComponent from './components/TestComponent';

import { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import Home from './components/Pages/Home';
import Preferences from './components/Pages/Preferences';
import Login from './components/login/Login'
import useToken from './components/login/useToken'


const App = () => {
  //login token with custom hook:  /components/login/useToken
  const { token, setToken } = useToken();

  if (!token) {
    //if there is no user token, then render login page

    // NOTE: COMMENT THIS LINE TO HIDE LOGIN PAGE
    //return <Login setToken={setToken} />
  }


  return (
    <>
      <TestComponent />
      {/* main container for ui */}
      <Container maxWidth="xs">
        <Home />
      </Container>
    </>
  );
}

export default App;
