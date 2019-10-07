import { render } from "react-dom";
import App from "./app";
import * as React from "react";
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";

const client2 = new ApolloClient({
        uri: "http://localhost:9000/query",
        // credentials: "include"
      });
// return (
//         <ApolloProvider client={client2}>
//                 {this.props.children}
//         </ApolloProvider>
// );


render(
        <ApolloProvider client={client2}>
                <App />
        </ApolloProvider>,
        document.getElementById("root"));
