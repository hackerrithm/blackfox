import React from "react";
import MButton from "../general/reusable/button/mbutton";
import { Grid, TextField, Button } from "@material-ui/core";
import useStyles from "./styles";
import "./styles.css"
import { Link } from "react-router-dom";
import CompanyIcon from "../general/reusable/icons/company";


const LoginForm = ({ username, password, isLoading, dispatch }: any) => {
        const classes = useStyles("");
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
                                                        // if (validateLoginForm(username, password)) {
                                                        // 	authHandler();
                                                        // }
                                                }}
                                        >
                                                <CompanyIcon />
                                                <div className="header">Login with username</div>
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
                                                                onChange={(e: any) => {
                                                                        dispatch({
                                                                                type: "field",
                                                                                fieldName: "username",
                                                                                payload: e.currentTarget.value
                                                                        });
                                                                }}
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
                                                                onChange={(e: any) => {
                                                                        console.log("changed");

                                                                        dispatch({
                                                                                type: "field",
                                                                                fieldName: "password",
                                                                                payload: e.currentTarget.value
                                                                        });
                                                                }}
                                                        />
                                                </div>
                                                {/* <div className="form-group">
                                                        <label htmlFor="username">Username</label>
                                                        <input
                                                                type="text"
                                                                name="username"
                                                                placeholder="username"
                                                                value={username}
                                                                onChange={(e: any) => {
                                                                        dispatch({
                                                                                type: "field",
                                                                                fieldName: "username",
                                                                                payload: e.currentTarget.value
                                                                        });
                                                                }}
                                                        />
                                                </div>
                                                <div className="form-group">
                                                        <label htmlFor="password">Password</label>
                                                        <input
                                                                type="text"
                                                                name="password"
                                                                placeholder="password"
                                                                value={password}
                                                                onChange={(e: any) => {
                                                                        console.log("changed");

                                                                        dispatch({
                                                                                type: "field",
                                                                                fieldName: "password",
                                                                                payload: e.currentTarget.value
                                                                        });
                                                                }}
                                                        />
                                                </div> */}
                                                <div>
                                                        <Button variant="contained" color="primary" className={classes.submitButton}>
                                                                {isLoading ? "Logging in..." : "Login"}
                                                        </Button>
                                                </div>
                                                <div className="account-creation-link">
                                                <label color="inherit">Need an account? </label>
                                                        <Link
                                                                to="/register"
                                                        >
                                                                <label>Create one</label>
                                                        </Link>
                                                </div>
                                        </form>
                                </div>
                        </Grid>
                        <Grid item xs={3} sm={3}>
                                <div className={"right-panel"}>
                                </div>
                                {/* <div className={"ads-data-right"}>
                                <h3>This is an Ad</h3>
                        </div> */}
                        </Grid>
                </Grid>
        )
}

export default LoginForm;