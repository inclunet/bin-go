package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/bingo"
	"github.com/inclunet/bin-go/pkg/braille"
)

func main() {
	var port string
	var host string
	var dir string

	flag.StringVar(&port, "port", "80", "Port to listen on")
	flag.StringVar(&host, "host", "", "Host to listen on")
	flag.StringVar(&dir, "dir", "./", "Directory to serve")
	flag.Parse()

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	b := bingo.NewBingo()
	brl, err := braille.New("classes.json")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/card/{round}/new/{type}", b.AddRoundsHandler)
	r.HandleFunc("/api/card/{round}", b.GetRoundsHandler)
	r.HandleFunc("/api/card/{round}/0", b.AddCardsHandler)
	r.HandleFunc("/api/card/{round}/{card}", b.GetCardsHandler)
	r.HandleFunc("/api/card/{round}/{card}/autoplay", b.ToggleCardsAutoplayHandler)
	r.HandleFunc("/api/card/{round}/1/0", b.DrawHandler)
	r.HandleFunc("/api/card/{round}/{card}/live", b.LiveHandler)
	r.HandleFunc("/api/card/{round}/{card}/{number}", b.ToggleNumbersHandler)
	r.HandleFunc("/qr/{round}/{card}", b.GetCardsQRHandler)
	//braille routs;
	r.HandleFunc("/api/braille/new", brl.AddPlayerHandler)
	r.HandleFunc("/api/braille/{player}", brl.GetPlayerHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/braille/{player}", brl.CheckChallengeRepplyHandler).Methods(http.MethodPost)

	r.PathPrefix("/card/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, dir+"index.html") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
	err = http.ListenAndServe(host+":"+port, r)

	if err != nil {
		panic(err)
	}
}
