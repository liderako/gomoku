package minimax

import "gomoku/board"

func CellsAdjacentToOccupied(b board.Board, breadth int) []board.Coords {
	moves := make([]board.Coords, 0, len(b))
	if b == (board.Board{}) {
		// Suggest to place into the center of the board if it's empty
		moves = append(moves, board.Coords{X: 10, Y: 10})
	} else {
		// Go through all cells and add those that are empty and withing 2 cells
		// from occupied to the set of possible moves
		// In golang, set data structures are implemented as map[key]bool
		b.ForEach(func(cell int8, coords board.Coords) {
			if b.CellIsEmpty(coords) && IsWithinNFrom(coords, breadth, b.CellIsOccupied) {
				moves = append(moves, coords)
			}
		})
	}
	return moves
}
