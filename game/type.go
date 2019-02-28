package game

import (
	"errors"
	"fmt"
	"gomoku/board"
	"gomoku/minimax/heuristic"
)

type State struct {
	Player     int8        `json:"player"`
	BlackScore int8        `json:"blackScore"`
	WhiteScore int8        `json:"whiteScore"`
	Board      board.Board `json:"board"`
	Winner     int8        `json:"winner"`
}

// Move returns new state after applying
func (s State) Move(coords board.Coords) (State, error) {
	if s.Winner != 0 {
		return s, fmt.Errorf("game is finished")
	}
	if !s.Board.CellIsEmpty(coords) {
		return s, fmt.Errorf("cell is occupied")
	}
	s.Board.SetCell(coords, s.Player)
	if caputedCells := GetCaptures(s.Board, coords); len(caputedCells) != 0 {
		for _, capturedCell := range caputedCells {
			s.Board.SetCell(capturedCell, 0)
		}
		if s.Player == board.BLACK_PLAYER {
			s.BlackScore += int8(len(caputedCells))
		} else {
			s.WhiteScore += int8(len(caputedCells))
		}
	} else if heuristic.FindDoubleThreeThreat(s.Board, s.Player, coords) {
		return s, errors.New("double three threat")
	}
	s.Winner = heuristic.IsTerminate(s.Board, s.BlackScore, s.WhiteScore)
	if s.Winner == 0 {
		s.switchPlayer()
	}
	return s, nil
}

func (s *State) switchPlayer() {
	if s.Player == board.BLACK_PLAYER {
		s.Player = board.WHITE_PLAYER
	} else {
		s.Player = board.BLACK_PLAYER
	}
}
