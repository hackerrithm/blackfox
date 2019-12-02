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

type Actions = "add" | "reset" | "delete" | "completed" | "edit";

export interface Action {
	id?: any;
	type: Actions;
	text?: any;
	payload?: any;
	completed?: any
}