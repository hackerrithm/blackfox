import * as React from "react";
import Home from "./components/home/home";
import Middlwares from "./middleware";

export default class App extends React.Component {
	constructor(props: any) {
		super(props);
	}

	public render() {
		return (
			<div>
				<Middlwares children={this.context} />
			</div>
		);
	}
}
