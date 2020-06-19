import * as ActionTypes from "./actionTypes"

export const loginSuccess = () => {
	return {
		type: ActionTypes.LOGIN_SUCCESS
	};
};

export const loginFailure = () => {
	return {
		type: ActionTypes.LOGIN_FAILURE
	};
};

export const logout = () => {
	return {
		type: ActionTypes.LOGOUT
	};
};

export const isLoggedIn = () => {
	return {
		type: ActionTypes.IS_LOGGED_IN
	};
};
