package heuristic

import (
	"gomoku/board"
)

func searchY(b board.Board, x int, y int, player int8, len int, i int) bool {
	if y-1 >= 0 && y+len < HEIGHT {
		if (((b.GetCell(board.Coords{x, y - 1}) == 0 && b.GetCell(board.Coords{x, y + len}) != 0 && b.GetCell(board.Coords{x, y + len}) != player) ||
			(b.GetCell(board.Coords{x, y - 1}) != 0 && b.GetCell(board.Coords{x, y + len}) == 0 && b.GetCell(board.Coords{x, y - 1}) != player)) &&
			(b.GetCell(board.Coords{x, y + i}) == player)) {
			return true
		}
	} else if y == 0 && b.GetCell(board.Coords{x, y + len}) == 0 && (b.GetCell(board.Coords{x, y + i}) == player) {
		return true
	} else if y-1 >= 0 && b.GetCell(board.Coords{x, y - 1}) == 0 && y+i < HEIGHT && (b.GetCell(board.Coords{x, y + i}) == player) {
		return true
	}
	return false
}

func searchX(b board.Board, x int, y int, player int8, len int, i int) bool {
	if x+len < HEIGHT && x-1 >= 0 {
		if ((b.GetCell(board.Coords{x - 1, y}) == 0 && b.GetCell(board.Coords{x + len, y}) != 0 && b.GetCell(board.Coords{x + len, y}) != player) ||
			(b.GetCell(board.Coords{x - 1, y}) != 0 && (b.GetCell(board.Coords{x + len, y}) == 0) && b.GetCell(board.Coords{x - 1, y}) != player)) &&
			b.GetCell(board.Coords{x + i, y}) == player {
			return true
		}
	} else if x == 0 && b.GetCell(board.Coords{x + len, y}) == 0 && (b.GetCell(board.Coords{x + i, y}) == player) {
		return true
	} else if x-1 >= 0 && b.GetCell(board.Coords{x - 1, y}) == 0 && x+i < HEIGHT && (b.GetCell(board.Coords{x + i, y}) == player) {
		return true
	}
	return false
}

func searchRightZ(b board.Board, x int, y int, player int8, len int, i int) bool {
	if (x+len < HEIGHT && y+len < HEIGHT && x-1 >= 0 && y-1 >= 0 &&
		(((b.GetCell(board.Coords{x - 1, y - 1}) == 0 &&
			b.GetCell(board.Coords{x + len, y + len}) != 0 &&
			b.GetCell(board.Coords{x + len, y + len}) != player) ||
			(b.GetCell(board.Coords{x - 1, y - 1}) != 0 && b.GetCell(board.Coords{x + len, y + len}) == 0 && b.GetCell(board.Coords{x - 1, y - 1}) != player)) &&
			(b.GetCell(board.Coords{x + i, y + i}) == player))) {
		return true
	} else if (x == 0 || y == 0) && (x+len < HEIGHT && y+len < HEIGHT) && b.GetCell(board.Coords{x + len, y + len}) == 0 && (b.GetCell(board.Coords{x + i, y + i}) == player) {
		return true
	} else if (x-1 >= 0 && y-1 >= 0 && b.GetCell(board.Coords{x - 1, y - 1}) == 0) && (x+i < HEIGHT && y+i < HEIGHT && b.GetCell(board.Coords{x + i, y + i}) == player) {
		if y+len < HEIGHT && x+len < HEIGHT && (b.GetCell(board.Coords{x+len, y+len}) == 0 || b.GetCell(board.Coords{x+len, y+len}) == player){
			return false
		}
		return true
	}
	return false
}

func searchLeftZ(b board.Board, x int, y int, player int8, len int, i int) bool {
	if y-1 >= 0 && y+len < HEIGHT && x+1 < HEIGHT && x-len >= 0 &&
		((b.GetCell(board.Coords{x + 1, y - 1}) == 0 && b.GetCell(board.Coords{x - len, y + len}) != 0 && b.GetCell(board.Coords{x - len, y + len}) != player) ||
			(b.GetCell(board.Coords{x + 1, y - 1}) != 0 && b.GetCell(board.Coords{x - len, y + len}) == 0 && b.GetCell(board.Coords{x + 1, y - 1}) != player)) &&
		b.GetCell(board.Coords{x - i, y + i}) == player {
		return true
	} else if (x == HEIGHT-1 || y == 0) && (x-i >= 0 && y+len < HEIGHT && x-len >= 0) && b.GetCell(board.Coords{x - len, y + len}) == 0 && b.GetCell(board.Coords{x - i, y + i}) == player {
		return true
	} else if x+1 < HEIGHT && y-1 >= 0 && b.GetCell(board.Coords{x + 1, y - 1}) == 0 && y+i < HEIGHT && x-i >= 0 && b.GetCell(board.Coords{x - i, y + i}) == player {
		if y+len < HEIGHT && x-len >= 0 && (b.GetCell(board.Coords{x-len, y+len}) == 0 || b.GetCell(board.Coords{x-len, y+len}) == player) {
			return false
		}
		return true
	}
	return false
}

func SearchThreatRowClose(b board.Board, threat []Threat, len int) []Threat {
	for y, row := range b {
		for x := range row {
			if (b.GetCell(board.Coords{x, y}) != 0) {
				i := 0
				positionsX := []board.Coords{}
				positionsY := []board.Coords{}
				positionsRightZ := []board.Coords{}
				positionsLeftZ := []board.Coords{}
				amountX := 0
				amountY := 0
				amountRightZ := 0
				amountLeftZ := 0
				player := b.GetCell(board.Coords{x, y})
				for i < len {
					if searchRightZ(b, x, y, player, len, i) == true {
						amountRightZ++
						positionsRightZ = append(positionsRightZ, board.Coords{x + i, y + i})
					}
					if searchX(b, x, y, player, len, i) == true {
						amountX++
						positionsX = append(positionsX, board.Coords{x + i, y})
					}
					if searchY(b, x, y, player, len, i) == true {
						amountY++
						positionsY = append(positionsY, board.Coords{x, y + i})
					}
					if searchLeftZ(b, x, y, player, len, i) == true {
						amountLeftZ++
						positionsLeftZ = append(positionsLeftZ, board.Coords{x - i, y + i})
					}
					i++
				}
				if amountX == len {
					threat = append(threat, Threat{owner: player, positions: positionsX, size: int8(len), status: CLOSE_THREAT})
				}
				if amountY == len {
					threat = append(threat, Threat{owner: player, positions: positionsY, size: int8(len), status: CLOSE_THREAT})
				}
				if amountRightZ == len {
					threat = append(threat, Threat{owner: player, positions: positionsRightZ, size: int8(len), status: CLOSE_THREAT})
				}
				if amountLeftZ == len {
					threat = append(threat, Threat{owner: player, positions: positionsLeftZ, size: int8(len), status: CLOSE_THREAT})
				}
			}
		}
	}
	return threat
}
