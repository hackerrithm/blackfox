import React, { useState } from "react";
import {
	Grid,
	Paper,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core";
import { getThemeProps } from "@material-ui/styles";
import TopTrendingUsers from "./userList";

const Leaderboard = () => {
	return (
		<div className="discover">
			<RepositoriesSection />
		</div>
	);
}

const useStylesRepositoriesSection = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			flexGrow: 1
		},
		paper: {
			padding: theme.spacing(2),
			textAlign: "center",
			color: theme.palette.text.secondary
		}
	})
);

function RepositoriesSection() {
	const classes = useStylesRepositoriesSection(getThemeProps);
	const [state, setState] = useState([])
	return (
		<div className={classes.root}>
			<Grid container spacing={3}>
				<Grid item xs={12}>
					<Paper className={classes.paper}>
						<TopTrendingUsers state={state} setState={setState} />
					</Paper>
				</Grid>
				<Grid item xs={12} sm={6}></Grid>
				<Grid item xs={12} sm={6}></Grid>
			</Grid>
		</div>
	);
}

export default Leaderboard;
