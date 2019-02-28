package board

import (
	"math/rand"
	"time"
	"fmt"
)

type Row [19]int8
type Board [19]Row

func (b Board) String() string {
	str := ""
	for _, row := range b {
		for _, c := range row {
			if c == 1 || c == 2 {
				str += fmt.Sprintf("%d", c)
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

// IsEmpty if the board is completely empty
func (b Board) IsEmpty(c Coords) bool {
	return b == Board{}
}

// GetCell returns a cell under coordinates
func (b Board) GetCell(c Coords) int8 {
	return b[c.Y][c.X]
}

func (b *Board) SetCell(c Coords, val int8) int8 {
	b[c.Y][c.X] = val
	return val
}

func (b Board) SetCellImmut(c Coords, val int8) Board {
	b[c.Y][c.X] = val
	return b
}

// SetCells takes an array of coordinates and values ​​and sets them on the board
func (b *Board) SetCells(coords []Coords, vals []int8) {
	for i, c := range coords {
		b.SetCell(c, vals[i]) // not very safe, since i may be out of the range
	}
}

// CellIsEmpty tells if a particular cell is not occupied
func (b Board) CellIsEmpty(c Coords) bool {
	return b.GetCell(c) == 0
}

// CellIsEmpty tells if a particular cell is not occupied
func (b Board) CellIsOccupied(c Coords) bool {
	return !b.CellIsEmpty(c)
}

// ForEach calls a provided function on every board cell with its values and coordinates
func (b Board) ForEach(f func(int8, Coords)) {
	for y, row := range b {
		for x, cell := range row {
			f(cell, Coords{
				X: x,
				Y: y,
			})
		}
	}
}

func (b Board) RandomCoords() Coords {
	rand.Seed(time.Now().UnixNano())
	return Coords{
		X: rand.Intn(len(b)),
		Y: rand.Intn(len(b)),
	}
}

// Fill n random cells with val
func (b *Board) FillRandomly(n int, val int8) {
	coords := make(map[Coords]bool)
	for ; n > 0; n-- {
		randomCoords := b.RandomCoords()
		_, ok := coords[randomCoords]
		for ok {
			randomCoords := b.RandomCoords()
			_, ok = coords[randomCoords]
		}
		coords[randomCoords] = true
		b.SetCell(randomCoords, val)
	}
}
