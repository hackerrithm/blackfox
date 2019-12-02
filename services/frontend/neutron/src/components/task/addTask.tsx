import React, { useState } from "react";

const AddTask = ({add}: any) => {
	const [text, setText] = useState("");
	return (
		<div>
			<div className="AddTodo">
				<input
					value={text}
					onChange={e => setText(e.target.value)}
					className="AddTodoInput"
				/>
				<button
					className="AddTodoButton"
					onClick={() => {
                        add(text);
                        setText("");
					}}
				>
					Add
				</button>
			</div>
		</div>
	);
};

export default AddTask;
