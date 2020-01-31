import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { makeStyles, Card, CardActionArea } from '@material-ui/core';
import React from 'react';

const useStyles = makeStyles({
  card: {
    maxWidth: 345,
  },
  media: {
    height: 140,
  },
});

export default function BasicInfoCard({searchItemID, searchItemTitle, searchItemURL}: any) {
  const classes = useStyles("");

  return (
    <Card className={classes.card}>
      <CardActionArea>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {searchItemURL}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
          {searchItemID}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}