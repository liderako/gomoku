package heuristic

import (
	"gomoku/board"
)

// IsTerminate a function that checks if there is a winner, and returns who won.
func IsTerminateFigure(b board.Board) int8 {
	threat := []Threat{}
	threat = SearchThreatRowClose(b, threat, 2)
	for y, row := range b {
		for x := range row {
			if (b.GetCell(board.Coords{x, y}) != 0) {
				current := b.GetCell(board.Coords{x, y})
				i := 0
				amountX := 0
				amountY := 0
				amountRightZ := 0
				amountLeftZ := 0
				for i < 5 {
					if (x+i < HEIGHT && b.GetCell(board.Coords{x + i, y}) == current) {
						if len(threat) == 0 || isCapture(threat, x+i, y) == 1 {
							amountX++
						}
					}
					if (y+i < HEIGHT && b.GetCell(board.Coords{x, y + i}) == current) {
						if len(threat) == 0 || isCapture(threat, x, y+i) == 1 {
							amountY++
						}
					}
					if (x+i < HEIGHT && y+i < HEIGHT && b.GetCell(board.Coords{x + i, y + i}) == current) {
						if len(threat) == 0 || isCapture(threat, x+i, y+i) == 1 {
							amountRightZ++
						}
					}
					if (x-i >= 0 && y+i < HEIGHT && b.GetCell(board.Coords{x - i, y + i}) == current) {
						if len(threat) == 0 || isCapture(threat, x-i, y+i) == 1 {
							amountLeftZ++
						}
					}
					i++
				}
				if amountX == 5 {
					return current
				}
				if amountY == 5 {
					return current
				}
				if amountRightZ == 5 {
					return current
				}
				if amountLeftZ == 5 {
					return current
				}
			}
		}
	}
	return 0
}

func isCapture(threat []Threat, x int, y int) int {
	for _, value := range threat {
		if value.positions[0].X == x && value.positions[0].Y == y {
			return 0
		} else if value.positions[1].X == x && value.positions[1].Y == y {
			return 0
		}
	}
	return 1
}

func IsTerminate(b board.Board, amountPointMinPlayer int8, amountPointMaxPlayer int8) int8 {
	if amountPointMinPlayer >= 10 {
		return board.BLACK_PLAYER
	} else if amountPointMaxPlayer >= 10 {
		return board.WHITE_PLAYER
	}
	return IsTerminateFigure(b)
}
