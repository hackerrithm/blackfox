import React, { useContext } from "react";
import Grid from "@material-ui/core/Grid";
import "./styles.css"
import CountdownTimer from "../general/reusable/countdown/countdown";
import MediaCard from "../general/reusable/card/simpleCard";
import VerticalLinearStepper from "./goalsStepper";

const IndexStartView = () => {
        const listOfCards: any = [1,2,3,4,5,6,7,8]
        return (
                <div>
                        <Grid container className={"container-1"}>
                                <Grid item xs={12} sm={12}>
                                        <div className={"mid-view"}>
                                                <h2 className="coming-soon-view">big heading</h2>
                                        </div>
                                </Grid>
                        </Grid>
                        <Grid container className={"container-2"}>
                                <Grid item xs={2} sm={2}>
                                        <div className={"left-panel"}>
                                        </div>
                                </Grid>
                                <Grid item xs={8} sm={8}>
                                        <div className={"mid-view-2"}>
                                                {listOfCards.map((x:any, index: number) => {
                                                        // return (<MediaCard className={"card"} key={index} searchItemID={index.toLocaleString()} searchItemTitle={"hey"} searchItemURL={""} imageURL={""} />)
                                                        return (
                                                        <div className={"card"} key={index}>
                                                                <MediaCard className={"cardItem"} key={index} searchItemID={index.toLocaleString()} searchItemTitle={"hey"} searchItemURL={"https://smalltotall.info/wp-content/uploads/2017/04/google-favicon-vector-400x400.png"} imageURL={"https://smalltotall.info/wp-content/uploads/2017/04/google-favicon-vector-400x400.png"} />
                                                        </div>)
                                                })
                                                }
                                        </div>
                                </Grid>
                                <Grid item xs={2} sm={2}>
                                        <div className={"right-panel"}>
                                        </div>
                                </Grid>
                        </Grid>
                                                <VerticalLinearStepper />       
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
                        {2 < 1? <AppLaunchView/>: <IndexStartView/>}
                </div>
        );
}

export default Start;
