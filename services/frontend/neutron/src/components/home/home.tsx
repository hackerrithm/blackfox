import * as React from "react";
import { IHomeProps } from ".";
// import Button from "@material-ui/core/Button";

import { gql } from "apollo-boost";
import { Query } from "react-apollo";

const GET_EXAMPLE = gql`
  query {
    getUser(id: "5d7daf474da65a16aa8a6c05") {
        password,
        username,
        name,
        id
    }
  }
`;
export default class Home extends React.Component<IHomeProps, any> {
    
    public weirdAlgorithm(params:any, params1: any): any {
        return(params * params1);
    }

    componentDidMount() {
        console.log(this.context.data)
    }
    
    /**
     * Render Homepage
     */
    public render(): JSX.Element {
        const array = [1,2,3,4,5,6];
        return (
            // <div>
            //     <br/>
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
            <Query query={GET_EXAMPLE}>
                {({ loading, error, data }: any) => {
                if (loading) {
                    return <div>Loading...</div>;
                }
                if (error) {
                    return <div>Error :(</div>;
                }

                return <div>
                        <br/>
                        <br/>
                        here
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        here
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        here
                        <br/>
                        <br/>
                        <br/>
                        {data.getUser.username}
                        {data.getUser.password}
                        {
                            console.log('data here:: ', data)
                        }
                    </div>;
                }}
            </Query>
        );
    }
}
