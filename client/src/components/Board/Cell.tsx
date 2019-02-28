import * as React from "react";
import c from "classnames";
import { IBoardCell, ICellClickHandler, IPlayer } from "../../types";
import BoardStyles from "./Board.module.css";

interface ICellProps {
  children: IBoardCell;
  x: number;
  y: number;
  onClick: ICellClickHandler;
}

function stoneClassName(c: IBoardCell): string {
  if (c === 0) return BoardStyles.stone__none;
  if (c === IPlayer.Black) return BoardStyles.stone__black;
  if (c === IPlayer.White) return BoardStyles.stone__white;
  return BoardStyles.stone__suggestion;
}

class Cell extends React.PureComponent<ICellProps> {
  render() {
    const { children: cell, x, y, onClick } = this.props;
    return (
      <div className={BoardStyles.cell} onClick={() => onClick({ x, y })}>
        <div
          title={cell.toString()}
          className={c(BoardStyles.stone, cell && stoneClassName(cell))}
        />
      </div>
    );
  }
}

export default Cell;
