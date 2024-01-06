package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/legle/bin-go/pkg/bingo"
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

	r := mux.NewRouter()
	r.HandleFunc("/api/card/{round}/new/{type}", b.AddRound)
	r.HandleFunc("/api/card/{round}", b.GetRound)
	r.HandleFunc("/api/card/{round}/0", b.AddBingoCard)
	r.HandleFunc("/api/card/{round}/{card}", b.GetBingoCard)
	r.HandleFunc("/api/card/{round}/{card}/autoplay", b.ToggleAutoplay)
	r.HandleFunc("/api/card/{round}/1/0", b.Draw)
	r.HandleFunc("/api/card/{round}/{card}/{number}", b.ToggleNumber)
	r.HandleFunc("/qr/{round}/{card}", b.GetCardQR)
	r.PathPrefix("/card/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, dir+"index.html") })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(host+":"+port, r)

	if err != nil {
		panic(err)
	}
}
