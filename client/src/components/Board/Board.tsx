import * as React from "react";
import Row from "./Row";
import { IBoard, ICellClickHandler } from "../../types";

interface IBoardProps {
  children: IBoard;
  onClick: ICellClickHandler;
}

class Board extends React.PureComponent<IBoardProps> {
  static init(): IBoard {
    return Array.from(Array(19)).map(() =>
      Array.from(Array(19)).fill(0)
    ) as IBoard;
  }

  render() {
    const { children: rows, onClick } = this.props;
    return (
      <div>
        {rows.map((row, y) => (
          <Row y={y} onClick={onClick} key={y}>
            {row}
          </Row>
        ))}
      </div>
    );
  }
}

export default Board;
