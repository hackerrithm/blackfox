import { useState, useReducer, useEffect } from "react";

import { DataFetchReducer } from "../book";
import axios from "axios";

const UseDataApi = (initialUrl: any, initialData: any) => {
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

export default UseDataApi;