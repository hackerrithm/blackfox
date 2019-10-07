import * as React from "react";
// import { BrowserRouter } from "react-router-dom";
import {
        BrowserRouter as Router,
        Route,
        Switch,
        Link,
      } from "react-router-dom";
import Home from "../components/home/home";
import PrimarySearchAppBar from "../components/navbar/navbar";

export default function({children}: {children: React.ReactChild}) {
        return (
                <Router>
                        <div>
                                <header>
                                        {/* <Link to="/about">About</Link>
                                        {" "}
                                        <Link to="/">Login</Link>
                                        {" "}
                                        <strong>neutron</strong> */}
                                        <PrimarySearchAppBar />
                                </header>
                                <main>

                                        <Switch>
                                                <Route exact path="/" component={Home} />
                                        </Switch>
                                </main>
                        </div>
                </Router>
        );
}
