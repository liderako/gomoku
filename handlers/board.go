package handlers

import (
	"encoding/json"
	"fmt"
	"gomoku/board"
	"io/ioutil"
	"log"
	"net/http"
)

func Board(_ http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	brd := board.Board{}
	err = json.Unmarshal(body, &brd)
	if err != nil {
		fmt.Println("Invalid board sent")
	} else {
		//threat := []heuristic.Threat{}
		fmt.Println("START")
		//fmt.Println(heuristic.SearchThreatComplexFourRowClose(brd, threat))
		//fmt.Println(heuristic.SearchThreatRowOpen(brd, threat, 5))
		//fmt.Println(game.GetCaptures(brd, board.Coords{2,7}))
		//fmt.Println(heuristic.IsTerminate(brd, 0, 0))
		//c := board.Coords{6, 7}
		//fmt.Println(heuristic.IsCorrectMove(brd, 1, c))
	}
}
