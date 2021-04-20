//component for login page. feel free to change anything
//might refactor this to use material ui components - gaven
import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

//passes user login info to backend
async function loginUser(credentials) {
    //replace url with correct api
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
    //replace url with correct api
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
    //states for username and password
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();
    const [isRegistered, setIsRegistered] = useState(true);

    const handleSubmit = async e => {
        //stops page from reloading on form submission
        e.preventDefault();

        // TODO: handle registration info

        //gets response from api
        const token = await loginUser({
            username,
            password
        });
        setToken(token);
        console.log(token);
    }

    // form rendered depends on if user is registered (login or register form)
    if (isRegistered) {
        return (
            <>
                <LoginForm handleSubmit={handleSubmit} setUsername={setUsername} setPassword={setPassword} setIsRegistered={setIsRegistered} />
            </>
        )
    } else {
        return (
            //TODO: work on registration form. there's nothing handling form entries atm
            <>
                <RegisterForm setIsRegistered={setIsRegistered} />
            </>
        )
    }
}

Login.propTypes = {
    setToken: PropTypes.func.isRequired
}

export default Login
