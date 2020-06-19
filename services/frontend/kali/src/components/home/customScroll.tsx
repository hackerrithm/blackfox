// import React, { useState, useEffect } from "react";
// import Post from "../post/post";

// const CustomInfiniteScroll = (props: any) => {
//         const [loadMore, setLoadMore] = useState(true);

// 	useEffect(() => {
// 		// let listCount = count 
// 		window.addEventListener("scroll", handleScroll);
// 		return () => window.removeEventListener("scroll", handleScroll);
//         }, []);
        
//         const getData = (load: any) => {
//                 if (load) {
//                         fetch('https://dog.ceo/api/breeds/image/random/15')
//                                 .then(res => {
//                                         return !res.ok
//                                                 ? res.json().then(e => Promise.reject(e))
//                                                 : res.json();
//                                 })
//                                 .then(res => {
//                                         props.setState([...props.state, ...res.message]);
//                                 });
//                 }
//         };

// 	// useEffect(() => {
//         //         getData(loadMore);
//         //         setLoadMore(true);
//         // }, [loadMore]);
        
//         useEffect(() => {
// 		// console.log("started now: ", items)
// 		if (!loadMore) return;
// 		getData(true);
// 	}, [loadMore]);

// 	function handleScroll() {
// 		if (
// 			window.innerHeight + document.documentElement.scrollTop !==
// 			document.documentElement.offsetHeight ||
// 			loadMore
// 		)
// 			return;
//                 setLoadMore(true);
// 	}

// 	return (
// 		<>
//                         {props.state.map((img: any, index: number) => (
//                                 <Post key={index} avatar={"https://images.alphacoders.com/901/901573.jpg"} details={"random stuff"} title={"The Big Picture"} image={img} username={"hackerrithm"} fullname={"Kemar G"} />
//                         ))}
//                         {loadMore && <Post loading={true} />}
// 		</>
// 	);
// };

// export default CustomInfiniteScroll;
