//component for login page. feel free to change anything
//might refactor this to use material ui components - gaven
import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';

import './Login.css';
import { GolfCourseSharp, LocalGroceryStoreSharp } from '@material-ui/icons';

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
    const [username, setUserName] = useState();
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

    // this block is gross, might refactor it into seperate components later on -gaven
    if (isRegistered){
        return (
            <div className="login-wrapper">
                <h1>Log In</h1>
                <form onSubmit={handleSubmit}>
                    <label>
                        <p>Username</p>
                        <input type="text" onChange={e => setUserName(e.target.value)} />
                    </label>
                    <label>
                        <p>Password</p>
                        <input type="password" onChange={e => setPassword(e.target.value)} />
                    </label>
                    <div>
                        <button type="submit">Submit</button>
                    </div>
                </form>
                <button onClick={() => setIsRegistered(false)}>Don't have an account? Register here</button>
            </div>
        )
    } else {
        return (
            <div className="login-wrapper">
                <h1>Register</h1>
                <form onSubmit={handleSubmit}>
                    <label>
                        <p>New Username</p>
                        <input type="text" onChange={e => setUserName(e.target.value)} />
                    </label>
                    <label>
                        <p>New Password</p>
                        <input type="password" onChange={e => setPassword(e.target.value)} />
                    </label>
                    <label>
                        <p>Enter Password again</p>
                        <input type="password" onChange={e => setPassword(e.target.value)} />
                    </label>
                    <div>
                        <button type="submit">Submit</button>
                    </div>
                </form>
                <button onClick={() => setIsRegistered(true)}>Already have an account? Log in here</button>
            </div>
        )
    }
}

Login.propTypes = {
    setToken: PropTypes.func.isRequired
}

export default Login
