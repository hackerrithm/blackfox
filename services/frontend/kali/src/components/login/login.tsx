import React from "react";
import { useHistory } from "react-router-dom";
import { useMutation } from "@apollo/react-hooks";
import { Button } from "@material-ui/core";
import * as ActionType from "./actionTypes";
import { LOGIN_USER, User } from "./types";
import { AuthContext } from "../context/authContext";
import LoginForm from "./loginForm";

const Login: React.FC = () => {
	const { dispatch } = React.useContext(AuthContext);
	const history = useHistory();
	const initialState: any = {
		username: "",
		password: "",
		isSubmitting: false,
		errorMessage: null
	};

	const [statedata, setData] = React.useState(initialState);

	const handleInputChange = (event: any) => {
		setData({
			...statedata,
			[event.target.name]: event.target.value
		});
	};

	const { username, password } = statedata;
	const [handleFormSubmit, { error, data }] = useMutation<
		{ loginUser: User },
		{ user: { username: string; password: string } }
	>(LOGIN_USER, {
		update: (proxy, mutationResult) => {
			console.log("login called:: ", proxy, " and ", mutationResult);

			setData({
				...statedata,
				isSubmitting: true,
				errorMessage: null
			});

			try {
				dispatch({
					type: ActionType.LOGIN_SUCCESS,
					payload: mutationResult.data
				});
				console.log("success");
				localStorage.setItem(
					"token",
					JSON.stringify(mutationResult.data.loginUser)
				);

				dispatch({
					type: ActionType.IS_LOGGED_IN,
					payload: mutationResult.data
				});
				localStorage.setItem("isLoggedIn", "true")
				history.push("/home");
			} catch (error) {
				console.log("error");
			}
		},
		variables: { user: { username, password } }
	});

	return (
		<div>
			{/* <p>Login</p>
			<form>
				<h1>Login</h1>

				<label htmlFor="username">
					Username
					<input
						type="text"
						value={statedata.username}
						onChange={handleInputChange}
						name="username"
						id="username"
					/>
				</label>

				<label htmlFor="password">
					Password
					<input
						type="password"
						value={statedata.password}
						onChange={handleInputChange}
						name="password"
						id="password"
					/>
				</label>

				{statedata.errorMessage && (
					<span className="form-error">{statedata.errorMessage}</span>
				)}

				<Button
					className={"bit-button"}
					disabled={statedata.isSubmitting}
					onClick={() => handleFormSubmit()}
				>
					{statedata.isSubmitting ? (
						<img
							className="spinner"
							src={"logo"}
							alt="loading icon"
						/>
					) : (
						"Login"
					)}
				</Button>
			</form> */}
			<LoginForm
				loginUserSubmit={handleFormSubmit}
				data={data}
				error={error}
				formError={statedata.errorMessage}
				handleInputChange={handleInputChange}
				username={statedata.username}
				password={statedata.password}
				isLoading={statedata.isSubmitting}
				dispatch={dispatch}
			/>

		</div>
	);
};

export default Login;
