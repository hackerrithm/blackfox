import * as React from "react";
import { useMutation } from "@apollo/react-hooks";
import gql from "graphql-tag";
import { useState } from "react";
import RegisterForm from "./regsiterForm";

const SAVE_USER = gql`
	mutation registerUser($user: RegisterInput!) {
		registerUser(user: $user)
	}
`;

interface User {
	username: string;
	password: string;
	firstname: string;
	lastname: string;
	emailAddress: string;
	gender: string;
}

interface RegisterDetails {
	username: string;
	password: string;
	firstname: string;
	lastname: string;
	emailAddress: string;
	gender: string;
}

export default function Register() {
	const [username, setUsername] = useState<any>("");
	const [firstname, setFirstname] = useState<any>("");
	const [lastname, setLastname] = useState<any>("");
	const [emailAddress, setEmailAddress] = useState<any>("");
	const [password, setPassword] = useState<any>("");
	const [gender, setGender] = useState<any>("");

	const [registerUser, { error, data }] = useMutation<
		{ registerUser: User },
		{ user: RegisterDetails }
	>(SAVE_USER, {
		variables: {
			user: {
				username,
				password,
				firstname,
				lastname,
				emailAddress,
				gender
			}
		}
	});

	return (
		<div className="base-container-register">
			<RegisterForm
				username={username}
				password={password}
				emailaddress={emailAddress}
				gender={gender}
				error={error}
				registerUser={registerUser}
				data={data}
				setUsername={setUsername}
				setEmailAddress={setEmailAddress}
				setPassword={setPassword}
			/>
		</div>
	);
	// }
}
