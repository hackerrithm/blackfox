import React, { useState, useEffect } from "react";
import Box from "@material-ui/core/Box";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import FavoriteIcon from '@material-ui/icons/Favorite';
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import { List, ListItem, ListItemAvatar, Avatar, ListItemText, Typography } from "@material-ui/core";
import useStyles from "./spacesStyles";

interface TabPanelProps {
	children?: React.ReactNode;
	index: any;
	value: any;
}

function TabPanel(props: TabPanelProps) {
	const { children, value, index, ...other } = props;

	return (
		<Typography
			component="div"
			role="tabpanel"
			hidden={value !== index}
			id={`vertical-tabpanel-${index}`}
			aria-labelledby={`vertical-tab-${index}`}
			{...other}
		>
			{value === index && <Box p={3}>{children}</Box>}
		</Typography>
	);
}

function a11yProps(index: any) {
	return {
		id: `vertical-tab-${index}`,
		'aria-controls': `vertical-tabpanel-${index}`,
	};
}


const Spaces = (props: any) => {
	const classes = useStyles("");
	const [isFetching, setIsFetching] = useState(true);
	const [tabValue, setTabValue] = React.useState(0);

	const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
		setTabValue(newValue);
	};

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
			<div className={classes.spacesContainer}>

				<Tabs
					orientation="vertical"
					variant="scrollable"
					value={tabValue}
					onChange={handleChange}
					aria-label="Vertical tabs example"
					className={classes.tabs}
				>
					{props.state.map((listItem: any, index: number) => {
						return (

							<Tab label={listItem.name} {...a11yProps(index)} />
						)
					})
					}
				</Tabs>
				{props.state.map((listItem: any, index: number) => {
					return (
						<TabPanel value={tabValue} index={index}>
							{listItem.name}
						</TabPanel>
					)
				})
				}
			</div>
			{/* <List className={classes.root}>
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
			</List> */}
		</>
	);
};

export default Spaces;
