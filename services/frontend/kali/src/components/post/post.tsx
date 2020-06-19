import React from "react";
import { createStyles, Theme, makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Avatar from "@material-ui/core/Avatar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import Skeleton from "@material-ui/lab/Skeleton";
import CardActions from "@material-ui/core/CardActions";
import FavoriteIcon from '@material-ui/icons/Favorite';
import ShareIcon from '@material-ui/icons/Share';
import CommentIcon from '@material-ui/icons/Comment';

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		card: {
			maxWidth: 524,
			margin: theme.spacing(2)
		},
		media: {
			height: 390
		}
	})
);

interface MediaProps {
	loading?: boolean;
	avatar?: any;
	username?: any;
	fullname?: any;
	title?: any;
	tags?: any;
	details?: any;
	status?: any;
	image?: any;
	followers?: any;
	likes?: any;
	shares?: any;
	impact?: any;
	datePosted?: any;
	age?: any;
}

const Post = (props: MediaProps) => {
	const { loading = false, avatar, username, fullname, title, tags, details, status, image, followers, likes, shares, impact, datePosted, age } = props;
	const classes = useStyles("");

	return (
		<Card className={classes.card}>
			<CardHeader
				avatar={
					loading ? (
						<Skeleton variant="circle" width={30} height={30} />
					) : (
							<Avatar
								alt={username}
								src={avatar}
							/>
						)
				}
				action={
					loading ? null : (
						<IconButton aria-label="settings">
							<MoreVertIcon />
						</IconButton>
					)
				}
				title={
					loading ? (
						<Skeleton
							height={10}
							width="80%"
							style={{ marginBottom: 6 }}
						/>
					) : (
							title
						)
				}
				subheader={
					loading ? (
						<Skeleton height={10} width="40%" />
					) : (
							age
						)
				}
			/>
			{loading ? (
				<Skeleton variant="rect" className={classes.media} />
			) : (
					<CardMedia
						className={classes.media}
						image={image}
						title={title}
					/>
				)}
			<CardContent>
				{loading ? (
					<React.Fragment>
						<Skeleton height={10} style={{ marginBottom: 6 }} />
						<Skeleton height={10} width="80%" />
					</React.Fragment>
				) : (
						<Typography
							variant="body2"
							color="textSecondary"
							component="p"
						>
							{details}
						</Typography>
					)}
			</CardContent>
			<CardActions disableSpacing>
				<IconButton aria-label="add to favorites">
					<FavoriteIcon />
				</IconButton>
				<IconButton aria-label="share">
					<ShareIcon />
				</IconButton>
				<IconButton aria-label="Comment">
					<CommentIcon />
				</IconButton>
			</CardActions>
		</Card>
	);
}

export default Post;
