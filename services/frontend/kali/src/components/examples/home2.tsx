// import * as React from "react";
// import { IHomeProps } from ".";
// import { gql } from "apollo-boost";
// import { Query } from "react-apollo";
// import List from "./list";
// import { Grid } from "@material-ui/core";
// import DevUpdateToggler from "../profile/userInfoToggler";
// import TestData from "../examples/youtube/testData";
// import InfiniteList from "./infiniteList";


// const GET_EXAMPLE = gql`
// 	query {
// 		getTask(id: "5dbfe03a583ba72328bb89ae") {
// 			text
// 		}
// 	}
// `;
// export default class Home extends React.Component<IHomeProps, any> {
// 	public postList:Array<any> = [];
// 	public postListCounter: number = 0;
// 	public weirdAlgorithm(params: any, params1: any): any {
// 		return params * params1;
// 	}
	
// 	constructor(props: any) {
// 		super(props);
// 		this.state = {
// 			dataList: [],
// 			count: 0
// 		}
// 	}
	
// 	componentDidMount() {
// 		// this.postList = await this.getAllData();
// 		TestData.forEach((x:any) => {
// 			this.postList.push(x)
// 			this.postListCounter++;
// 		})
// 		this.setState({
// 			dataList: this.postList,
// 			count: this.postListCounter,
// 		})
// 	}
	
// 	getAllData() {
// 		var results = TestData.map(x => x)
// 		return results;
// 	}
	
// 	const [state, setState] = useState([]);
// 	/**
// 	 * Render Homepage
// 	 */
// 	public render(): JSX.Element {
// 		let arrayLenth: number = Number(this.postListCounter);
// 		let listSize: number = 8;
// 		return (
// 			<div className="homepage">

// 				<Query query={GET_EXAMPLE}>
// 					{({ loading, error, data }: any) => {
// 						if (loading) {
// 							return <div>Loading...</div>;
// 						}
// 						if (error) {
// 							return <div>No results at the moment</div>;
// 						}

// 						return (
// 							<div>
// 								{data.getTask.text}
// 								{/* {data.getUser.password} */}
// 								{console.log("data here:: ", data)}
// 							</div>
// 						);
// 					}}
// 				</Query>
// 				<Grid container>
// 					<Grid item xs={12} sm={12}>
// 						1
// 					</Grid>
// 					<Grid item xs={3} sm={3}>
// 						<div className={"left-panel"}>
// 							1
// 						</div>
// 					</Grid>
// 					<Grid item xs={6} sm={6}>
// 						<div className={"mid-view"}>
// 							{/* <List postsList={this.state.dataList} initialListCount={listSize} />	 */}
// 							<InfiniteList />
// 						</div>
// 					</Grid>
// 					<Grid item xs={3} sm={3}>
// 						<div className={"right-panel"}>
// 							3
// 						</div>
// 					</Grid>
// 				</Grid>
// 				{/* //</div>    <br/>
//                     //     <br/>
//                     //     <div />

//                     //     <div></div>

//                     //     <div>{false}</div>

//                     //     <div>{null}</div>

//                     //     <div>{undefined}</div>

//                     //     <div>{true}</div>
//                     //     <h1>Home page</h1>
//                     //     <ul>
//                     //     {
//                     //         array.map((element: any, index: any) => {
//                     //             return (
//                     //                 <div key={index}>
//                     //                     <Button variant="contained" color="primary">
//                     //                         {this.weirdAlgorithm(element, index)}
//                     //                     </Button>
//                     //                     <br />
//                     //                     <br />
//                     //                     <br />
//                     //                     <br />
//                     //                     <br />
//                     //                     {index}
//                     //                     <br />
//                     //                     <br />
//                     //                     <br />
//                     //                     <br />
//                     //                     <br />
//                     //                 </div>
//                     //             )
//                     //         })
//                     //     }
//                     //     </ul>
//                     //     {this.props.test}
//                 // </div>
//                     */}
// 			</div>
// 		);
// 	}
// }
