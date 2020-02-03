import {
        fade,
        makeStyles,
        Theme,
        createStyles
} from "@material-ui/core/styles";

const useStyles = makeStyles((theme: Theme) =>
        createStyles({
                root: {
                        marginTop: 130,
                        alignContent: "center"
                        // marginBottom: 50
                },
                username: {
                        position: "relative",
                        borderRadius: theme.shape.borderRadius,
                        backgroundColor: fade(theme.palette.common.white, 0.7),
                        WebkitBorderRadius: 1,
                        color: theme.palette.common.black,
                        "&:hover": {
                                backgroundColor: fade(theme.palette.common.white, 0.9),
                                color: "black"
                        },
                        marginRight: theme.spacing(2),
                        marginLeft: 0,
                        width: "200%",
                        [theme.breakpoints.up("sm")]: {
                                marginLeft: theme.spacing(3),
                                width: "600px"
                        }
                },
                password: {
                        position: "relative",
                        borderRadius: theme.shape.borderRadius,
                        backgroundColor: fade(theme.palette.common.white, 0.7),
                        WebkitBorderRadius: 1,
                        color: theme.palette.common.black,
                        "&:hover": {
                                backgroundColor: fade(theme.palette.common.white, 0.9),
                                color: "black"
                        },
                        marginRight: theme.spacing(2),
                        marginLeft: 0,
                        width: "200%",
                        [theme.breakpoints.up("sm")]: {
                                marginLeft: theme.spacing(3),
                                width: "600px"
                        }
                },
                submitButton: {
                        width: "600px",
                        boxShadow: 'none',
                        textTransform: 'none',
                        fontSize: 16,
                        padding: '6px 12px',
                        border: '1px solid',
                        lineHeight: 1.5,
                        backgroundColor: '#007bff',
                        borderColor: '#007bff',
                        marginRight: theme.spacing(-1),
                        marginTop: theme.spacing(2),
                        fontFamily: [
                                '-apple-system',
                                'BlinkMacSystemFont',
                                '"Segoe UI"',
                                'Roboto',
                                '"Helvetica Neue"',
                                'Arial',
                                'sans-serif',
                                '"Apple Color Emoji"',
                                '"Segoe UI Emoji"',
                                '"Segoe UI Symbol"',
                        ].join(','),
                        '&:hover': {
                                backgroundColor: '#0069d9',
                                borderColor: '#5588FF',
                                boxShadow: 'none',
                        },
                        '&:active': {
                                // boxShadow: 'none',
                                backgroundColor: '#5588FF',
                                borderColor: '#005cbf',
                        },
                        '&:focus': {
                                boxShadow: '0 0 0 0.2rem rgba(0,123,255,.5)',
                        },
                },
	})
);

export default useStyles;