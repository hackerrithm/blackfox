import React, { useState, useEffect } from "react";
import useStyles from "./styles";
import EnhancedTable from "./leaderboardTable";

const TopTrendingUsers = (props: any) => {
	const classes = useStyles("");
	const [isFetching, setIsFetching] = useState(true);

	useEffect(() => {
		console.log("started now: ")
		if (!isFetching) return;
		getData(isFetching);
		setIsFetching(true);
	}, [isFetching]);

	const getData = (load: any) => {
		if (load) {
			fetch('https://jsonplaceholder.typicode.com/users')
				.then(res => {
					return !res.ok
						? res.json().then(e => Promise.reject(e))
						: res.json();
				})
				.then(res => {
					props.setState([...props.state, ...res]);
				});
		}
	};

	return (
		<>
                        <EnhancedTable usersLeaderboardCurrentList={props.state} />
		</>
	);
};

export default TopTrendingUsers;
