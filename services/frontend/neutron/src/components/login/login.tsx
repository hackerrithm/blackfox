import React, { useReducer, useEffect } from "react";
// import Button from "@material-ui/core/Button";
import { useHistory } from "react-router-dom";
import { useMutation } from "@apollo/react-hooks";
import gql from "graphql-tag";
import { login, LoginReducer, InitialState } from ".";
import Profile from "../profile/profile";
import MButton from "../general/reusable/button/mbutton";
import { authContext } from "../utils/authContext";
import { apiRequest, validateLoginForm } from "../utils/Helpers";

export const Context = React.createContext(InitialState);

const LOGIN_USER = gql`
	mutation loginUser($user: LoginInput!) {
		loginUser(user: $user)
	}
`;

interface User {
	username: string;
	password: string;
}

interface LoginDetails {
	username: string;
	password: string;
}

const Login: React.FC = () => {
	const auth = React.useContext(authContext);
	const history = useHistory();

	const [state, dispatch] = useReducer(LoginReducer, InitialState);
	const { username, password, isLoading, error, isLoggedIn } = state;
	const [loading, setLoading] = React.useState(false);

	const onSubmit = async (e: any) => {
		console.log("login called");
		e.preventDefault();

		dispatch({ type: "login" });

		try {
			await login({ username, password });
			dispatch({ type: "success" });
			// localStorage.setItem('isLoggedIn', context.accessToken)
			history.push("/profile");
		} catch (error) {
			dispatch({ type: "error" });
		}
	};

	useEffect(() => {
		localStorage.setItem("userAuthData", JSON.stringify(state));
	}, [state]);

	const authHandler = async () => {
		try {
			setLoading(true);
			const userData = await apiRequest(
				"https://jsonplaceholder.typicode.com/users",
				// "http://www.mocky.io/v2/5ded90a83300006d002b9178",
				"post",
				{ username: username, password: password }
			);
			const { id, user } = userData;
			auth.setAuthStatus({ id, user });
			dispatch({ type: "success" });
			history.push("/profile");
		} catch (err) {
			setLoading(false);
			//   showError(err.message);
			console.log("error my g");
		}
	};

	/**
	 * Render Login
	 */
	// public render() {
	// const [username, setUsername] = useState<any>("");
	// const [password, setPassword] = useState<any>("");

	// const [loginUserSubmit, { error, data }] = useMutation<
	// 	{ loginUser: User },
	// 	{ user: LoginDetails }
	// >(LOGIN_USER, {
	// 	variables: { user: { username, password } }
	// });
	return (
		<div className="base-container">
			<Context.Provider
				value={{
					username: state.username,
					// password: state.password,
					isLoading: state.isLoading,
					isLoggedIn: state.isLoggedIn,
					error: state.error
				}}
			>
				<div className="header">Login</div>
				{/* {error ? <p>Oh no! {error.message}</p> : null}
			{data && loginUserSubmit ? <p>Login!</p> : null} */}
				{/* {isLoggedIn ? (
				<>
					<Profile username={username} />
				</>
			) :  */}
				<form
					className="form"
					onSubmit={(e: any) => {
						e.preventDefault();
						if (validateLoginForm(username, password)) {
							authHandler();
						}
					}}
				>
					<div className="form-group">
						<label htmlFor="username">Username</label>
						<input
							type="text"
							name="username"
							placeholder="username"
							value={username}
							onChange={(e: any) => {
								dispatch({
									type: "field",
									fieldName: "username",
									payload: e.currentTarget.value
								});
							}}
						/>
					</div>
					<div className="form-group">
						<label htmlFor="password">Password</label>
						<input
							type="text"
							name="password"
							placeholder="password"
							value={password}
							onChange={(e: any) => {
								console.log("changed");

								dispatch({
									type: "field",
									fieldName: "password",
									payload: e.currentTarget.value
								});
							}}
						/>
					</div>
					<MButton
						className="submit"
						type="submit"
						disabled={isLoading}
						color="primary"
						variant="contained"
						label={isLoading ? "Logging in..." : "Log In"}
					/>
				</form>
				{/* // } */}
			</Context.Provider>
		</div>
	);
};

export default Login;
