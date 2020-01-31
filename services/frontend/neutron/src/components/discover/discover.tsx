import * as React from "react";
import {
	Grid,
	Paper,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core";
import { getThemeProps } from "@material-ui/styles";
import Books from "../book/books";

const Discover = () => {
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
	return (
		<div className={classes.root}>
			<Grid container spacing={3}>
				<Grid item xs={12}>
					<Paper className={classes.paper}>
						<Books />
					</Paper>
				</Grid>
				<Grid item xs={12} sm={6}></Grid>
				<Grid item xs={12} sm={6}></Grid>
			</Grid>
		</div>
	);
}

export default Discover;