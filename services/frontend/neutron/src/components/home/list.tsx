import React, { useState, useEffect } from "react";
import FacebookExampleFeedLoader from "../examples/facebook/fb";

const List = ({initialListCount}: any) => {
	const [listItems, setListItems] = useState(
		Array.from(Array(initialListCount).keys(), n => n + 1)
	);
	const [isFetching, setIsFetching] = useState(false);

	useEffect(() => {
		window.addEventListener("scroll", handleScroll);
		return () => window.removeEventListener("scroll", handleScroll);
	}, []);

	useEffect(() => {
		if (!isFetching) return;
		fetchMoreListItems();
	}, [isFetching]);

	function handleScroll() {
		if (
			window.innerHeight + document.documentElement.scrollTop !==
				document.documentElement.offsetHeight ||
			isFetching
		)
			return;
		setIsFetching(true);
	}

	function fetchMoreListItems() {
		setTimeout(() => {
			setListItems(prevState => [
				...prevState,
				...Array.from(Array(initialListCount).keys(), n => n + prevState.length + 1)
			]);
			setIsFetching(false);
		}, 2000);
	}

	return (
		<>
			{/* <ul className="list-group mb-2"> */}
				{listItems.map((listItem: any, index: number) => (
					// <li key={index} className="list-group-item">
					// 	List Item {listItem}
					// </li>
					<FacebookExampleFeedLoader />
				))}
			{/* </ul> */}
			{isFetching && <FacebookExampleFeedLoader loading={true}/>}
		</>
	);
};

export default List;
