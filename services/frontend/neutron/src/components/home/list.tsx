import React, { useState, useEffect } from "react";
import Post from "../post/post";

const List = (props: any) => {
	// const [listItems, setListItems] = useState(
	// 	Array.from(Array(props.initialListCount).keys(), n => n + 1)
	// );
	// const [items, setItems] = useState(Array.from(Array(initialListCount).keys(), n => n + 1));
	const [isFetching, setIsFetching] = useState(false);


	useEffect(() => {
		// let listCount = count 
		window.addEventListener("scroll", handleScroll);
		return () => window.removeEventListener("scroll", handleScroll);
	}, []);

	useEffect(() => {
		console.log("started now: ")
		if (!isFetching) return;
		getData(isFetching);
		setIsFetching(true);				
	}, [isFetching]);

	function handleScroll() {
		if (
			window.innerHeight + document.documentElement.scrollTop !==
			document.documentElement.offsetHeight ||
			isFetching
		)
			return;
		// setItems(postsList)		
		setIsFetching(true);
	}

	const getData = (load: any) => {
                if (load) {
                        fetch('https://dog.ceo/api/breeds/image/random/15')
                                .then(res => {
                                        return !res.ok
                                                ? res.json().then(e => Promise.reject(e))
                                                : res.json();
                                })
                                .then(res => {
                                        props.setState([...props.state, ...res.message]);
				});
			}		
        };

	function fetchMoreListItems() {
		// setTimeout(() => {
		// 	setListItems(prevState => [
		// 		...prevState,
		// 		...Array.from(Array(props.initialListCount).keys(), n => n + prevState.length + 1)
		// 	]);
		// 	// setItems((prevState: any) => [
		// 	// 	...prevState,
		// 	// 	...Array.from(Array(initialListCount).keys(), n => n + prevState.length + 1)
		// 	// ]);
		// 	// setItems(items)
		// }, 2000);
		setIsFetching(false);
	}

	return (
		<>
			{props.state.map((listItem: any, index: number) => (
				console.log("yah suh ", listItem),
				<Post avatar={"https://images.alphacoders.com/901/901573.jpg"} details={"random stuff"} title={"The Big Picture"} image={listItem} username={"hackerrithm"} fullname={"Kemar G"} />
			))}
			{/* {items.map((a: any, index: number) => {
				console.log("yah suh ", a);
				<Post avatar={"https://images.alphacoders.com/901/901573.jpg"} details={"random stuff"} title={"The Big Picture"} image={postsList[index]} username={"hackerrithm"} fullname={"Kemar G"} />

			})} */}
			{props.state.isFetching && <Post loading={true} />}
		</>
	);
};

export default List;
