import React, { Component } from "react";
import {
  BrowserRouter as Router,
  NavLink,
  Route,
  RouteComponentProps
} from "react-router-dom";
import queryString from "querystring";
import GameCreator from "./components/GameCreator/GameCreator";
import Notifications from "./components/Notifications/Notifications";
import Game from "./components/Game/Game";
import { GameType } from "./types";

class App extends Component {
  renderGame = ({
    match: {
      params: { type }
    },
    location: { search }
  }: RouteComponentProps<{
    type: GameType;
  }>) => (
    <Notifications timeout={2e3}>
      {(notifications: string[], notify: (x: string) => void) => (
        <Game
          type={type}
          aiPlayer={+queryString.parse(search.replace("?", "")).aiPlayer}
          notifications={notifications}
          notify={notify}
        />
      )}
    </Notifications>
  );

  render() {
    return (
      <Router>
        <React.Fragment>
          <h1>
            <NavLink to="/">Gomoku</NavLink>
          </h1>
          <Route path="/" exact={true} component={GameCreator} />
          <Route path="/game/:type" exact={true} render={this.renderGame} />
        </React.Fragment>
      </Router>
    );
  }
}

export default App;
