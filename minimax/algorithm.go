package minimax

import (
	"gomoku/board"
	"gomoku/game"
	"math"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func MaxInt64(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func Minimax(state game.State, width, depth int,
	heuristic func(b board.Board, blackScore, whiteScore int8) int64,
	alpha, beta int64) Moves {
	if depth == 0 {
		return Moves{
			{
				State:      state,
				Evaluation: heuristic(state.Board, state.BlackScore, state.WhiteScore),
			},
		}
	}
	// get all cells adjacent to occupied cells
	cellsAdjacentToOccupied := CellsAdjacentToOccupied(state.Board, 1)
	moves := make(Moves, 0, len(cellsAdjacentToOccupied))
	movesCh := make(chan *Move)
	winsCh := make(chan *Move)
	for _, coords := range cellsAdjacentToOccupied {
		coords := coords // needed to use in the goroutine
		go func() {
			state, err := state.Move(coords)
			if err != nil {
				movesCh <- nil
				return
			}
			move := &Move{
				Coords: coords,
				State:  state,
			}
			if state.Winner != 0 {
				if state.Winner == board.BLACK_PLAYER {
					move.Evaluation = math.MinInt64
				} else {
					move.Evaluation = math.MaxInt64
				}
				winsCh <- move
				return
			}
			move.Evaluation = heuristic(state.Board, state.BlackScore, state.WhiteScore)
			movesCh <- move
		}()
	}
	for range cellsAdjacentToOccupied {
		select {
		case move := <-winsCh:
			return Moves{move}
		case move := <-movesCh:
			if move != nil {
				moves = append(moves, move)
			}
		}
	}
	sort.Sort(moves)
	moves = moves[:Min(width, moves.Len())]
	// take best moves and proceed with them, discard the rest
	for _, move := range moves {
		opponentMoves := Minimax(move.State, width, depth-1, heuristic, alpha, beta)
		move.Evaluation = opponentMoves[0].Evaluation
		switch move.State.Player {
		case MIN_PLAYER:
			beta = MinInt64(move.Evaluation, beta)
		case MAX_PLAYER:
			alpha = MaxInt64(move.Evaluation, alpha)
		}
		if beta < alpha {
			break
		}
	}
	sort.Sort(moves)
	return moves
}
