//component for login page. feel free to change anything
import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'
import firebase from "firebase";
import { useHistory } from "react-router-dom";






//passes user login info to backend
async function loginUser(credentials, callback) {
    //logs in with firebase and gets credentials
  //let user
  firebase.auth().signInWithEmailAndPassword( credentials.username, credentials.password)
  .then((userCredential) => {
    // Signed in 
    let user = userCredential.user;
    callback(user);
    // ...
  })
  .catch((error) => {
    const errorCode = error.code;
    const errorMessage = error.message;
    console.log(errorCode)
    console.log(errorMessage)
    // ..
  });

  //return user ? user:null;
    
    // return fetch('http://localhost:3003/userLogin', {
    //     method: 'POST',
    //     headers: {
    //         'Content-Type': 'application/json'
    //     },
    //     body: JSON.stringify(credentials)
    // })
    //     .then(data => data.json())
}


async function loginAppleUser(callback){

  var provider = new firebase.auth.OAuthProvider('apple.com');

  firebase
    .auth()
    .signInWithPopup(provider)
    .then((result) => {
      /** @type {firebase.auth.OAuthCredential} */
      var credential = result.credential;

      // The signed-in user info.
      var user = result.user;

      // You can also get the Apple OAuth Access and ID Tokens.
      var accessToken = credential.accessToken;
      var idToken = credential.idToken;

      callback(user)
      // ...
    })
    .catch((error) => {
      // Handle Errors here.
      var errorCode = error.code;
      var errorMessage = error.message;
      // The email of the user's account used.
      var email = error.email;
      // The firebase.auth.AuthCredential type that was used.
      var credential = error.credential;

      // ...
    });
}



//passes user registration info to backend
//haven't done anything with this yet - gaven
async function registerUser(credentials,callback) {

  let user
  firebase.auth().createUserWithEmailAndPassword(credentials.username, credentials.password)
  .then((userCredential) => {
    // Registered.  Maybe return user to sign in?
    user = userCredential.user;
    user.displayName = `${credentials.firstName} ${credentials.lastName}`
    callback(user)
    
    // ...
  })
  .catch((error) => {
    const errorCode = error.code;
    const errorMessage = error.message;
    console.log(errorCode)
    console.log(errorMessage)
    // ..
  });
  // console.log(user)

    // return fetch('http://localhost:3003/userCreate', {
    //     method: 'POST',
    //     headers: {
    //         'Content-Type': 'application/json'
    //     },
    //     body: JSON.stringify(credentials)
    // })
    //     .then(data => data.json())
}

const Login = ({ setToken, handleGoogleAuth /*Function to call for google auth*/ }) => {
    //states for sign in username and password
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    //boolean to track which form needs to be displayed
    const [isRegistered, setIsRegistered] = useState(true);

    //states for signup form
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const history = useHistory();
    function returnLogin() {
      history.push("/login");
    }

    const getToken = () => {
      //tokens are stored locally so user doesn't have to keep logging in
      const token = sessionStorage.getItem('token');
      return token
    };

    //function for handling login submission
    const handleSubmitLogin = async e => {
        //stops page from reloading on form submission
        e.preventDefault();

        //gets response from api
        const token = await loginUser({
            username,
            password,
            isRegistered
        }, (token) =>{ sessionStorage.setItem('token',JSON.stringify(token));setToken(token); console.log(token)});
        //setToken(token);
        //console.log(token);
        //once token is set, home page renders
    }

    //function for handling registration submission
    const handleSubmitRegister = async e => {
        //stops page from reloading on form submission
        e.preventDefault();

        //gets response from api
        const token = await registerUser({
            username,
            password,
            firstName,
            lastName,
            isRegistered
        },(user)=>{if (user){
          user.displayName=firstName+" "+lastName;

          firebase.auth().signOut().then(() => {
            // Sign-out successful.
            console.log("The User has been logged out.")
            sessionStorage.setItem('token', "");
            returnLogin();
            setIsRegistered(true);
          }).catch((error) => {
            // An error happened.
            console.log(error)
          });
        }});
        // setToken(token);
        // console.log(token);
        //once token is set, home page renders
    }

    const handleGoogleSignIn = async e => {
      //stops page from reloading on form submission
      e.preventDefault();


      //Popup googleAuth on login render
      if(!getToken()) {
        handleGoogleAuth( (err, email, token) => {
          console.log("err: " + err);
          console.log("email: " + email);
          console.log("token: " + token);
          if(!err){ //if no errors on auth
            if(email){
              if(token){
                sessionStorage.setItem('token', token);
                sessionStorage.setItem('email', email);
                console.log('Everything worked');
              } else {
                console.log('auth token failed: ' + token);
              }
            } else {
              console.log('auth email failed: ' + email);
            }
          } else {
            console.log('auth failed: ' + err);
          }
        });
      } else {
        console.log('token was registered');
      }
    }

    // form rendered depends on if user is registered (login or register form)
    if (getToken()){
      console.log('Used this home oute login.usename = true');
      //this.history.push('/home');
      return (
         <Redirect to="/home" />
      );
    }
    if (isRegistered) {
        return (
            <>
                <LoginForm handleSubmit={handleSubmitLogin} setUsername={setUsername} setPassword={setPassword} setIsRegistered={setIsRegistered} />
            </>
        )
    } else {
        return (
            <>
                <RegisterForm handleSubmit={handleSubmitRegister} setUsername={setUsername} setPassword={setPassword} setFirstName={setFirstName} setLastName={setLastName} setIsRegistered={setIsRegistered} />
            </>
        )
    }
}

Login.propTypes = {
    setToken: PropTypes.func.isRequired
}

export default Login
