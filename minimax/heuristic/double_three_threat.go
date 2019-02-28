package heuristic

import (
	"gomoku/board"
)

const (
	MIN_X_5X5_BOARD int = 7
	MAX_X_5X5_BOARD int = 11
	MIN_X_7X7_BOARD int = 6
	MAX_X_7X7_BOARD int = 12

	MIN_Y_5X5_BOARD int = 7
	MAX_Y_5X5_BOARD int = 11
	MIN_Y_7X7_BOARD int = 6
	MAX_Y_7X7_BOARD int = 12
)

func SearchDoubleThreat(b *board.Board, threat []Threat, len int) []Threat {
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
				current := b.GetCell(board.Coords{x, y})
				for i <= len {
					if x+len < 19 && x-1 >= 0 {
						if (y-1 >= 0 && y+len < 19 &&
							(b.GetCell(board.Coords{x - 1, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x + len, y + len}) == 0) &&
							(b.GetCell(board.Coords{x + i, y + i}) == current)) {
							amountRightZ++
							positionsRightZ = append(positionsRightZ, board.Coords{x + i, y + i})
						}
						if ((b.GetCell(board.Coords{x - 1, y}) == 0) &&
							(b.GetCell(board.Coords{x + len, y}) == 0) &&
							(b.GetCell(board.Coords{x + i, y}) == current)) {
							amountX++
							positionsX = append(positionsX, board.Coords{x + i, y})
						}
					}
					if y-1 >= 0 && y+len < 19 {
						if ((b.GetCell(board.Coords{x, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x, y + len}) == 0) &&
							(b.GetCell(board.Coords{x, y + i}) == current)) {
							amountY++
							positionsY = append(positionsY, board.Coords{x, y + i})
						}
						if (x+1 < 19 && x-len >= 0 &&
							(b.GetCell(board.Coords{x + 1, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x - len, y + len}) == 0) &&
							(b.GetCell(board.Coords{x - i, y + i}) == current)) {
							amountLeftZ++
							positionsLeftZ = append(positionsLeftZ, board.Coords{x - i, y + i})
						}
					}
					i++
				}
				if amountX == len {
					corner := []board.Coords{{positionsX[0].X - 1, positionsX[0].Y}, {positionsX[2].X + 1, positionsX[2].Y}}
					threat = append(threat, Threat{owner: current, positions: positionsX, size: int8(len), status: 1, corner: corner})
				}
				if amountY == len {
					corner := []board.Coords{{positionsY[0].X, positionsY[0].Y - 1}, {positionsY[2].X, positionsY[2].Y + 1}}
					threat = append(threat, Threat{owner: current, positions: positionsY, size: int8(len), status: 1, corner: corner})
				}
				if amountRightZ == len {
					corner := []board.Coords{{positionsRightZ[0].X - 1, positionsRightZ[0].Y - 1}, {positionsRightZ[2].X + 1, positionsRightZ[2].Y + 1}}
					threat = append(threat, Threat{owner: current, positions: positionsRightZ, size: int8(len), status: 1, corner: corner})
				}
				if amountLeftZ == len {
					corner := []board.Coords{{positionsLeftZ[0].X + 1, positionsLeftZ[0].Y - 1}, {positionsLeftZ[2].X - 1, positionsLeftZ[2].Y + 1}}
					threat = append(threat, Threat{owner: current, positions: positionsLeftZ, size: int8(len), status: 1, corner: corner})
				}
			}
		}
	}
	return threat
}

