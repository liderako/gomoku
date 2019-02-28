package handlers

import (
	"encoding/json"
	"gomoku/game"
	"gomoku/minimax"
	"gomoku/minimax/heuristic"
	"net/http"
	"math"
	"strconv"
)

var difficulties = [...]struct {
	width, depth int
}{
	{width: 3, depth: 5},
	{width: 4, depth: 6},
	{width: 4, depth: 7},
}

func SuggestMoves(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	difficultyIndex, err := strconv.Atoi(req.URL.Query().Get("difficulty"))
	if err != nil || difficultyIndex < 0 || difficultyIndex >= len(difficulties) ||
	json.NewDecoder(req.Body).Decode(&state) != nil || state.Winner != 0 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	difficulty := difficulties[difficultyIndex]
	moves := minimax.Minimax(state, difficulty.width,
		difficulty.depth, heuristic.Evaluation, math.MinInt64, math.MaxInt64)
	resBody, err := json.Marshal(moves)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(resBody)
}
