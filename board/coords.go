package board

type Coords struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coords) AreWithin() bool {
	return c.X >= 0 && c.X < 19 && c.Y >= 0 && c.Y < 19
}

func (c Coords) Neighbors(r int8) []Coords {
	neighbors := make([]Coords, 0)
	for y := -2; y <= 2; y++ {
		for x := -2; x <= 2; x++ {
			if x == 0 && y == 0 {
				continue
			}
			neighbCoords := Coords{
				X: c.X + x,
				Y: c.Y + y,
			}
			if neighbCoords.AreWithin() {
				neighbors = append(neighbors, neighbCoords)
			}
		}
	}
	return neighbors
}

func DiffCoords(expected, received []Coords) bool {
	if len(expected) != len(received) {
		return false
	}
	expMap := make(map[Coords]bool)
	for _, x := range expected {
		expMap[x] = true
	}
	for _, x := range received {
		if !expMap[x] {
			return false
		}
	}
	return true
}
