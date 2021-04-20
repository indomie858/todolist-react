// custom hook for login token
import { useState } from 'react'

export default function useToken() {
  //gets token from memory
  const getToken = () => {
    //tokens are stored locally so user doesn't have to keep logging in
    const tokenString = localStorage.getItem('token');
    const userToken = JSON.parse(tokenString);
    return userToken?.token
  };

  const [token, setToken] = useState(getToken());

  //saves token to state
  const saveToken = userToken => {
    //tokens are stored locally so user doesn't have to keep logging in
    localStorage.setItem('token', JSON.stringify(userToken));
    setToken(userToken.token);
  };

  return {
    setToken: saveToken,
    token
  }
}