import React, { useState, useEffect } from "react";
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import ListSubheader from '@material-ui/core/ListSubheader';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Collapse from '@material-ui/core/Collapse';
import InboxIcon from '@material-ui/icons/MoveToInbox';
import DraftsIcon from '@material-ui/icons/Drafts';
import SendIcon from '@material-ui/icons/Send';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import StarBorder from '@material-ui/icons/StarBorder';

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			width: '100%',
			maxWidth: 360,
			backgroundColor: theme.palette.background.paper,
		},
		nested: {
			paddingLeft: theme.spacing(4),
		},
	}),
);

const Trending = (props: any) => {
	const classes = useStyles("");
	const [isFetching, setIsFetching] = useState(true);
	const [open, setOpen] = React.useState(false);

	const handleClick = () => {
		setOpen(!open);
	};

	useEffect(() => {
		console.log("started now: ")
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
					console.log("week: ", res);
					props.setState([...props.state, ...res]);
				});
		}
	};

	return (
		<>
			<List
				component="nav"
				aria-labelledby="nested-list-subheader"
				subheader={
					<ListSubheader component="div" id="nested-list-subheader">
						Nested List Items
  					</ListSubheader>
				}
				className={classes.root}
			>

				{
					props.state.map((listItem: any, index: number) => {
						return (
						<>
							<ListItem button onClick={handleClick}>
								<ListItemIcon>
									<InboxIcon />
								</ListItemIcon>
								<ListItemText primary={listItem.website} />
								{open ? <ExpandLess /> : <ExpandMore />}
							</ListItem>
							<Collapse in={open} timeout="auto" unmountOnExit>
								<List component="div" disablePadding>
									<ListItem button className={classes.nested}>
										<ListItemIcon>
											<StarBorder />
										</ListItemIcon>
										<ListItemText primary="Starred" />
									</ListItem>
								</List>
							</Collapse>
						</>
						)
					}
					)}

			</List>
		</>
	);
};

export default Trending;
