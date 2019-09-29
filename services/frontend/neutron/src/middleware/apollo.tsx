import * as React from "react";
import { withClientState } from "@tgrx/apollo-link-state";
import { InMemoryCache } from "apollo-cache-inmemory";
import { ApolloClient } from "apollo-client";
import { ApolloLink } from "apollo-link";
import { onError } from "apollo-link-error";
import { HttpLink } from "apollo-link-http";
import { ApolloProvider } from "react-apollo";

import { ResolverDefaults,  Resolvers } from "../gql";

interface IProps {
        graphqlURL: string;
        children: React.ReactChild;
}

export default class ApolloClientProvider extends React.Component<IProps> {

        // tslint:disable-next-line:variable-name
        private _apolloClient: ApolloClient<any>;

        // tslint:disable-next-line:variable-name
        private _cache = new InMemoryCache();

        // tslint:disable-next-line:variable-name
        private _stateLink = withClientState({
                cache: this._cache,
                defaults: ResolverDefaults,
                resolvers: Resolvers,
        });

        // tslint:disable-next-line:variable-name
        private _httpLink = new HttpLink({
                uri: this.props.graphqlURL,
        });

        // tslint:disable-next-line:variable-name
        private _errorLink = onError(({ graphQLErrors, networkError }) => {
                if (graphQLErrors) {
                  graphQLErrors.map(({ message, locations, path }) => {
                    // tslint:disable-next-line:no-console
                    console.error(`[GraphQL] Message: ${message}, Location: ${locations}, Path: ${path}`)
                  });
                }
                if (networkError) {
                  // tslint:disable-next-line:no-console
                  console.error(`[Network] ${networkError}`);
                }
        });

        constructor(props: IProps) {
                super(props);
                // tslint:disable-next-line:member-ordering
                const links = [
                        this._stateLink,
                        this._errorLink,
                        this._httpLink,
                ];

                this._apolloClient = new ApolloClient({
                        cache: this._cache,
                        connectToDevTools: true,
                        link: ApolloLink.from(links),
                });
        }

        public render() {
                return (
                        <ApolloProvider client={this._apolloClient}>
                                {this.props.children}
                        </ApolloProvider>
                );
        }

}
