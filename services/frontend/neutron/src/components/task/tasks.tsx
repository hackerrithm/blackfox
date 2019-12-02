import React, { useReducer, useEffect, useRef } from "react";
import AddTask from "./addTask";
import { State, Action } from ".";
import { TaskReducer, InitialState } from "./reducer";
import Task from "./task";

export const Context = React.createContext(InitialState);

const useEffectOnce = (cb: any) => {
	const didRun = useRef(false);

	useEffect(() => {
		if (!didRun.current) {
			cb();
			didRun.current = true;
		}
	});
};

const Tasks = () => {
	const [state, dispatch] = useReducer<React.Reducer<State, Action>>(
		TaskReducer,
		InitialState
	);

	useEffectOnce(() => {
		const raw = localStorage.getItem("data");
		dispatch({ type: "reset", payload: JSON.parse(raw) });
	});

	useEffect(() => {
		localStorage.setItem("data", JSON.stringify(state));
	}, [state]);

	return (
		<>
			<Context.Provider value={dispatch as any}>
				<AddTask
					add={(text: any) => dispatch({ type: "add", text: text })}
				/>
				<ul>
					{[] && (state.task!.map((item: any) => {
						return (
							<li key={item.id}>
								<Task
									edit={item.id}
									checked={item.completed}
									todo={item.text}
									remove={() => dispatch({type: "delete", id: item.id, payload: item.id})}
									key={item.id}
								/>
							</li>
						);
					}))}
				</ul>
			</Context.Provider>
		</>
	);
};

export default Tasks;
