import * as React from "react";
import { IHomeProps } from ".";

export default class Home extends React.Component<IHomeProps, {}> {
    /**
     * Render Homepage
     */
    public render(): JSX.Element {
        return (
            <div>
                <h1>Home page</h1>
                {this.props.test}
            </div>
        );
    }
}
