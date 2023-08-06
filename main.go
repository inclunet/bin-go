package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/legle/bin-go/pkg/bingo"
)

func main() {
	b := bingo.NewBingo()
	r := mux.NewRouter()
	r.HandleFunc("/api/card/{round}/new/{type}", b.AddRound)
	r.HandleFunc("/api/card/{round}", b.GetRound)
	r.HandleFunc("/api/card/{round}/0", b.AddBingoCard)
	r.HandleFunc("/api/card/{round}/{card}", b.GetBingoCard)
	r.HandleFunc("/api/card/{round}/{card}/autoplay", b.ToggleAutoplay)
	r.HandleFunc("/api/card/{round}/1/0", b.Draw)
	r.HandleFunc("/api/card/{round}/{card}/{number}", b.ToggleNumber)
	r.HandleFunc("/qr/{round}/{card}", b.GetCardQR)
	r.PathPrefix("/card/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./bingo/build/index.html") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./bingo/build/")))
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
