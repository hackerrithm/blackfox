import React from 'react';
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';

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

export default function AlignItemsList({list}: any) {
        const classes = useStyles("");
        let bigList = list

        return (
                bigList.map((listItem: any) => {
                        <List className={classes.root}>
                                <ListItem alignItems="flex-start">
                                        <ListItemAvatar>
                                                <Avatar alt="Remy Sharp" src="/static/images/avatar/1.jpg" />
                                        </ListItemAvatar>
                                        <ListItemText
                                                primary="Brunch this weekend?"
                                                secondary={
                                                        <React.Fragment>
                                                                <Typography
                                                                        component="span"
                                                                        variant="body2"
                                                                        className={classes.inline}
                                                                        color="textPrimary"
                                                                >
                                                                        Ali Connors
                </Typography>
                                                                {" — I'll be in your neighborhood doing errands this…"}
                                                        </React.Fragment>
                                                }
                                        />
                                </ListItem>
                        </List>

                })
        );
}
