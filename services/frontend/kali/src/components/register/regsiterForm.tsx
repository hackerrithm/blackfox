import React from "react";
import { useHistory } from "react-router-dom";
import { Grid, TextField, Button } from "@material-ui/core";
import useStyles from "./styles";
import "./styles.css"
import { Link } from "react-router-dom";
import CompanyIcon from "../general/reusable/icons/company";


const RegisterForm = ({ username, password, emailaddress, gender, error, registerUser, data, setUsername, setEmailAddress, setPassword }: any) => {
        const classes = useStyles("");
        const history = useHistory();
        return (
                <Grid container>
                        <Grid item xs={3} sm={3}>
                                <div className={"left-panel"}>
                                </div>
                        </Grid>
                        <Grid item xs={6} sm={6}>
                                <div className={"mid-view"}>
                                        <form
                                                className={classes.root}
                                                onSubmit={(e: any) => {
                                                        e.preventDefault();
                                                        console.log("submit");
                                                        
                                                }}
                                        >
                                                <CompanyIcon />
                                                <div className="header">Create an account</div>
                                                <div>
                                                        {error ? <p>Oh no! {error.message}</p> : null}
                                                        {data && registerUser ? history.push("/profile") : null}
                                                </div>
                                                <div className="emailAddress-form-group">
                                                        <TextField
                                                                label="Email Address"
                                                                id="outlined-emailaddress-small"
                                                                className={classes.emailAddress}
                                                                defaultValue="Email Address"
                                                                variant="outlined"
                                                                size="small"
                                                                name="emailaddress"
                                                                value={emailaddress}
                                                                onChange={e => setEmailAddress(e.target.value)}
                                                        />
                                                </div>
                                                <div className="username-form-group">
                                                        <TextField
                                                                label="Username"
                                                                id="outlined-username-small"
                                                                className={classes.username}
                                                                defaultValue="Username"
                                                                variant="outlined"
                                                                size="small"
                                                                name="username"
                                                                value={username}
                                                                onChange={e => setUsername(e.target.value)}
                                                        />
                                                </div>
                                                <div className="password-form-group">
                                                        <TextField
                                                                label="Password"
                                                                id="outlined-password-small"
                                                                className={classes.password}
                                                                type="password"
                                                                defaultValue="Password"
                                                                variant="outlined"
                                                                size="small"
                                                                name="password"
                                                                value={password}
                                                                onChange={e => setPassword(e.target.value)}
                                                        />
                                                </div>
                                                <div>
                                                        <Button variant="contained" 
                                                                color="primary" 
                                                                className={classes.submitButton}
                                                                onClick={() => registerUser()}>
                                                                Create an account
                                                        </Button>
                                                </div>
                                                <div className="account-existing-link">
                                                        <label color="white">Already have an account? </label>
                                                        <Link
                                                                to="/login"
                                                        >
                                                                <label>Login here</label>
                                                        </Link>
                                                </div>
                                        </form>
                                </div>
                        </Grid>
                        <Grid item xs={3} sm={3}>
                                <div className={"right-panel"}>
                                </div>
                        </Grid>
                </Grid>
        )
}

export default RegisterForm;