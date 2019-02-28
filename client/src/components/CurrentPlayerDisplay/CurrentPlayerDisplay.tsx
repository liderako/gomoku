import * as React from "react";
import { IPlayer } from "../../types";
import CurrentPlayerDisplayStyles from "./CurrentPlayerDisplay.module.css";
import c from "classnames";

interface ICurrentPlayerDisplayProps {
  player: IPlayer;
  blackScore: number;
  whiteScore: number;
}

const CurrentPlayerDisplay: React.FunctionComponent<
  ICurrentPlayerDisplayProps
> = ({ player, blackScore, whiteScore }) => (
  <div className={CurrentPlayerDisplayStyles.container}>
    <div
      className={c(
        CurrentPlayerDisplayStyles.btn,
        CurrentPlayerDisplayStyles.btn__black,
        player === IPlayer.Black && CurrentPlayerDisplayStyles.btn__selected
      )}
    >
      {blackScore}
    </div>
    <div
      className={c(
        CurrentPlayerDisplayStyles.btn,
        CurrentPlayerDisplayStyles.btn__white,
        player === IPlayer.White && CurrentPlayerDisplayStyles.btn__selected
      )}
    >
      {whiteScore}
    </div>
  </div>
);

export default CurrentPlayerDisplay;
