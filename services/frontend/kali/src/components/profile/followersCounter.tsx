import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";

const useStyles = makeStyles({
	table: {
		minWidth: 100
	}
});

export default function InfluenceMetricSummary({
	followers,
	following,
	level,
	influence
}: any) {
	const classes = useStyles("");

	return (
		<Table className={classes.table} aria-label="simple table">
			<TableHead>
				<TableRow>
					<TableCell align="right">Followers</TableCell>
					<TableCell align="right">Following</TableCell>
					<TableCell align="right">Level</TableCell>
					<TableCell align="right">Influence</TableCell>
				</TableRow>
			</TableHead>
			<TableBody>
				{/* {rows.map(row => ( */}
				<TableRow key={1}>
					<TableCell align="right">{followers}</TableCell>
					<TableCell align="right">{following}</TableCell>
					<TableCell align="right">{level}</TableCell>
					<TableCell align="right">{influence}</TableCell>
				</TableRow>
				{/* ))} */}
			</TableBody>
		</Table>
	);
}