func SearchDoubleThreatMid(b *board.Board, threat []Threat, size int) []Threat {
	for y, row := range b {
		for x := range row {
			if (b.GetCell(board.Coords{x, y}) != 0) {
				current := b.GetCell(board.Coords{x, y})
				if x+size+1 < 19 && x-1 >= 0 {
					if (y-1 >= 0 && y+size+1 < 19 && (b.GetCell(board.Coords{x - 1, y - 1}) == 0 && b.GetCell(board.Coords{x + size+1, y + size+1}) == 0)) {
						if b.GetCell(board.Coords{x + 1, y + 1}) == current && b.GetCell(board.Coords{x + 2, y + 2}) == 0 && b.GetCell(board.Coords{x + 3, y + 3}) == current {
							positionsRightZ := []board.Coords{{x, y}, {x + 1, y + 1}, {x + 3, y + 3}}
							cornerRZ := []board.Coords{{positionsRightZ[0].X - 1, positionsRightZ[0].Y - 1}, {positionsRightZ[2].X + 1, positionsRightZ[2].Y + 1}, {positionsRightZ[0].X + 2, positionsRightZ[0].Y + 2}}
							threat = append(threat, Threat{owner: current, positions: positionsRightZ, size: int8(3), status: 1, corner: cornerRZ})
						} else if b.GetCell(board.Coords{x + 1, y + 1}) == 0 && b.GetCell(board.Coords{x + 2, y + 2}) == current && b.GetCell(board.Coords{x + 3, y + 3}) == current {
							positionsRightZ := []board.Coords{{x, y}, {x + 2, y + 2}, {x + 3, y + 3}}
							cornerRZ := []board.Coords{{positionsRightZ[0].X - 1, positionsRightZ[0].Y - 1}, {positionsRightZ[2].X + 1, positionsRightZ[2].Y + 1}, {positionsRightZ[0].X + 1, positionsRightZ[0].Y + 1}}
							threat = append(threat, Threat{owner: current, positions: positionsRightZ, size: int8(3), status: 1, corner: cornerRZ})
						}
					}
					if ((b.GetCell(board.Coords{x - 1, y}) == 0) && (b.GetCell(board.Coords{x + size + 1, y}) == 0)) {
						if b.GetCell(board.Coords{x + 1, y}) == current && b.GetCell(board.Coords{x + 2, y}) == 0 && b.GetCell(board.Coords{x + 3, y}) == current {
							positionsX := []board.Coords{{x, y}, {x + 1, y}, {x + 3, y}}
							cornerX := []board.Coords{{positionsX[0].X - 1, positionsX[0].Y}, {positionsX[2].X + 1, positionsX[2].Y}, {positionsX[0].X + 2, positionsX[0].Y}}
							threat = append(threat, Threat{owner: current, positions: positionsX, size: int8(3), status: 1, corner: cornerX})
						} else if b.GetCell(board.Coords{x + 1, y}) == 0 && b.GetCell(board.Coords{x + 2, y}) == current && b.GetCell(board.Coords{x + 3, y}) == current {
							positionsX := []board.Coords{{x, y}, {x + 2, y}, {x + 3, y}}
							cornerX := []board.Coords{{positionsX[0].X - 1, positionsX[0].Y}, {positionsX[2].X + 1, positionsX[2].Y}, {positionsX[0].X + 1, positionsX[0].Y}}
							threat = append(threat, Threat{owner: current, positions: positionsX, size: int8(3), status: 1, corner: cornerX})
						}
					}
					if y-1 >= 0 && y+size+1 < 19 {
						if ((b.GetCell(board.Coords{x, y - 1}) == 0) && (b.GetCell(board.Coords{x, y + size + 1}) == 0)) {
							if b.GetCell(board.Coords{x, y + 1}) == current && b.GetCell(board.Coords{x, y + 2}) == 0 && b.GetCell(board.Coords{x, y + 3}) == current {
								positionsY := []board.Coords{{x, y}, {x, y + 1}, {x, y + 3}}
								corner := []board.Coords{{positionsY[0].X, positionsY[0].Y - 1}, {positionsY[2].X, positionsY[2].Y + 1}, {positionsY[0].X, positionsY[0].Y + 2}}
								threat = append(threat, Threat{owner: current, positions: positionsY, size: int8(3), status: 1, corner: corner})
							} else if b.GetCell(board.Coords{x, y + 1}) == 0 && b.GetCell(board.Coords{x, y + 2}) == current && b.GetCell(board.Coords{x, y + 3}) == current {
								positionsY := []board.Coords{{x, y}, {x, y + 2}, {x, y + 3}}
								corner := []board.Coords{{positionsY[0].X, positionsY[0].Y - 1}, {positionsY[2].X, positionsY[2].Y + 1}, {positionsY[0].X, positionsY[0].Y + 1}}
								threat = append(threat, Threat{owner: current, positions: positionsY, size: int8(3), status: 1, corner: corner})
							}
						}
						if x+1 < 19 && x-(size+1) >= 0 && b.GetCell(board.Coords{x + 1, y - 1}) == 0 && b.GetCell(board.Coords{x - (size + 1), y + (size + 1)}) == 0 {
							if b.GetCell(board.Coords{x - 1, y + 1}) == 0 && b.GetCell(board.Coords{x - 2, y + 2}) == current && b.GetCell(board.Coords{x - 3, y + 3}) == current {
								positionsLeftZ := []board.Coords{{x, y}, {x - 2, y + 2}, {x - 3, y + 3}}
								corner := []board.Coords{{positionsLeftZ[0].X + 1, positionsLeftZ[0].Y - 1}, {positionsLeftZ[2].X - 1, positionsLeftZ[2].Y + 1}, {positionsLeftZ[0].X - 1, positionsLeftZ[0].Y + 1}}
								threat = append(threat, Threat{owner: current, positions: positionsLeftZ, size: int8(3), status: 1, corner: corner})
							} else if b.GetCell(board.Coords{x - 1, y + 1}) == current && b.GetCell(board.Coords{x - 2, y + 2}) == 0 && b.GetCell(board.Coords{x - 3, y + 3}) == current {
								positionsLeftZ := []board.Coords{{x, y}, {x - 1, y + 1}, {x - 3, y + 3}}
								corner := []board.Coords{{positionsLeftZ[0].X + 1, positionsLeftZ[0].Y - 1}, {positionsLeftZ[2].X - 1, positionsLeftZ[2].Y + 1}, {positionsLeftZ[0].X - 2, positionsLeftZ[0].Y + 2}}
								threat = append(threat, Threat{owner: current, positions: positionsLeftZ, size: int8(3), status: 1, corner: corner})
							}
						}
					}
				}
			}

		}
	}
	return threat
}

func FindDoubleThreeThreat(b board.Board, player int8, cord board.Coords) bool {
	b.SetCell(cord, player)
	var threat []Threat
	threat = SearchDoubleThreat(&b, threat, 3)
	threat = SearchDoubleThreatMid(&b, threat, 3)
	amountPlayer := 0
	flag := 0
	for _, value := range threat {
		if value.owner == player {
			amountPlayer++
			for _, position := range value.positions {
				if position.X == cord.X && position.Y == cord.Y {
					flag++
				}
			}
		}
	}
	if flag == 0 {
		return false
	}
	if amountPlayer == 2 {
		i := 0
		for i < len(threat[1].corner) {
			j := 0
			for j < len(threat[0].corner) {
				if threat[0].corner[j].X == threat[1].corner[i].X && threat[0].corner[j].Y == threat[1].corner[i].Y {
					return false
				}
				j++
			}
			i++
		}
		return true
	}
	return false
}