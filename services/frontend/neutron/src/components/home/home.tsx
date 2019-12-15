import * as React from "react";
import { IHomeProps } from ".";
import { gql } from "apollo-boost";
import { Query } from "react-apollo";
import {
	Grid,
	Paper,
	makeStyles,
	Theme,
	createStyles
} from "@material-ui/core";
import { getThemeProps } from "@material-ui/styles";
import CustomizedInputBase from "../general/reusable/input/search";
import YouTube from "../examples/youtube/yt";
import Books from "../book/books";


const GET_EXAMPLE = gql`
	query {
		getUser(id: "5d7daf474da65a16aa8a6c05") {
			password
			username
			name
			id
		}
	}
`;
export default class Home extends React.Component<IHomeProps, any> {
	public weirdAlgorithm(params: any, params1: any): any {
		return params * params1;
	}

	componentDidMount() {
		console.log(this.context.data);
	}

	/**
	 * Render Homepage
	 */
	public render(): JSX.Element {
		const array = [1, 2, 3, 4, 5, 6];
		return (
			<div className="homepage">
				<br />
				<br />
				<br />
				<br />
				{/** GOALS */}
				<GoalsSection />
				{/** GOALS */}
				<Query query={GET_EXAMPLE}>
					{({ loading, error, data }: any) => {
						if (loading) {
							return <div>Loading...</div>;
						}
						if (error) {
							return <div>No results at the moment</div>;
						}

						return (
							<div>
								{data.getUser.username}
								{data.getUser.password}
								{console.log("data here:: ", data)}
							</div>
						);
					}}
				</Query>
				{/* //</div>    <br/>
                    //     <br/>
                    //     <div />

                    //     <div></div>

                    //     <div>{false}</div>

                    //     <div>{null}</div>

                    //     <div>{undefined}</div>

                    //     <div>{true}</div>
                    //     <h1>Home page</h1>
                    //     <ul>
                    //     {
                    //         array.map((element: any, index: any) => {
                    //             return (
                    //                 <div key={index}>
                    //                     <Button variant="contained" color="primary">
                    //                         {this.weirdAlgorithm(element, index)}
                    //                     </Button>
                    //                     <br />
                    //                     <br />
                    //                     <br />
                    //                     <br />
                    //                     <br />
                    //                     {index}
                    //                     <br />
                    //                     <br />
                    //                     <br />
                    //                     <br />
                    //                     <br />
                    //                 </div>
                    //             )
                    //         })
                    //     }
                    //     </ul>
                    //     {this.props.test}
                // </div>
                    */}
			</div>
		);
	}
}

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

function GoalsSection() {
	const classes = useStylesGoalsSection(getThemeProps);
	const goalsList = [1, 2, 3, 4];
	return (
		<div className={classes.root}>
			<Grid container spacing={3}>
				<Grid item xs={12}>
					<Paper className={classes.paper}>
						{" "}
						<Books />
					</Paper>
				</Grid>
				<Grid item xs={12} sm={6}></Grid>
				<Grid item xs={12} sm={6}></Grid>
				{/* {goalsList.map((val, index) => {
					return (
						<Grid
							alignContent={"center"}
							item
							xs={6}
							sm={3}
							key={index}
						>
							<ExpandableCard />
						</Grid>
					);
				})} */}
			</Grid>
		</div>
	);
}
