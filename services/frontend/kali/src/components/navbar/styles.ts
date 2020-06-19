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
		appbarPosition: {
			zIndex: theme.zIndex.drawer + 1,
		},
		grow: {
			flexGrow: 1
		},
		root: {
			flexGrow: 1,
			// marginTop: 0,
			position: "fixed",
			// backgroundColor: theme.palette.common.white,
			// color: theme.palette.common.black,
			// marginBottom: 50
		},
		menuButton: {
			marginRight: theme.spacing(2)
		},
		title: {
			display: "none",
			[theme.breakpoints.up("sm")]: {
				display: "block"
			},
			// color: theme.palette.common.black
		},
		search: {
			position: "relative",
			borderRadius: theme.shape.borderRadius,
			backgroundColor: fade(theme.palette.common.black, 0.1),
			color: theme.palette.common.black,
			"&:hover": {
				backgroundColor: "#5588FF",
				marginRight: theme.spacing(8),
				color: "white"
			},
			marginRight: theme.spacing(9),
			marginLeft: 0,
			width: "1000%",
			[theme.breakpoints.up("sm")]: {
				marginLeft: theme.spacing(9),
				width: "1000%"
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
			color: "#000",
			"&:hover": {
				color: "#fff"
			},
		},
		inputInput: {
			padding: theme.spacing(1, 1, 1, 7),
			transition: theme.transitions.create("width"),
			width: "1000%",
			[theme.breakpoints.up("md")]: {
				width: "auto",
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