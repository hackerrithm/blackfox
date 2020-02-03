import {
	fade,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core/styles";

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			width: '100%',
			maxWidth: 360,
			backgroundColor: theme.palette.background.paper,
		},
	})
);

export default useStyles;