import React from "react";
import MButton from "../general/reusable/button/mbutton";
import { Grid, TextField, Button } from "@material-ui/core";
import useStyles from "./styles";
import "./styles.css"
import { Link } from "react-router-dom";
import CompanyIcon from "../general/reusable/icons/company";


const LoginForm = ({ loginUserSubmit, error, formError, data, handleInputChange, username, password, isLoading, dispatch }: any) => {
        const classes = useStyles("");
        return (
                <Grid container>
                        <Grid item xs={3} sm={3}>
                                <div className={"left-panel"}>
                                </div>
                        </Grid>
                        <Grid item xs={6} sm={6}>
                                <div className={"mid-view"}>
                                {error ? <p>Oh no! {error.message}</p> : null}
			        {data && loginUserSubmit ? <p>Login!</p> : null}
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
                                                                onChange={handleInputChange}
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
                                                                onChange={handleInputChange}
                                                        />
                                                </div>
                                                {formError && (
					<span className="form-error">{formError}</span>
				)}
                                                <div>
                                                        <Button variant="contained" color="primary" className={classes.submitButton} disabled={isLoading} onClick={() => loginUserSubmit()}>
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