package heuristic

import (
	"gomoku/board"
	"fmt"
)

type Threat struct {
	positions	[]board.Coords
	corner		[]board.Coords
	owner		int8
	size 		int8
	status		int8
	rate		int64
}

const (
	CLOSE_THREAT = 0
	OPEN_THREAT = 1
	HEIGHT = 19
)


func (t Threat) log() {
	fmt.Println("Positions: ", t.positions)
	fmt.Println("corner: ", t.corner)
	fmt.Println("owner: ", t.owner)
	fmt.Println("size: ", t.size)
	fmt.Println("status: ", t.status)
	fmt.Println("rate: ", t.rate)
}
