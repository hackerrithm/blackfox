import React, { useReducer } from "react";

const initialTodos = [
	{
		id: "a",
		task: "Learn React",
		complete: false
	},
	{
		id: "b",
		task: "Learn Firebase",
		complete: false
	}
];

const todoReducer = (state: any, action: any) => {
	switch (action.type) {
        case "ADD": {
            const newTask = {
				id: Date.now(),
				task: action.text,
				completed: false
			};
			return {
				// ...state.task,
				// {
				// {
				// 	id: Date.now(),
				// 	text: "",
				// 	completed: false
				// }
				// task: "bless"
				task: [...state.task, newTask]
			};
        }
            return 
		case "DO_TODO":
			return state.map((todo: any) => {
				if (todo.id === action.id) {
					return { ...todo, complete: true };
				} else {
					return todo;
				}
			});
		case "UNDO_TODO":
			return state.map((todo: any) => {
				if (todo.id === action.id) {
					return { ...todo, complete: false };
				} else {
					return todo;
				}
			});
		default:
			return state;
	}
};

const Todo = () => {
	const [todos, dispatch] = React.useReducer(todoReducer, initialTodos);

	const handleChange = (todo: any) => {
		dispatch({
			type: todo.complete ? "UNDO_TODO" : "DO_TODO",
			id: todo.id
		});
	};

	return (
		<>
			<button
				className="AddTodoButton"
				onClick={() => {
					console.log('being clicked');
                    
					dispatch({ type: "add", text: "random shit" });
				}}
			>
				Add todo
			</button>
			<ul>
				{todos.map((todo: any) => (
					<li key={todo.id}>
						<label>
							<input
								type="checkbox"
								checked={todo.complete}
								onChange={() => handleChange(todo)}
							/>
							{todo.task}
						</label>
					</li>
				))}
			</ul>
		</>
	);
};

export default Todo;
