import { render } from "react-dom";
import App from "./app";
import * as React from "react";
// import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";
import { ApolloClient } from 'apollo-client';
import { createHttpLink } from 'apollo-link-http';
import { setContext } from 'apollo-link-context';
import { InMemoryCache } from 'apollo-cache-inmemory';

const httpLink = createHttpLink({
        uri: "http://localhost:9000/query",
        credentials: 'same-origin'
});

const authLink = setContext((_:any, { headers }: any) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem('token');
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
});

const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache()
});

// const client2 = new ApolloClient({
//         uri: "http://localhost:9000/query",
//         // credentials: "include"
//       });
// return (
//         <ApolloProvider client={client2}>
//                 {this.props.children}
//         </ApolloProvider>
// );


render(
        <ApolloProvider client={client}>
                <App />
        </ApolloProvider>,
        document.getElementById("root"));
