package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/legle/bin-go/pkg/bingo"
)

func main() {
	b := bingo.NewBingo()
	r := mux.NewRouter()
	r.HandleFunc("/api/card/0/{type}", b.AddRound)
	r.HandleFunc("/api/card/{round}", b.GetRound)
	r.HandleFunc("/api/card/{round}/0", b.AddBingoCard)
	r.HandleFunc("/api/card/{round}/{card}", b.GetBingoCard)
	r.HandleFunc("/api/card/{round}/{card}/{number}", b.CheckNumber)
	http.Handle("/app/", http.FileServer(http.Dir("./bingo/build")))
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
