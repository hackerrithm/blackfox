import React, {
	useReducer,
	useEffect,
	useRef,
	useState,
	Fragment
} from "react";
import axios from "axios";
import { DataFetchReducer } from ".";
import ExpandableCard from "../general/reusable/card/card";
import {
	Grid,
	Paper,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core";
import { getThemeProps } from "@material-ui/styles";
import GridList from "@material-ui/core/GridList";
import GridListTile from "@material-ui/core/GridListTile";
import GridListTileBar from "@material-ui/core/GridListTileBar";
import IconButton from "@material-ui/core/IconButton";
import StarBorderIcon from "@material-ui/icons/StarBorder";
import ExampleLoadPage from "../examples/youtube/yt";
import CustomizedInputBase from "../general/reusable/input/search";
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

const useDataApi = (initialUrl: any, initialData: any) => {
	const [url, setUrl] = useState(initialUrl);

	const [state, dispatch] = useReducer(DataFetchReducer, {
		isLoading: false,
		isError: false,
		data: initialData
	});

	useEffect(() => {
		let didCancel = false;

		const fetchData = async () => {
			dispatch({ type: "FETCH_INIT" });

			try {
				const result = await axios(url);

				if (!didCancel) {
					dispatch({ type: "FETCH_SUCCESS", payload: result.data });
				}
			} catch (error) {
				if (!didCancel) {
					dispatch({ type: "FETCH_FAILURE" });
				}
			}
		};

		fetchData();

		return () => {
			didCancel = true;
		};
	}, [url]);

	return [state, setUrl];
};

// function useAsyncHook(searchBook: any) {
// 	const [result, setResult] = useState([]);
// 	const [loading, setLoading] = useState("false");

// 	useEffect(() => {
// 		async function fetchBookList() {
// 			try {
// 				setLoading("true");
// 				const response = await fetch(
// 					`https://www.googleapis.com/books/v1/volumes?q=${searchBook}`
// 				);

// 				const json = await response.json();
// 				// console.log(json);
// 				setResult(
// 					json.items.map((item: any) => {
// 						console.log(item.volumeInfo.title);
// 						return item.volumeInfo.title;
// 					})
// 				);
// 			} catch (error) {
// 				setLoading("null");
// 			}
// 		}

// 		if (searchBook !== "") {
// 			fetchBookList();
// 		}
// 	}, [searchBook]);

// 	return [result, loading];
// }

const useStylesGoalsSection = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			flexGrow: 1
		},
		paper: {
			padding: theme.spacing(2),
			textAlign: "center",
			color: theme.palette.text.secondary
		}
	})
);

const Books = () => {
	let count = 0;
	const [query, setQuery] = useState("");
	const [
		{ data, isLoading, isError },
		doFetch
	] = useDataApi("http://hn.algolia.com/api/v1/search?query=redux", {
		hits: []
	});

	// const classes = useStylesGoalsSection(getThemeProps);
	const classes = useStyles("");
	return (
		<Fragment>
			<CustomizedInputBase
				value={query}
				onChange={(event: any) => setQuery(event.target.value)}
				btnType="submit"
				onSubmit={(event: any) => {
					doFetch(
						`http://hn.algolia.com/api/v1/search?query=${query}`
					);

					event.preventDefault();
				}}
			/>

			{isError && <div>Something went wrong ...</div>}

			{isLoading ? (
				<div>
					<ExampleLoadPage loading={true} />
				</div>
			) : (
				<div className={classes.root}>
					{/* {data.hits.map((item: any) => (
					  <li key={item.objectID}>
						<a href={item.url}>{item.title}</a>
						{/* <MediaCard searchItemID={item.objectID} searchItemTitle={item.title.toString()} searchItemURL={item.url.toString()}/> */}
					{/* </li> */}
					{/* ))} */}
					{/* <AdvancedGridList goalList={data.hits} />
					 */}
					{/* <GridList
						cellHeight={200}
						spacing={1}
						className={classes.gridList}
					>
						{data.hits.map((goal: any) => (
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
					</GridList> */}

					<ExampleLoadPage data={data.hits} />
				</div>

				// <div className={classes.root}>
				// 	<Grid container spacing={3}>
				// 		<Grid item xs={12}>
				// 			<Paper className={classes.paper}>
				// 				{" "}
				// 				<Books />
				// 			</Paper>
				// 		</Grid>
				// 		<Grid item xs={12} sm={6}></Grid>
				// 		<Grid item xs={12} sm={6}></Grid>
				// <ul>
				// {data.hits.map((searchItem: any) => {
				// return (
				// <Grid
				// 	alignContent={"center"}
				// 	item
				// 	xs={6}
				// 	sm={3}
				// 	key={searchItem.objectID}
				// >
				// <li>
				// 	<ExpandableCard
				// 		url={searchItem.url}
				// 		title={searchItem.title}
				// 	/>
				// </li>
				// 									  <li key={searchItem.objectID}>
				// 		<a href={searchItem.url}>{searchItem.title}</a>
				// 	  </li>
				// 				// </Grid>
				// 			// );
				// 		})}
				// 		</ul>
				// // 	</Grid>
				// </div>
			)}
		</Fragment>
	);
};

export default Books;
