import gql from "graphql-tag";

export const LOGIN_USER = gql`
	mutation loginUser($user: LoginInput!) {
		loginUser(user: $user)
	}
`;

export interface User {
	username: string;
	password: string;
}

export interface LoginDetails {
	username: string;
	password: string;
}