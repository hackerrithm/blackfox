import {
        makeStyles,
        Theme,
        createStyles
} from "@material-ui/core/styles";
import { deepOrange } from "@material-ui/core/colors";

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
                tabs: {
                        borderRight: `1px solid ${theme.palette.divider}`,
                        textAlign: 'center',
                        alignContent: 'center',
                },
                spacesContainer: {
                        textAlign: 'center',
                        alignContent: 'center',
                },
                avatarSquared: {
                        color: theme.palette.getContrastText(deepOrange[500]),
                        backgroundColor: deepOrange[500],
                }
        })
);

export default useStyles;