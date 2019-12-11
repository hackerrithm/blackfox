import React from "react";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Skeleton from "@material-ui/lab/Skeleton";

interface MediaProps {
	loading?: boolean;
	data?: any;
}

export default function ExampleLoadPage(props: MediaProps) {
	const { loading = false, data } = props;

	return (
		<Grid container wrap="nowrap">
			{(loading ? Array.from(new Array(5)) : data).map(
				(item: any, index: number) =>
						<Box key={index} width={210} marginRight={0.5} my={5}>
							{item ? (
								<img
									style={{ width: 210, height: 118 }}
									alt={item.title}
									src={item.src}
								/>
							) : (
								<Skeleton
									variant="rect"
									width={210}
									height={118}
								/>
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
										{item.channel}
									</Typography>
									<Typography
										variant="caption"
										color="textSecondary"
									>
										{`${item.views} • ${item.createdAt}`}
									</Typography>
								</Box>
							) : (
								<>
								<Box pt={0.5}>
									<Skeleton />
									<Skeleton width="60%" />
								</Box>
								<br/>
								</>
							)}
						</Box>
			)}
		</Grid>
	);
}

// export default function YouTube() {
// 	return (
// 		<Box overflow="hidden">

// 		</Box>
// 	);
// }
