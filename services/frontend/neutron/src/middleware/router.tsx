import * as React from "react";
// import { BrowserRouter } from "react-router-dom";
import { BrowserRouter as Router, Route, Redirect, Switch } from "react-router-dom";
import Home from "../components/home/home";
// import PrimarySearchAppBar from "../components/navbar/navbar";
import Profile from "../components/profile/profile";
import About from "../components/about/about";
import Feed from "../components/feed/feed";
import Dashboard from "../components/dashboard/dashboard";
import Workspace from "../components/workspace/workspace";
import Login from "../components/login/login";
import Register from "../components/register/register";
import Discover from "../components/discover/discover";
import PrimarySearchAppBar from "../components/navbar/navbar";
// import "./style.css";
import Leaderboard from "../components/leaderboard/leaderboard";


export default function({ children }: { children: React.ReactChild }) {
	return (
		<Router>
			<>
			{/* <ButtonAppBar /> */}
				<header>
					<PrimarySearchAppBar />
				</header>
				<main className={"content"}>
					<Switch>
						<Route exact path="/" component={Home} />
						<Route exact path="/login" component={Login} />
						<Route exact path="/register" component={Register} />
						<Route exact path="/about" component={About} />
						<Route exact path="/feed" component={Feed} />
						<Route exact path="/dashboard" component={Dashboard} />
						<Route exact path="/workspace" component={Workspace} />
						<Route exact path="/profile" component={Profile} />
						<Route exact path="/leaderboard" component={Leaderboard} />
						<Route exact path="/discover" component={Discover} />

					</Switch>
				</main>
			</>
		</Router>
	);
}
