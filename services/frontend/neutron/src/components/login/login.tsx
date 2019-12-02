import * as React from "react";
// import Button from "@material-ui/core/Button";
import { useMutation } from '@apollo/react-hooks';
import gql from 'graphql-tag';
import { useState } from "react";

const LOGIN_USER = gql`
  mutation loginUser($user: LoginInput!) {
    loginUser(user: $user)
  }
`;

interface User {
    username: string;
    password: string;
  }

interface LoginDetails {
    username: string;
    password: string;
}

export default function Login() {
    // constructor(props: any) {
    //     super(props);
    // }
    // componentDidMount() { 
    //     console.log('here bruh');
    // }
    
    /**
     * name
     */
    function login() {
        console.log('login here');
    }

    
    /**
     * Render Login
     */
    // public render() {
        const [username, setUsername] = useState<any>('');
        const [password, setPassword] = useState<any>('');


        const [loginUserSubmit, { error, data }] = useMutation<
        { loginUser: User },
        { user: LoginDetails }
        >(LOGIN_USER, {
            variables: { user: { username, password } }
        });
        return (
                <div className="base-container">
                    <div className="header">Login</div>
                    {error ? <p>Oh no! {error.message}</p> : null}
                    {data && loginUserSubmit ? <p>Login!</p> : null}
                    <div className="form">
                        <div className="form-group">
                        <label htmlFor="username">Username</label>
                        <input type="text" name="username" placeholder="username" onChange={e => setUsername(e.target.value)}/>
                        </div>
                        <div className="form-group">
                        <label htmlFor="password">Password</label>
                        <input type="text" name="password" placeholder="password" onChange={e => setPassword(e.target.value)}/>
                        </div>
                    </div>
                    <div className="footer">
                    <button type="button" className="btn" onClick={() => loginUserSubmit()}>
                        Login
                    </button>
                    </div>
                </div>     
        );
    // }
}
