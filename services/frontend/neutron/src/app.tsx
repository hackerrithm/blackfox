import * as React from "react";
import Middlwares from "./middleware";

export default class App extends React.Component {
	constructor(props: any) {
		super(props);
	}

	public render() {
		return (
				<Middlwares children={this.context} />
		);
	}
}
