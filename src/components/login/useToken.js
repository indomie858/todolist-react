// custom hook for login token
import { useState } from 'react'

export default function useToken() {
  //gets token from memory
  const getToken = () => {
    //tokens are stored locally so user doesn't have to keep logging in
    const tokenString = sessionStorage.getItem('token');
    //const userToken = JSON.parse(tokenString);
    return tokenString
  };

  const [token, setToken] = useState(getToken());

  //saves token to state
  const saveToken = userToken => {
    //tokens are stored locally so user doesn't have to keep logging in
    sessionStorage.setItem('token', userToken);
    setToken(userToken);
  };

  return {
    setToken: saveToken,
    token
  }
}
