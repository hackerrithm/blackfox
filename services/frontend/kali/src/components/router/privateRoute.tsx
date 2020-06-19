import React from "react";
import { Route, Redirect } from "react-router-dom";

const PrivateRoute = ({ component: Component, ...rest }: any) => {
	return (
		<Route
			{...rest}
			render={props =>
				(localStorage.getItem("isLoggedIn") === "true") ? <Component {...props} /> : <Redirect to="/login" />
			}
		/>
	);
}

export default PrivateRoute;
