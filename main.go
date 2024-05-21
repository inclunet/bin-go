package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/bingo"
	"github.com/inclunet/bin-go/pkg/braille"
	"github.com/inclunet/bin-go/pkg/server"
)

func main() {

	b := bingo.NewBingo()
	brl, err := braille.New("classes.json")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/bingo/{round}/new/{type}", b.AddRoundsHandler)
	r.HandleFunc("/api/bingo/{round}", b.GetRoundsHandler)
	r.HandleFunc("/api/bingo/{round}/0", b.AddCardsHandler)
	r.HandleFunc("/api/bingo/{round}/{card}", b.GetCardsHandler)
	r.HandleFunc("/api/bingo/{round}/{card}/autoplay", b.ToggleCardsAutoplayHandler)
	r.HandleFunc("/api/bingo/{round}/1/0", b.DrawHandler)
	r.HandleFunc("/ws/bingo/{round}/{card}", b.LiveHandler)
	r.HandleFunc("/api/bingo/{round}/{card}/{number}", b.ToggleNumbersHandler)
	r.HandleFunc("/qr/{round}/{card}", b.GetCardsQRHandler)
	//braille routs;
	r.HandleFunc("/api/braille/new", brl.AddPlayerHandler)
	r.HandleFunc("/api/braille/{player}", brl.GetPlayerHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/braille/{player}", brl.CheckChallengeRepplyHandler).Methods(http.MethodPost)

	server.AddFileServer(r)
	server.Start(r)
}
