import React from "react";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import InputBase from "@material-ui/core/InputBase";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import DirectionsIcon from "@material-ui/icons/Directions";

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			padding: "2px 4px",
			display: "flex",
			alignItems: "center",
			width: 400
		},
		input: {
			marginLeft: theme.spacing(1),
			flex: 1
		},
		iconButton: {
			padding: 10
		},
		divider: {
			height: 28,
			margin: 4
		}
	})
);

export default function CustomizedInputBase({val, onChange, btnType, onSubmit, placeholder}: any) {
	const classes = useStyles("");

	return (
		<Paper component="form" onSubmit={onSubmit} className={classes.root}>
			{/* <IconButton className={classes.iconButton} aria-label="menu">
				<MenuIcon />
			</IconButton> */}
			<InputBase
                type="text"
                value={val}
				className={classes.input}
				placeholder={placeholder === undefined? "Search Blackfox": placeholder}
                inputProps={{ "aria-label": "search blackfox" }}
                onChange={onChange}
			/>
			<IconButton
				type={btnType}
				className={classes.iconButton}
				aria-label="search"
			>
				<SearchIcon />
			</IconButton>
			{/* <Divider className={classes.divider} orientation="vertical" /> */}
			{/* <IconButton
				color="primary"
				className={classes.iconButton}
				aria-label="directions"
			>
				<DirectionsIcon />
			</IconButton> */}
		</Paper>
	);
}
