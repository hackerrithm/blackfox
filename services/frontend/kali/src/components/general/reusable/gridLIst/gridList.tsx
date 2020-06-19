import React from "react";
import { createStyles, Theme, makeStyles } from "@material-ui/core/styles";
import GridList from "@material-ui/core/GridList";
import GridListTile from "@material-ui/core/GridListTile";
import GridListTileBar from "@material-ui/core/GridListTileBar";
import IconButton from "@material-ui/core/IconButton";
import StarBorderIcon from "@material-ui/icons/StarBorder";
// import tileData from './tileData';

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			display: "flex",
			flexWrap: "wrap",
			justifyContent: "space-around",
			overflow: "hidden",
			backgroundColor: theme.palette.background.paper
		},
		gridList: {
			width: 500,
			height: 450,
			// Promote the list into his own layer on Chrome. This cost memory but helps keeping high FPS.
			transform: "translateZ(0)"
		},
		titleBar: {
			background:
				"linear-gradient(to bottom, rgba(0,0,0,0.7) 0%, " +
				"rgba(0,0,0,0.3) 70%, rgba(0,0,0,0) 100%)"
		},
		icon: {
			color: "white"
		}
	})
);

/**
 * The example data is structured as follows:
 *
 * import image from 'path/to/image.jpg';
 * [etc...]
 *
 * */
const tileData = [
	{
		objectID: 1,
		//   img: image,
		title: "Image 1",
		author: "author",
		featured: true
	},
	{
		objectID: 2,
		// img: image,
		title: "Image 2",
		author: "author",
		featured: true
	},
	{
		objectID: 3,
		// img: image,
		title: "Image 3",
		author: "author",
		featured: true
	}
];

const AdvancedGridList = (goalList: any) => {
	const classes = useStyles("");

	return (
		<div className={classes.root}>
			<GridList cellHeight={200} spacing={1} className={classes.gridList}>
				{goalList.map((goal: any) => (
					<GridListTile
						key={goal.objectID}
						cols={true ? 2 : 1}
						rows={true ? 2 : 1}
					>
						<img src={""} alt={goal.title} />
						<GridListTileBar
							title={goal.title}
							titlePosition="top"
							actionIcon={
								<IconButton
									aria-label={`star ${goal.title}`}
									className={classes.icon}
								>
									<StarBorderIcon />
								</IconButton>
							}
							actionPosition="left"
							className={classes.titleBar}
						/>
					</GridListTile>
				))}
			</GridList>
		</div>
	);
}

export default AdvancedGridList;