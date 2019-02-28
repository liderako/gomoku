import React from "react";
import { Link } from "react-router-dom";
import { GameType, IPlayer } from "../../types";

const GameCreator: React.FunctionComponent<{}> = function() {
  return (
    <React.Fragment>
      <Link to={`/game/${GameType.vsFriend}`}>Play a Friend</Link>
      <br />
      <Link to={`/game/${GameType.vsComputer}?aiPlayer=${IPlayer.White}`}>
        Play Computer as Black
      </Link>
      <br />
      <Link to={`/game/${GameType.vsComputer}?aiPlayer=${IPlayer.Black}`}>
        Play Computer as White
      </Link>
    </React.Fragment>
  );
};

export default GameCreator;
