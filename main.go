package main

import (
	"github.com/rs/cors"
	"gomoku/handlers"
	"log"
	"net/http"
	"flag"
	"os"
	"runtime/pprof"
	"time"
)

const PORT = "0.0.0.0:5555"

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		go func() {
			time.Sleep(time.Second * 30)
			pprof.StopCPUProfile()
			log.Println("profiling stopped")
		}()
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/board", handlers.Board)
	mux.HandleFunc("/make-move", handlers.MakeMove)
	mux.HandleFunc("/suggest-moves", handlers.SuggestMoves)

	println("Starting server at " + PORT)
	log.Fatalln(http.ListenAndServe(PORT, cors.Default().Handler(mux)))
}
