package heuristic

import (
	"gomoku/board"
)

func SearchThreatRowOpen(b board.Board, threat []Threat, len int) []Threat{
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
					if x + len < 19 && x - 1 >= 0 {
						if (y - 1 >= 0 && y + len < 19 &&
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
					if y - 1 >= 0 && y + len < 19 {
						if ((b.GetCell(board.Coords{x, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x, y + len}) == 0) &&
							(b.GetCell(board.Coords{x, y + i}) == current)) {
							amountY++
							positionsY = append(positionsY, board.Coords{x, y + i})
						}
						if (x + 1 < 19 && x - len >= 0 &&
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
					threat = append(threat, Threat{owner:current, positions:positionsX, size:int8(len), status:OPEN_THREAT})
				}
				if amountY == len {
					threat = append(threat, Threat{owner:current, positions:positionsY, size:int8(len), status:OPEN_THREAT})
				}
				if amountRightZ == len {
					threat = append(threat, Threat{owner:current, positions:positionsRightZ, size:int8(len), status:OPEN_THREAT})
				}
				if amountLeftZ == len {
					threat = append(threat, Threat{owner:current, positions:positionsLeftZ, size:int8(len), status:OPEN_THREAT})
				}
			}
		}
	}
	return threat
}