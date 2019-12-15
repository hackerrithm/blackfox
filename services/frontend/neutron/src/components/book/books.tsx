import React, { useReducer, useEffect, useState, Fragment } from "react";
import axios from "axios";
import { DataFetchReducer } from ".";
import {
	makeStyles,
	Theme,
	createStyles,
} from "@material-ui/core";
import ExampleLoadPage from "../examples/youtube/yt";
import CustomizedInputBase from "../general/reusable/input/search";
import useStyles from "./styles";

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
					console.log("here is counter: ", result);
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
	const [query, setQuery] = useState("");
	const [{ data, isLoading, isError }, doFetch] = useDataApi(
		"http://hn.algolia.com/api/v1/search?query=japan",
		{
			hits: []
		}
	);

	// const classes = useStylesGoalsSection(getThemeProps);
	const classes = useStyles("");

	const [embla, setEmbla] = useState(null);

	useEffect(() => {
		if (embla) {
			embla.on("select", () => {
				console.log(`Current index is ${embla.selectedScrollSnap()}`);
			});
		}
	}, [embla]);
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
					<ExampleLoadPage loading={true} length={12} />
				</div>
			) : (
				<div>

						{/* <EmblaCarouselReact
							emblaRef={setEmbla}
							options={{
								loop: true,
								draggable: true,
								dragFree: false,
								speed: 10,
								startIndex: 0
							}}
						>
							<div style={{ display: "flex" }}>
								<div>
									<img
										alt={"1"}
										width="500px"
										height="200px"
										src={
											"https://66.media.tumblr.com/e7b8cf45f9d29ffe6644ebeaf9e87419/tumblr_oy6vcn7U7l1vtzqkfo1_400.jpg"
										}
									/>
								</div>
								<div style={{ flex: "0 0 100%" }}>
									{" "}
									<img
										alt={"2"}
										height="200px"
										width="500px"
										src={
											"https://i.ytimg.com/vi/pLqipJNItIo/hqdefault.jpg?sqp=-oaymwEYCNIBEHZIVfKriqkDCwgBFQAAiEIYAXAB&rs=AOn4CLBkklsyaw9FxDmMKapyBYCn9tbPNQ"
										}
									/>
								</div>
								<div style={{ flex: "0 0 100%" }}>
									{" "}
									<img
										alt={"3"}
										height="200px"
										width="500px"
										src={
											"https://i.ytimg.com/vi/kkLk2XWMBf8/hqdefault.jpg?sqp=-oaymwEYCNIBEHZIVfKriqkDCwgBFQAAiEIYAXAB&rs=AOn4CLB4GZTFu1Ju2EPPPXnhMZtFVvYBaw"
										}
									/>
								</div>
							</div>
						</EmblaCarouselReact>
						<button onClick={() => embla.scrollPrev()}>Prev</button>
						<button onClick={() => embla.scrollNext()}>Next</button> */}
					{/* </div> */}
					<div className={classes.root}>
						<ExampleLoadPage data={data.hits} />
					</div>
				</div>
			)}
		</Fragment>
	);
};

export default Books;
