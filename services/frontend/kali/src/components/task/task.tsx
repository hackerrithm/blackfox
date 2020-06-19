import React, { useContext } from "react";
import {
	Paper,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core";
import { getThemeProps } from "@material-ui/styles";
import { Context } from "./tasks";

const useStylesTask = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			flexGrow: 1
		},
		paper: {
			padding: theme.spacing(2),
			textAlign: "left",
			color: theme.palette.text.secondary
		}
	})
);

const Task = ({ todo, id, completed, edit, remove }: any) => {
    const dispatch = useContext<any>(Context)
	const classes = useStylesTask(getThemeProps);
    const handleChange = (e: any) => {
        console.log('e value: ', e.target.checked);
		dispatch({
			type: "completed",
            id: edit,
            completed: e.target.checked
		});
    };

    const handleChangedText = (e: any) => {
        console.log('e value: ', e.target.value);
        
		dispatch({
			type: "edit",
            id: edit,
            payload: edit,
            text: e.target.value,
		});
    };

	return (
        <div>
            <Paper className={classes.paper}>
                <input type="checkbox" onChange={ (e) => {handleChange(e)}} checked={completed}/>
                <input key={id}  defaultValue={todo} onChange={(e) => handleChangedText(e)}/>
                <button onClick={remove}>delete</button>
            </Paper>
        </div>
	);
};

export default Task;
