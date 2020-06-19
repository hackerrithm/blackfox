import * as React from "react";
// import { BrowserRouter } from "react-router-dom";
import { BrowserRouter as Router, Route, Redirect, Switch } from "react-router-dom";
import Start from "../start/start";
import CustomNavbar from "../navbar/navbar";
import Login from "../login/login";
import Profile from "../profile/profile";
import PrivateRoute from "./privateRoute";
import Register from "../register/register";
import Home from "../home/home";
import Dashboard from "../dashboard/dashboard";
import Workspace from "../workspace/workspace";
import Leaderboard from "../leaderboard/leaderboard";
import Discover from "../discover/discover";
import About from "../about/about";


export default function({ children }: { children: React.ReactChild }) {
	return (
		<Router>
			<>
				<header>
					<CustomNavbar />
				</header>
				<main className={"content"}>
					<Switch>
						<Route exact path="/" component={Start} />
						<Route exact path="/about" component={About} />
						<Route exact path="/register" component={Register} />
						<Route exact path="/login" component={Login} />
						<PrivateRoute exact path="/profile" component={Profile} />
						<PrivateRoute exact path="/home" component={Home} />
						<PrivateRoute exact path="/dashboard" component={Dashboard} />
						<PrivateRoute exact path="/workspace" component={Workspace} />
						<PrivateRoute exact path="/leaderboard" component={Leaderboard} />
						<PrivateRoute exact path="/discover" component={Discover} />
					</Switch>
				</main>
			</>
		</Router>
	);
}
