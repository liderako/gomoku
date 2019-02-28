package minimax

import "gomoku/board"

// IsWithinNFrom checks whether any of the coords withing maximum n steps from coords c
// satisfy predicate pred
func IsWithinNFrom(c board.Coords, n int, pred func(coords board.Coords) bool) bool {
	for y := -n; y <= n; y++ {
		for x := -n; x <= n; x++ {
			neighboringCoords := board.Coords{
				X: c.X + x,
				Y: c.Y + y,
			}
			if neighboringCoords.AreWithin() && pred(neighboringCoords) {
				return true
			}
		}
	}
	return false
}

