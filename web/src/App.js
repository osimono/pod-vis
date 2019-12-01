import React from 'react';
import 'react-bulma-components/dist/react-bulma-components.min.css';
import {Columns, Container, Menu, Navbar} from "react-bulma-components";

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {brand: "Ford"};
    }

    render() {
        return (
            <div>
                <Navbar color="dark" active="true">
                    <Navbar.Brand>
                        <Navbar.Item renderAs="a" href="#">
                            <img src="https://bulma.io/images/bulma-logo.png"
                                 alt="Bulma: a modern CSS framework based on Flexbox" width="112" height="28"/>
                        </Navbar.Item>
                        <Navbar.Burger/>
                    </Navbar.Brand>
                </Navbar>
                <Columns className={"main-columns"}>
                    <Columns.Column size={"one-fifth"}>
                        <Menu>
                            <Menu.List title={clusterName}>
                                <Menu.List.Item>
                                    Info
                                </Menu.List.Item>
                                <Menu.List.Item>
                                    Pods
                                </Menu.List.Item>
                            </Menu.List>
                        </Menu>
                    </Columns.Column>
                </Columns>
            </div>
        );
    }
}

export default App;
