package heuristic

import (
	"gomoku/board"
)

func searchYTwo(b board.Board, x int, y int, player int8, len int, i int, flag int, amount int) bool {
	if y-1 >= 0 && y+len < HEIGHT {
		if (((b.GetCell(board.Coords{x, y - 1}) == 0 && b.GetCell(board.Coords{x, y + len}) != 0 && b.GetCell(board.Coords{x, y + len}) != player) ||
			(b.GetCell(board.Coords{x, y - 1}) != 0 && b.GetCell(board.Coords{x, y + len}) == 0 && b.GetCell(board.Coords{x, y - 1}) != player)) &&
			(b.GetCell(board.Coords{x, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x, y + i}) == 0 && flag == 0))) {
			return true
		}
	} else if y == 0 && b.GetCell(board.Coords{x, y + len}) == 0 && (b.GetCell(board.Coords{x, y + i}) == player|| (amount != 4 && b.GetCell(board.Coords{x, y + i}) == 0 && flag == 0)) {
		return true
	} else if y-1 >= 0 && b.GetCell(board.Coords{x, y - 1}) == 0 && y+i < HEIGHT && (b.GetCell(board.Coords{x, y + i}) == player|| (amount != 4 && b.GetCell(board.Coords{x, y + i}) == 0 && flag == 0)) {
		return true
	}
	return false
}

func searchXTwo(b board.Board, x int, y int, player int8, len int, i int, flag int, amount int) bool {
	if x+len < HEIGHT && x-1 >= 0 {
		if ((b.GetCell(board.Coords{x - 1, y}) == 0 && b.GetCell(board.Coords{x + len, y}) != 0 && b.GetCell(board.Coords{x + len, y}) != player) ||
			(b.GetCell(board.Coords{x - 1, y}) != 0 && (b.GetCell(board.Coords{x + len, y}) == 0) && b.GetCell(board.Coords{x - 1, y}) != player)) &&
			(b.GetCell(board.Coords{x + i, y}) == player || (amount != 4 && b.GetCell(board.Coords{x + i, y}) == 0 && flag == 0)) {
			return true
		}
	}  else if x == 0 && b.GetCell(board.Coords{x + len, y}) == 0 && (b.GetCell(board.Coords{x + i, y}) == player || (amount != 4 && b.GetCell(board.Coords{x + i, y}) == 0 && flag == 0)) {
		return true
	} else if x-1 >= 0 && b.GetCell(board.Coords{x - 1, y}) == 0 && x+i < HEIGHT && (b.GetCell(board.Coords{x + i, y}) == player || (amount != 4 && b.GetCell(board.Coords{x + i, y}) == 0 && flag == 0)) {
		return true
	}
	return false
}

func searchRightZTwo(b board.Board, x int, y int, player int8, len int, i int, flag int, amount int) bool {
	if (x+len < HEIGHT && y+len < HEIGHT && x-1 >= 0 && y-1 >= 0 &&
		(((b.GetCell(board.Coords{x - 1, y - 1}) == 0 &&
			b.GetCell(board.Coords{x + len, y + len}) != 0 &&
			b.GetCell(board.Coords{x + len, y + len}) != player) ||
			(b.GetCell(board.Coords{x - 1, y - 1}) != 0 && b.GetCell(board.Coords{x + len, y + len}) == 0 && b.GetCell(board.Coords{x - 1, y - 1}) != player)) &&
			(b.GetCell(board.Coords{x + i, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x + i, y + i}) == 0 && flag == 0)))) {
		return true
	}else if (x == 0 || y == 0) && (x+len < HEIGHT && y+len < HEIGHT) && b.GetCell(board.Coords{x + len, y + len}) == 0 && (b.GetCell(board.Coords{x + i, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x + i, y + i}) == 0 && flag == 0))   {
		return true
	} else if (x-1 >= 0 && y-1 >= 0 && b.GetCell(board.Coords{x - 1, y - 1}) == 0) && (x+i < HEIGHT && y+i < HEIGHT && (b.GetCell(board.Coords{x + i, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x + i, y + i}) == 0 && flag == 0))) {
		if y+len < HEIGHT && x+len < HEIGHT && (b.GetCell(board.Coords{x+len, y+len}) == 0 || b.GetCell(board.Coords{x+len, y+len}) == player){
			return false
		}
		return true
	}
	return false
}

