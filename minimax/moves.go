package minimax

import (
	"gomoku/board"
	"fmt"
	"gomoku/game"
)

type Move struct {
	board.Coords
	Evaluation int64      `json:"evaluation"`
	State      game.State `json:"state"`
}

type Moves []*Move

func (pq Moves) Len() int {
	return len(pq)
}

func (pq Moves) Less(i, j int) bool {
	if pq.Len() == 0 {
		return true
	}
	if pq[0].State.Player == MIN_PLAYER {
		return pq[i].Evaluation > pq[j].Evaluation
	}
	if pq[0].State.Player == MAX_PLAYER {
		return pq[i].Evaluation < pq[j].Evaluation
	}
	panic(fmt.Errorf("pq.Player must be one of %d, %d, got %d", MIN_PLAYER,
		MAX_PLAYER, pq[0].State.Player))
}

func (pq Moves) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (moves Moves) String() string {
	movesString := ""
	for _, move := range moves {
		movesString += fmt.Sprintf("X: %d Y: %d Eval: %d\n", move.X, move.Y, move.Evaluation)
	}
	return movesString
}

func (moves Moves) Print() {
	fmt.Printf("%s\n", moves)
}
