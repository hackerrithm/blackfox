import * as React from "react";
// import Button from "@material-ui/core/Button";
import { useMutation } from '@apollo/react-hooks';
import gql from 'graphql-tag';
import { useState } from "react";
import "./style.css";

const SAVE_USER = gql`
  mutation registerUser($user: RegisterInput!) {
    registerUser(user: $user)
  }
`;

interface User {
    username: string;
    password: string;
    firstname: string;
    lastname: string;
    emailAddress: string;
    gender: string;
  }

interface RegisterDetails {
    username: string;
    password: string;
    firstname: string;
    lastname: string;
    emailAddress: string;
    gender: string;
}

export default function Register() {
    // constructor(props: any) {
    //     super(props);
    // }
    // componentDidMount() { 
    //     console.log('here bruh');
    // }
    
    /**
     * name
     */
    function register() {
        console.log('register here');
    }

    
    /**
     * Render Register
     */
    // public render() {
        const [username, setUsername] = useState<any>('');
        const [firstname, setFirstname] = useState<any>('');
        const [lastname, setLastname] = useState<any>('');
        const [emailAddress, setEmailAddress] = useState<any>('');
        const [password, setPassword] = useState<any>('');
        const [gender, setGender] = useState<any>('');


        const [registerUser, { error, data }] = useMutation<
        { registerUser: User },
        { user: RegisterDetails }
        >(SAVE_USER, {
            variables: { user: { username, 
                password, 
                firstname, 
                lastname, 
                emailAddress, 
                gender } }
        });
        return (
                <div className="base-container">
                    <div className="header">Register</div>
                    {error ? <p>Oh no! {error.message}</p> : null}
                    {data && registerUser ? <p>Saved!</p> : null}
                    <div className="form">
                        <div className="form-group">
                        <label htmlFor="username">Username</label>
                        <input type="text" name="username" placeholder="username" onChange={e => setUsername(e.target.value)}/>
                        </div>
                        <div className="form-group">
                        <label htmlFor="firstname">Firstname</label>
                        <input type="text" name="firstname" placeholder="firstname" onChange={e => setFirstname(e.target.value)} />
                        </div>
                        <div className="form-group">
                        <label htmlFor="lastname">Lastname</label>
                        <input type="text" name="lastname" placeholder="lastname" onChange={e => setLastname(e.target.value)} />
                        </div>
                        <div className="form-group">
                        <label htmlFor="password">Password</label>
                        <input type="text" name="password" placeholder="password" onChange={e => setPassword(e.target.value)}/>
                        </div>
                        <div className="form-group">
                        <label htmlFor="email">Email</label>
                        <input type="text" name="email" placeholder="email" onChange={e => setEmailAddress(e.target.value)} />
                        </div>
                        <div className="form-group">
                        <label htmlFor="gender">Gender</label>
                        <input type="text" name="gender" placeholder="gender" onChange={e => setGender(e.target.value)}/>
                        </div>
                    </div>
                    <div className="footer">
                    <button type="button" className="btn" onClick={() => registerUser()}>
                        Register
                    </button>
                    </div>
                </div>     
        );
    // }
}
