import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import firebase from 'firebase/app';
import "firebase/auth";


//Grabs all env variables from .env.local file
const {REACT_APP_API_KEY, REACT_APP_PROJECT_ID, REACT_APP_STORAGE_BUCKET, REACT_APP_AUTH_DOMAIN} = process.env;
console.log(REACT_APP_PROJECT_ID);
const config = {
  apiKey: "AIzaSyAHUKJNQj77dcxX0JeCrw1ZniwVZ52uwus",
  authDomain: "friday-584.firebaseapp.com",
  projectId: "friday-584",
  storageBucket: "friday-584.appspot.com",
  messagingSenderId: "748113457358",
  appId: "1:748113457358:web:08e92a55354cecf740a09f",
  measurementId: "G-2DC9L5VCCX"
}

firebase.initializeApp(config);
var provider = new firebase.auth.GoogleAuthProvider();
function handleGoogleAuth(callback){
  var err, email, token = null;
  firebase.auth()
    .signInWithPopup(provider)
    .then((result) => {
      /** @type {firebase.auth.OAuthCredential} */
      //var credential = result.credential;

      // This gives you a Google Access Token. You can use it to access the Google API.
      token = result.credential.accessToken;
      sessionStorage.setItem('token', JSON.stringify(token));
      // The signed-in user info.
      email = result.user.email;
      callback(err, email, token);
      // ...
    }).catch((error) => {
      // Handle Errors here.
      //var errorCode = error.code;
      err = error.message;
      // The email of the user's account used.
      email = error.email;
      // The firebase.auth.AuthCredential type that was used.
      //var credential = error.credential;
      // ...
    });
}

ReactDOM.render(
  <React.StrictMode>
    <App handleGoogleAuth={handleGoogleAuth}/>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
