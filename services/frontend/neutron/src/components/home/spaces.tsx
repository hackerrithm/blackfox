import React, { useState, useEffect } from "react";
// import List from '@material-ui/core/List';
// import ListItem from '@material-ui/core/ListItem';
// import ListItemText from '@material-ui/core/ListItemText';
// import Divider from '@material-ui/core/Divider';
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import { List, ListItem, ListItemAvatar, Avatar, ListItemText, Typography } from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			width: '100%',
			maxWidth: 360,
			backgroundColor: theme.palette.background.paper,
		},
		inline: {
			display: 'inline',
		},
	}),
);

const Spaces = (props: any) => {
	const classes = useStyles("");
	const [isFetching, setIsFetching] = useState(true);

	useEffect(() => {
		if (!isFetching) return;
		getData(isFetching);
		setIsFetching(true);
	}, [isFetching]);

	const getData = async (load: any) => {
		if (load) {
			await fetch('https://jsonplaceholder.typicode.com/users')
				.then(res => {
					return !res.ok
						? res.json().then(e => Promise.reject(e))
						: res.json();
				})
				.then(res => {
					props.setState([...props.state, ...res]);
				});
		}
	};

	return (
		<>
				<List className={classes.root}>
			{props.state.map((listItem: any, index: number) => {
				return (
					<ListItem key={index} alignItems="flex-start">
						<ListItemAvatar>
							<Avatar alt="Remy Sharp" src="/static/images/avatar/1.jpg" />
						</ListItemAvatar>
						<ListItemText
							primary={listItem.name}
							secondary={
								<React.Fragment>
									<Typography
										component="span"
										variant="h5"
										className={classes.inline}
										color="textPrimary"
									>
									</Typography>
								</React.Fragment>
							}
						/>
					</ListItem>
				)
			})
		}
		</List>
		</>
	);
};

export default Spaces;
