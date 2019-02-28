import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import Modal from "../Modal/Modal";
import {
  GameType,
  IBoard,
  ICommonGameState,
  ICommonGameStateWithBoard,
  ICoords,
  IPlayer,
  ISuggestion
} from "../../types";

import omit from "lodash/fp/omit";
import GameStyles from "./Game.module.css";
import * as API from "../../API";
import Controls from "./Controls";
import Dropdown from "../Dropdown/Dropdown";

export interface IGameProps {
  type: GameType;
  aiPlayer?: IPlayer;
  notify(x: string): void;
  notifications: string[];
}

interface IGameState extends ICommonGameState {
  boardHistory: IBoard[];
  boardHistoryIndex: number;
  suggestions: ISuggestion[];
  aiIsThinking: boolean;
  aiResponseTime: number;
  difficulty: number;
  suggestMoves: boolean;
  showWinnerModal: boolean;
}

function mergeBoardWithSuggestions(
  board: IBoard,
  suggestions: ISuggestion[]
): IBoard {
  if (!suggestions.length) return board;
  const merged = board.map(row => row.slice()) as IBoard;
  suggestions.forEach(({ x, y }, i) => !merged[y][x] && (merged[y][x] = 3 + i));
  return merged;
}

const INITIAL_STATE: IGameState = {
  player: IPlayer.Black,
  blackScore: 0,
  whiteScore: 0,
  winner: 0,
  suggestions: [],
  boardHistory: [Board.init()],
  boardHistoryIndex: 0,
  aiIsThinking: false,
  aiResponseTime: 0,
  difficulty: 0,
  suggestMoves: true,
  showWinnerModal: false
};

export default class Game extends React.Component<IGameProps, IGameState> {
  resetGame = () =>
    this.setState(INITIAL_STATE, () => {
      const { type, aiPlayer } = this.props;
      const { player } = this.state;
      if (type === GameType.vsComputer && player === aiPlayer) this.aiMove();
    });

  componentWillMount() {
    this.resetGame();
  }

  getCurrentBoard = () => {
    const { boardHistory } = this.state;
    return boardHistory[boardHistory.length - 1];
  };

  currentBoardIsDisplayed = (): boolean => {
    const { boardHistory, boardHistoryIndex } = this.state;
    return boardHistoryIndex >= boardHistory.length - 1;
  };

  humanMove = (coords: ICoords) => {
    !this.state.winner &&
      API.makeMove(
        {
          ...this.state,
          board: this.getCurrentBoard()
        },
        coords
      ).then(this.setGameState, this.props.notify);
  };

  aiFetch = () => {
    this.setState({
      aiIsThinking: true
    });
    const now = Date.now();
    return API.suggestMoves(
      {
        ...this.state,
        board: this.getCurrentBoard()
      },
      this.state.difficulty
    ).then(moves => {
      this.setState({
        aiIsThinking: false,
        aiResponseTime: Date.now() - now
      });
      return moves;
    });
  };

  aiMove = () => {
    this.setState({
      suggestions: []
    });
    if (this.state.winner) return;
    this.aiFetch().then(moves => {
      this.setGameState(moves[0].state);
    });
  };

  showSuggestions = () => {
    this.aiFetch().then(suggestions => {
      this.setState({
        suggestions
      });
    });
  };

  setGameState = (state: ICommonGameStateWithBoard) => {
    const {
      winner,
      boardHistory,
      boardHistoryIndex,
      suggestMoves
    } = this.state;
    if (winner) return;
    this.setState(
      {
        boardHistory: [...boardHistory, state.board as IBoard],
        boardHistoryIndex: this.currentBoardIsDisplayed()
          ? boardHistoryIndex + 1
          : boardHistoryIndex,
        ...omit(["board"], state),
        showWinnerModal: state.winner > 0
      },
      () => {
        if (state.player == this.props.aiPlayer) this.aiMove();
        else if (suggestMoves && !this.state.winner) this.showSuggestions();
      }
    );
  };

  handleCellClick = (coords: ICoords) => {
    if (!this.currentBoardIsDisplayed()) return;
    const { type, aiPlayer } = this.props;
    const { player, aiIsThinking } = this.state;

    if (aiIsThinking) {
      return;
    }
    if (
      type === GameType.vsFriend ||
      (type === GameType.vsComputer && player !== aiPlayer)
    ) {
      this.humanMove(coords);
    }
  };

  setDifficulty = ({
    currentTarget: { value }
  }: React.SyntheticEvent<HTMLInputElement>) =>
    this.setState({ difficulty: +value });

  toggleSuggestMoves = () =>
    this.setState(({ suggestMoves }) => ({
      suggestions: [],
      suggestMoves: !suggestMoves
    }));

  render() {
    const { type, notifications } = this.props;
    const {
      player,
      blackScore,
      whiteScore,
      winner,
      suggestions,
      boardHistory,
      boardHistoryIndex,
      aiResponseTime,
      difficulty,
      suggestMoves,
      showWinnerModal
    } = this.state;
    return (
      <div className={GameStyles.container}>
        <div className={GameStyles.centerWrapper}>
          <div className={GameStyles.infoPanel}>
            <CurrentPlayer
              player={player}
              blackScore={blackScore}
              whiteScore={whiteScore}
            />
            {type === GameType.vsComputer && !!aiResponseTime && (
              <div className={GameStyles.responseTime}>{aiResponseTime}ms</div>
            )}
            <Dropdown
              renderButton={(isOpen, toggle) => (
                <button onClick={toggle} className={GameStyles.settingsBtn} />
              )}
              renderContent={() => (
                <div className={GameStyles.dropdown}>
                  <label>Difficulty</label>
                  <input
                    onChange={this.setDifficulty}
                    type="range"
                    value={difficulty}
                    min={0}
                    max={2}
                    step={1}
                  />
                  <label>Suggest moves</label>
                  <input
                    onChange={this.toggleSuggestMoves}
                    type="range"
                    value={+suggestMoves}
                    min={0}
                    max={1}
                    step={1}
                  />
                </div>
              )}
            />
          </div>
          <Board onClick={this.handleCellClick}>
            {this.currentBoardIsDisplayed()
              ? mergeBoardWithSuggestions(this.getCurrentBoard(), suggestions)
              : boardHistory[boardHistoryIndex]}
          </Board>
          <Controls
            i={boardHistoryIndex}
            max={boardHistory.length - 1}
            onChange={i => {
              this.setState({
                boardHistoryIndex: i
              });
            }}
          />
          <div>
            {notifications.map((text, i) => (
              <div className={GameStyles.notification} key={i}>
                {text}
              </div>
            ))}
          </div>
          {showWinnerModal && (
            <Modal onClose={() => this.setState({ showWinnerModal: false })}>
              <>
                <h2>{winner === 1 ? "Black" : "White"} won!</h2>
                <button onClick={this.resetGame}>Restart</button>
              </>
            </Modal>
          )}
        </div>
      </div>
    );
  }
}