func searchLeftZTwo(b board.Board, x int, y int, player int8, len int, i int, flag int, amount int) bool {
	if y-1 >= 0 && y+len < HEIGHT && x+1 < HEIGHT && x-len >= 0 &&
		((b.GetCell(board.Coords{x + 1, y - 1}) == 0 && b.GetCell(board.Coords{x - len, y + len}) != 0 && b.GetCell(board.Coords{x - len, y + len}) != player) ||
			(b.GetCell(board.Coords{x + 1, y - 1}) != 0 && b.GetCell(board.Coords{x - len, y + len}) == 0 && b.GetCell(board.Coords{x + 1, y - 1}) != player)) &&
		(b.GetCell(board.Coords{x - i, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x - i, y + i}) == 0 && flag == 0)) {
		return true
	} else if (x == HEIGHT-1 || y == 0) && (x-i >= 0 && y+len < HEIGHT && x-len >= 0) && b.GetCell(board.Coords{x - len, y + len}) == 0 && (b.GetCell(board.Coords{x - i, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x - i, y + i}) == 0 && flag == 0)) {
		return true
	} else if x+1 < HEIGHT && y-1 >= 0 && b.GetCell(board.Coords{x + 1, y - 1}) == 0 && y+i < HEIGHT && x-i >= 0 && (b.GetCell(board.Coords{x - i, y + i}) == player || (amount != 4 && b.GetCell(board.Coords{x - i, y + i}) == 0 && flag == 0)){
		if y+len < HEIGHT && x-len >= 0 && (b.GetCell(board.Coords{x-len, y+len}) == 0 || b.GetCell(board.Coords{x-len, y+len}) == player) {
			return false
		}
		return true
	}
	return false
}

func SearchThreatComplexFourRowClose(b board.Board, threat []Threat) []Threat {
	len := 5
	for y, row := range b {
		for x := range row {
			if (b.GetCell(board.Coords{x, y}) != 0) {
				i := 0
				positionsX := []board.Coords{}
				positionsY := []board.Coords{}
				positionsRightZ := []board.Coords{}
				positionsLeftZ := []board.Coords{}
				amountX := 0
				flagX := 0
				amountY := 0
				flagY := 0
				amountRightZ := 0
				flagRightZ := 0
				amountLeftZ := 0
				flagLeftZ := 0
				player := b.GetCell(board.Coords{x, y})
				for i < len {
					if searchRightZTwo(b, x, y, player, len, i, flagRightZ, amountRightZ) == true {
						if b.GetCell(board.Coords{x + i, y + i}) == 0 {
							flagRightZ++
						} else {
							amountRightZ++
							positionsRightZ = append(positionsRightZ, board.Coords{x + i, y + i})
						}
					}
					if searchXTwo(b, x, y, player, len, i, flagX, amountX) == true {
						if b.GetCell(board.Coords{x + i, y}) == 0 {
							flagX++
						} else {
							amountX++
							positionsX = append(positionsX, board.Coords{x + i, y})
						}
					}
					if searchYTwo(b, x, y, player, len, i, flagY, amountY) == true {
						if b.GetCell(board.Coords{x, y + i}) == 0 {
							flagY++
						} else {
							amountY++
							positionsY = append(positionsY, board.Coords{x, y + i})
						}
					}
					if searchLeftZTwo(b, x, y, player, len, i, flagLeftZ, amountLeftZ) == true {
						if b.GetCell(board.Coords{x - i, y + i}) == 0 {
							flagLeftZ++
						} else {
							amountLeftZ++
							positionsLeftZ = append(positionsLeftZ, board.Coords{x - i, y + i})
						}
					}
					i++
				}
				if amountX == 4 && flagX == 1 {
					threat = append(threat, Threat{owner: player, positions: positionsX, size: int8(4), status: CLOSE_THREAT})
				}
				if amountY == 4 && flagY == 1 {
					threat = append(threat, Threat{owner: player, positions: positionsY, size: int8(4), status: CLOSE_THREAT})
				}
				if amountRightZ == 4 && flagRightZ == 1 {
					threat = append(threat, Threat{owner: player, positions: positionsRightZ, size: int8(4), status: CLOSE_THREAT})
				}
				if amountLeftZ == 4 && flagLeftZ == 1 {
					threat = append(threat, Threat{owner: player, positions: positionsLeftZ, size: int8(4), status: CLOSE_THREAT})
				}
			}
		}
	}
	return threat
}
