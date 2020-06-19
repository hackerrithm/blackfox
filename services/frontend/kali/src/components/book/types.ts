export interface State {
	task?: [
		{
			id: number;
			text: string;
			completed: boolean;
		}
	];
	dispatch?: any;
}

type Actions = "FETCH_INIT" | "FETCH_SUCCESS" | "FETCH_FAILURE";

export interface Action {
	id?: any;
	type: Actions;
	text?: any;
	payload?: any;
	completed?: any
}