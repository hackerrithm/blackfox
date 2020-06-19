import React, { useReducer, useState } from "react";
import { ApolloProvider } from "react-apollo";
import { ApolloClient } from "apollo-client";
import { createHttpLink } from "apollo-link-http";
import { setContext } from "apollo-link-context";
import { InMemoryCache } from "apollo-cache-inmemory";
import RouteManager from "./components/router/routeManager";
import { InitialState, Reducer } from "./components/login/reducer";
import * as Actions from "./components/login/actions";
import { AuthContext } from "./components/context/authContext";

const App = (props: any) => {
	const httpLink = createHttpLink({
		uri: "http://localhost:9000/query",
		credentials: "same-origin"
	});

	const authLink = setContext((_: any, { headers }: any) => {
		// get the authentication token from local storage if it exists
		const token = localStorage.getItem("token");
		// return the headers to the context so httpLink can read them
		return {
			headers: {
				...headers,
				authorization: token ? `Bearer ${token}` : ""
			}
		};
	});

	const client = new ApolloClient({
		link: authLink.concat(httpLink),
		cache: new InMemoryCache()
	});

	const [state, dispatch] = useReducer(Reducer, InitialState);

	React.useEffect(() => {
		// const user = JSON.parse(localStorage.getItem("user") || null);
		const token = JSON.parse(localStorage.getItem("token") || null);

		if (token) {
			dispatch({
				type: Actions.loginSuccess(),
				payload: {
					// user,
					token
				}
			});
		}
	}, []);

	return (
		<AuthContext.Provider
			value={{
				state,
				dispatch
			}}
		>
			<ApolloProvider client={client}>
					<RouteManager children={props.children} />
			</ApolloProvider>
		</AuthContext.Provider>
	);
};

export default App;
