import React, { useEffect, useState } from "react";
import Grid from "@material-ui/core/Grid";
import "./styles.css"
import CountdownTimer from "../general/reusable/countdown/countdown";

const IndexStartView = () => {
        return (
                <div>
                        <Grid container>
                                <Grid item xs={2} sm={2}>
                                        <div className={"left-panel"}>
                                        </div>
                                </Grid>
                                <Grid item xs={8} sm={8}>
                                        <div className={"mid-view"}>
                                                <h2 className="coming-soon-view">we are here</h2>
                                        </div>
                                </Grid>
                                <Grid item xs={2} sm={2}>
                                        <div className={"right-panel"}>
                                        </div>
                                </Grid>
                        </Grid>
                </div>
        );
}

const AppLaunchView = () => {
        return (
                <div>
                        <Grid container>
                                <Grid item xs={2} sm={2}>
                                        <div className={"left-panel"}>
                                        </div>
                                </Grid>
                                <Grid item xs={8} sm={8}>
                                        <div className={"mid-view"}>
                                                <h2 className="coming-soon-view">coming soon</h2>

                                                <h4 className="info-text-view">We will be celebrating the launch of our new site very soon!</h4>

                                                <div className="countdown-timer-view">
                                                        <CountdownTimer /> 
                                                </div>  
                                        </div>
                                </Grid>
                                <Grid item xs={2} sm={2}>
                                        <div className={"right-panel"}>
                                        </div>
                                </Grid>
                        </Grid>
                </div>
        );
}

const Start = () => {
        return (
                <div>
                        {2 > 1? <AppLaunchView/>: <IndexStartView/>}
                </div>
        );
}

export default Start;