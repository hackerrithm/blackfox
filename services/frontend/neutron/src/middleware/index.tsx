// import { GRAPHQL_ENDPOINT } from "_config/index";
import * as React from "react";
import Apollo from "./apollo";
import Router from "./router";
import { DataFetchReducer, SearchInitialState } from "../components/book";

export const AppContext = React.createContext(SearchInitialState);

export default function Middlwares({
	children
}: {
	children: React.ReactChild;
}) {
	const [query, setQuery] = React.useReducer(DataFetchReducer, SearchInitialState);
	// ROUTER must come before ALL OTHER middleware
	return (
		<Router>
			<Apollo graphqlURL={"http://localhost:9000/query"}>
			<AppContext.Provider value={{ query, setQuery }}>
				{children}
				</AppContext.Provider>
			</Apollo>
		</Router>
	);
}
