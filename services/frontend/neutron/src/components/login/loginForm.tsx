import React from "react";
import MButton from "../general/reusable/button/mbutton";
import { Grid } from "@material-ui/core";


const LoginForm = ({ username, password, isLoading, dispatch }: any) => {
        return (
                <Grid container>
                        <Grid item xs={3} sm={3}>
                                <div className={"left-panel"}>
                                </div>
                        </Grid>
                        <Grid item xs={6} sm={6}>
                                <div className={"mid-view"}>
                                        <div className="header">Login</div>
                                        <form
                                                className="form"
                                                onSubmit={(e: any) => {
                                                        e.preventDefault();
                                                        // if (validateLoginForm(username, password)) {
                                                        // 	authHandler();
                                                        // }
                                                }}
                                        >
                                                <div className="form-group">
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
                                                </div>
                                                <MButton
                                                        className="submit"
                                                        type="submit"
                                                        disabled={isLoading}
                                                        color="primary"
                                                        variant="contained"
                                                        label={isLoading ? "Logging in..." : "Log In"}
                                                />
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