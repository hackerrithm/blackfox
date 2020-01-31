// import React, { useReducer, useState } from "react";
// import { State, Action } from ".";

const SearchInitialState: any = {
		isLoading: false,
		isError: false,
		data: [],
};

// const TaskReducer: React.Reducer<State, Action> = (state: any, action: any) => {
// 	switch (action.type) {
// 		case "reset": {
// 			return action.payload;
// 		}
// 		case "add": {
// 			const newTask = {
// 				id: Date.now(),
// 				text: action.text,
// 				completed: false
// 			};
// 			return {
// 				task: [...state.task, newTask]
// 			};
// 		}
// 		case "edit": {
// 			const idx = state.task.findIndex((t: any) => t.id === action.id);
// 			const todo = Object.assign({}, state.task[idx]);
// 			todo.text = action.text;
// 			const todos = Object.assign([], state.task);
// 			todos.splice(idx, 1, todo);
// 			return {
// 				task: todos
// 			};
// 		}
// 		case "delete": {
// 			const idx = state.task.findIndex((t: any) => t.id === action.id);
// 			const todos = Object.assign([], state.task);
// 			todos.splice(idx, 1);
// 			return {
// 				task: todos
// 			};
// 		}
// 		// TODO: Needs work
// 		case "completed": {
// 			const idx = state.task.findIndex((t: any) => t.id === action.id);
// 			const todo = Object.assign({}, state.task[idx]);

// 			state.task.map((item: any) => {
// 				if (item.id === action.id) {
// 					todo.completed = !item.completed
// 					return {
// 						task: todo
// 					};
// 				}
// 				return item;
// 			});
// 		}
// 		default:
// 			return state;
// 	}
// };


const DataFetchReducer = (state: any, action: any) => {
	switch (action.type) {
		case "FETCH_INIT":
			return { ...state, isLoading: true, isError: false };
		case "FETCH_SUCCESS":
			return {
				...state,
				isLoading: false,
				isError: false,
				data: action.payload,
			};
		case "FETCH_FAILURE":
			return {
				...state,
				isLoading: false,
				isError: true
			};
		default:
			throw new Error();
	}
};

export { DataFetchReducer, SearchInitialState };
