import {
	fade,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core/styles";

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		body: {
			margin: theme.spacing(-10)
		},
		grow: {
			flexGrow: 1
		},
		menuButton: {
			marginRight: theme.spacing(2)
		},
		title: {
			display: "none",
			[theme.breakpoints.up("sm")]: {
				display: "block"
			},
			color: "#000"
		},
		search: {
			position: "relative",
			borderRadius: theme.shape.borderRadius,
			backgroundColor: fade(theme.palette.common.black, 1),
			color: theme.palette.common.white,
			"&:hover": {
				backgroundColor: fade(theme.palette.common.black, 1)
			},
			marginRight: theme.spacing(2),
			marginLeft: 0,
			width: "100%",
			[theme.breakpoints.up("sm")]: {
				marginLeft: theme.spacing(3),
				width: "auto"
			}
		},
		searchIcon: {
			width: theme.spacing(7),
			height: "100%",
			position: "absolute",
			pointerEvents: "none",
			display: "flex",
			alignItems: "center",
			justifyContent: "center"
		},
		inputRoot: {
			color: "#fff"
		},
		inputInput: {
			padding: theme.spacing(1, 1, 1, 7),
			transition: theme.transitions.create("width"),
			width: "100%",
			[theme.breakpoints.up("md")]: {
				width: 200
			}
		},
		sectionDesktop: {
			display: "none",
			[theme.breakpoints.up("md")]: {
				display: "flex"
			}
		},
		sectionMobile: {
			display: "flex",
			[theme.breakpoints.up("md")]: {
				display: "none"
			}
		}
	})
);

export default useStyles;