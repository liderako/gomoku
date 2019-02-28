package heuristic

import (
	"gomoku/board"
)

const (
	FiveROW        = 10000000000000 // # # # # #
	TwoRowCloseWIN = FiveROW
	FourRowOpen     = 100000000000 // # # # #
	FourRowClose    = 100000000    // * # # # #
	ThreeRowOpen   = 1000000      // # # #
	ThreeRowClose  = 10000        // * # # #
	TwoRowClose    = 1000         // * # #
	TwoRowOpen     = 10       // # #
	TwoRowCloseSix = ThreeRowClose * 2
	TwoRowCloseFour = TwoRowClose + TwoRowClose
)

func EvaluationRate(threat []Threat, amountPointMinPlayer int8, amountPointMaxPlayer int8) []Threat {
	for key, value := range threat {
		if value.size == 5 {
			if value.owner == board.WHITE_PLAYER {
				threat[key].rate = FiveROW
			} else {
				threat[key].rate = -FiveROW
			}
		} else if value.size == 4 {
			if value.status == OPEN_THREAT {
				if value.owner == board.WHITE_PLAYER {
					threat[key].rate = FourRowOpen
				} else {
					threat[key].rate = -FourRowOpen
				}
			} else {
				if value.owner == board.WHITE_PLAYER {
					threat[key].rate = FourRowClose
				} else {
					threat[key].rate = -FourRowClose
				}
			}
		} else if value.size == 3 {
			if value.status == OPEN_THREAT {
				if value.owner == board.WHITE_PLAYER {
					threat[key].rate = ThreeRowOpen
				} else {
					threat[key].rate = -ThreeRowOpen
				}
			} else {
				if value.owner == board.WHITE_PLAYER {
					threat[key].rate = ThreeRowClose
				} else {
					threat[key].rate = -ThreeRowClose
				}
			}
		} else if value.size == 2 {
			if value.status == OPEN_THREAT {
				if value.owner == board.WHITE_PLAYER {
					threat[key].rate = TwoRowOpen
				} else {
					threat[key].rate = -TwoRowOpen
				}
			} else {
				if value.owner == board.BLACK_PLAYER {
					if amountPointMaxPlayer >= 8 {
						threat[key].rate = TwoRowCloseWIN
					} else if amountPointMaxPlayer == 6 {
						threat[key].rate = TwoRowCloseSix
					} else if amountPointMaxPlayer == 4{
						threat[key].rate = TwoRowCloseFour
					} else {
						threat[key].rate = TwoRowClose
					}
				} else {
					if amountPointMinPlayer >= 8 {
						threat[key].rate = -TwoRowCloseWIN
					} else if amountPointMinPlayer == 6 {
						threat[key].rate = -TwoRowCloseSix
					} else if amountPointMinPlayer == 4 {
						threat[key].rate = -TwoRowCloseFour
					} else {
						threat[key].rate = -TwoRowClose
					}
				}
			}
		}
	}
	return threat
}

func Evaluation(brd board.Board, amountPointMinPlayer int8, amountPointMaxPlayer int8) int64 {
	var power int64 = 0

	i := 5
	threat := []Threat{}
	for i >= 2 {
		threat = SearchThreatRowClose(brd, threat, i)
		threat = SearchThreatRowOpen(brd, threat, i)
		i--
	}
	threat = SearchThreatComplexFourRowOpen(brd, threat)
	threat = SearchThreatComplexFourRowClose(brd, threat)
	threat = EvaluationRate(threat, amountPointMinPlayer, amountPointMaxPlayer)
	for _, value := range threat {
		power = power + value.rate
	}
	return power
}
