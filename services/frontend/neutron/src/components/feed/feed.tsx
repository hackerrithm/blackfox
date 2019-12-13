import React, { useReducer } from "react";
import { LoginReducer, InitialState } from "../login/reducer";
import { useHistory } from "react-router";
import List from "../home/list";

const Feed = () => {
    const [state, dispatch] = useReducer(LoginReducer, InitialState);
    const history = useHistory();
    let listSize: number = 3;
	return (
		<div>
			{/* <h1>Welcome {username}!</h1> */}
            <br/>
            <br/>
            <br/>
            <br/>
            <br/>
            <br/>

            <List initialListCount={listSize} />
		</div>
	);
};

export default Feed;
