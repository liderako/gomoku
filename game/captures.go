package game

import "gomoku/board"

func GetCaptures(b board.Board, c board.Coords) []board.Coords {
	player := b.GetCell(c)
	captures := []board.Coords(nil)
	y := c.Y
	x := c.X

	if player == 0 {
		return nil
	}

	if (x+3 < 19 && player == b.GetCell(board.Coords{x + 3, y})) {
		if (b.GetCell(board.Coords{x + 1, y}) != 0 && player != b.GetCell(board.Coords{x + 1, y}) &&
			b.GetCell(board.Coords{x + 1, y}) == b.GetCell(board.Coords{x + 2, y})) {
			captures = append(captures, board.Coords{x + 1, y}, board.Coords{x + 2, y})
		}
	}
	if (x-3 >= 0 && player == b.GetCell(board.Coords{x - 3, y})) {
		if (b.GetCell(board.Coords{x - 1, y}) != 0 && player != b.GetCell(board.Coords{x - 1, y}) &&
			b.GetCell(board.Coords{x - 1, y}) == b.GetCell(board.Coords{x - 2, y})) {
			captures = append(captures, board.Coords{x - 1, y}, board.Coords{x - 2, y})
		}
	}

	if (y+3 < 19 && player == b.GetCell(board.Coords{x, y + 3})) {
		if (b.GetCell(board.Coords{x, y + 1}) != 0 && player != b.GetCell(board.Coords{x, y + 1}) &&
			b.GetCell(board.Coords{x, y + 1}) == b.GetCell(board.Coords{x, y + 2})) {
			captures = append(captures, board.Coords{x, y + 1}, board.Coords{x, y + 2})
		}
	}

	if (y-3 >= 0 && player == b.GetCell(board.Coords{x, y - 3})) {
		if (b.GetCell(board.Coords{x, y - 1}) != 0 && player != b.GetCell(board.Coords{x, y - 1}) &&
			b.GetCell(board.Coords{x, y - 1}) == b.GetCell(board.Coords{x, y - 2})) {
			captures = append(captures, board.Coords{x, y - 1}, board.Coords{x, y - 2})
		}
	}

	if (y+3 < 19 && x+3 < 19 && player == b.GetCell(board.Coords{x + 3, y + 3})) {
		if (b.GetCell(board.Coords{x + 1, y + 1}) != 0 && player != b.GetCell(board.Coords{x + 1, y + 1}) &&
			b.GetCell(board.Coords{x + 1, y + 1}) == b.GetCell(board.Coords{x + 2, y + 2})) {
			captures = append(captures, board.Coords{x + 1, y + 1}, board.Coords{x + 2, y + 2})
		}
	}

	if (y-3 >= 0 && x-3 >= 0 && player == b.GetCell(board.Coords{x - 3, y - 3})) {
		if (b.GetCell(board.Coords{x - 1, y - 1}) != 0 && player != b.GetCell(board.Coords{x - 1, y - 1}) &&
			b.GetCell(board.Coords{x - 1, y - 1}) == b.GetCell(board.Coords{x - 2, y - 2})) {
			captures = append(captures, board.Coords{x - 1, y - 1}, board.Coords{x - 2, y - 2})
		}
	}

	if (y+3 < 19 && x-3 >= 0 && player == b.GetCell(board.Coords{x - 3, y + 3})) {
		if (b.GetCell(board.Coords{x - 1, y + 1}) != 0 && player != b.GetCell(board.Coords{x - 1, y + 1}) &&
			b.GetCell(board.Coords{x - 1, y + 1}) == b.GetCell(board.Coords{x - 2, y + 2})) {
			captures = append(captures, board.Coords{x - 1, y + 1}, board.Coords{x - 2, y + 2})
		}
	}

	if (y-3 >= 0 && x+3 < 19 && player == b.GetCell(board.Coords{x + 3, y - 3})) {
		if (b.GetCell(board.Coords{x + 1, y - 1}) != 0 && player != b.GetCell(board.Coords{x + 1, y - 1}) &&
			b.GetCell(board.Coords{x + 1, y - 1}) == b.GetCell(board.Coords{x + 2, y - 2})) {
			captures = append(captures, board.Coords{x + 1, y - 1}, board.Coords{x + 2, y - 2})
		}
	}

	return captures
}

