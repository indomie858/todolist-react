//component for login page. feel free to change anything
import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

//passes user login info to backend
async function loginUser(credentials) {
    //replace url with correct endpoint
    return fetch('http://localhost:3003/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

//passes user registration info to backend
//haven't done anything with this yet - gaven
async function registerUser(credentials) {
    //replace url with correct endpoint
    return fetch('http://localhost:3003/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

const Login = ({ setToken }) => {
    //states for sign in username and password
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    //boolean to track which form needs to be displayed
    const [isRegistered, setIsRegistered] = useState(true);

    //states for signup form
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");

    //function for handling login submission
    const handleSubmitLogin = async e => {
        //stops page from reloading on form submission
        e.preventDefault();

        //gets response from api
        const token = await loginUser({
            username,
            password,
            isRegistered
        });
        setToken(token);
        console.log(token);
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
        });
        setToken(token);
        console.log(token);
        //once token is set, home page renders
    }

    // form rendered depends on if user is registered (login or register form)
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
