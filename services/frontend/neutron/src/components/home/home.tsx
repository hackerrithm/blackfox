import * as React from "react";
import { gql } from "apollo-boost";
import { Query } from "react-apollo";
import { useState } from "react";
import CustomInfiniteScroll from "./customInfiniteScroll";
import './home.css';
import useStyles from "./styles";
import Grid from "@material-ui/core/Grid";
import { Suspense, lazy } from "react";

const Spaces = React.lazy(() => import("./spaces"));
const Trending = React.lazy(() => import("./trending"));

const GET_EXAMPLE = gql`
	query {
		getTask(id: "5dbfe03a583ba72328bb89ae") {
			text
		}
	}
`;



const Home = () => {
	const [state, setState] = useState([]);
	const classes = useStyles("");



	return (
		<div className="homepage">
			{/* <Query query={GET_EXAMPLE}>
				{({ loading, error, data }: any) => {
					if (loading) {
						return <div>Loading...</div>;
					}
					if (error) {
						return <div>No results at the moment</div>;
					}

					return (
						<div>
							{data.getTask.text}
							{data.getUser.password}
						</div>
					);
				}}
			</Query> */}

			<Grid container>
				<Grid item xs={3} sm={3}>
					<div className={"left-panel"}>
						<Suspense fallback={<div>Loading...</div>}>
							<h3 className={"spaces-label"}>Spaces</h3>
							<Spaces state={state} setState={setState} />
						</Suspense>
					</div>
				</Grid>
				<Grid item xs={6} sm={6}>
					<div className={"mid-view"}>
						<CustomInfiniteScroll />
						{/* <List postsList={null} initialListCount={17} state={state} setState={setState} />	 */}
						{/* <InfiniteList state={state} setState={setState} /> */}
						{/* <CustomInfiniteScroll state={state} setState={setState}/> */}
					</div>
				</Grid>
				<Grid item xs={3} sm={3}>
					<div className={"right-panel"}>
						<Suspense fallback={<div>Loading...</div>}>
							<h3 className={"trending-label"}>Trending</h3>
							<Trending state={state} setState={setState} />
						</Suspense>
					</div>
					{/* <div className={"ads-data-right"}>
						<h3>This is an Ad</h3>
					</div> */}
				</Grid>
			</Grid>
		</div>
	);
}

export default Home;