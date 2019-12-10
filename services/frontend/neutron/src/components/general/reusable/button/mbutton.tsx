import React from "react";
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      '& > *': {
        margin: theme.spacing(1),
      },
    },
  }),
);

const MButton = (props: any) => {
    const classes = useStyles('');
    // const className = `btn ${props.type}`
    return (
        <Button type={props.type} className={classes.root} variant={props.variant} color={props.color} disabled={props.disabled}>
            {props.label}
        </Button>
    )
}

export default MButton;