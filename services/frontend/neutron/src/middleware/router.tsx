import * as React from "react";
// import { BrowserRouter } from "react-router-dom";
import { BrowserRouter as Router, Route, Redirect, Switch } from "react-router-dom";
import Home from "../components/home/home";
import PrimarySearchAppBar from "../components/navbar/navbar";
import Profile from "../components/profile/profile";
import About from "../components/about/about";
import Feed from "../components/feed/feed";
import Dashboard from "../components/dashboard/dashboard";
import Workspace from "../components/workspace/workspace";
import Login from "../components/login/login";
import Register from "../components/register/register";

export default function({ children }: { children: React.ReactChild }) {
	return (
		<Router>
			<div>
				<header>
					<PrimarySearchAppBar />
				</header>
				<main>
					<Switch>
						<Route exact path="/" component={Home} />
						<Route exact path="/login" component={Login} />
						<Route exact path="/register" component={Register} />
						<Route exact path="/about" component={About} />
						<Route exact path="/feed" component={Feed} />
						<Route exact path="/dashboard" component={Dashboard} />
						<Route exact path="/workspace" component={Workspace} />
						<Route exact path="/profile" component={Profile} />
					</Switch>
				</main>
			</div>
		</Router>
	);
}
