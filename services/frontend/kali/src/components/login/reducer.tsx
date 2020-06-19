import * as ActionType from "./actionTypes";

export const InitialState: any = {
	isAuthenticated: false,
	user: null,
	token: null,
	isLoggedIn: false
};

export const Reducer = (state: any, action: any) => {
	switch (action.type) {
		case ActionType.LOGIN_SUCCESS:
			return {
				...state,
				isAuthenticated: true,
				user: action.payload.user,
				token: action.payload.token
			};
		case ActionType.IS_LOGGED_IN:
			return {
				...state,
				isLoggedIn: true,
			};
		case ActionType.LOGOUT:
			return {
				...state,
				isAuthenticated: false,
				isLoggedIn: false,
				user: null
			};
		default:
			return state;
	}
};
