import React from "react";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Skeleton from "@material-ui/lab/Skeleton";
import TestData from "./testData";
import { makeStyles, Theme, createStyles, Popover } from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		popover: {
			pointerEvents: "none"
		},
		paper: {
			padding: theme.spacing(1)
		}
	})
);

interface MediaProps {
	loading?: boolean;
	data?: any;
	length?: any;
}

export default function ExampleLoadPage(props: MediaProps) {
	const { loading = false, data, length } = props;

	const [anchorEl, setAnchorEl] = React.useState(null);

	const handlePopoverOpen = (event: any) => {
		setAnchorEl(event.currentTarget);
	};

	const handlePopoverClose = () => {
		setAnchorEl(null);
	};

	const open = Boolean(anchorEl);

	const classes = useStyles("");

	return (
		<Grid container>
			{(loading ? Array.from(new Array(length)) : data).map(
				(item: any, index: number) => (
					<Box
						key={index}
						width={250}
						marginRight={1.0}
						marginLeft={1.0}
						padding={0}
						my={1}
					>
						{item ? (
							// TODO: Change how image is decided
							<img
								style={{ width: 250, height: 218 }}
								alt={item.title}
								src={
									TestData[index]
										? TestData[index].src
										: TestData[
												Math.floor(
													Math.random() *
														(TestData.length -
															1 -
															0 +
															1)
												) + 0
										  ].src
								}
								aria-owns={
									open ? "mouse-over-popover" : undefined
								}
								aria-haspopup="true"
								onMouseEnter={handlePopoverOpen}
								onMouseLeave={handlePopoverClose}
							/>
						) : (
							<Skeleton variant="rect" width={210} height={118} />
						)}
						{item ? (
							<Box pr={2}>
								<Typography gutterBottom variant="body2">
									{item.title}
								</Typography>
								<Typography
									display="block"
									variant="caption"
									color="textSecondary"
								>
									{item.url}
								</Typography>
								<Typography
									variant="caption"
									color="textSecondary"
								>
									{`${item.relevancy_score} • ${item.createdAt}`}
								</Typography>
							</Box>
						) : (
							<>
								<Box pt={0.5}>
									<Skeleton />
									<Skeleton width="60%" />
								</Box>
								<br />
							</>
						)}
					</Box>
				)
			)}
			<Popover
				id="mouse-over-popover"
				className={classes.popover}
				classes={{
					paper: classes.paper
				}}
				open={open}
				anchorEl={anchorEl}
				anchorOrigin={{
					vertical: "top",
					horizontal: "right"
				}}
				transformOrigin={{
					vertical: "top",
					horizontal: "left"
				}}
				onClose={handlePopoverClose}
				disableRestoreFocus
			>
				<Box pr={2}>
					<Typography gutterBottom variant="body2">
						Big Detail
					</Typography>
					<Typography
						display="block"
						variant="caption"
						color="textSecondary"
					>
						A lot more details
					</Typography>
					<Typography variant="caption" color="textSecondary">
						Some more details
					</Typography>
				</Box>
			</Popover>
		</Grid>
	);
}

// export default function YouTube() {
// 	return (
// 		<Box overflow="hidden">

// 		</Box>
// 	);
// }
