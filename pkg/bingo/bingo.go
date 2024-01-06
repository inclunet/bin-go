package bingo

import (
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
)

type Bingo struct {
	Rounds Rounds
}

func (b *Bingo) AddRound(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding round", GetUrlIntParam(r, "round")-1, GetUrlIntParam(r, "type"))
	log.Println(b)
	currentRound, err := b.Rounds.AddRound(GetUrlIntParam(r, "round")-1, GetUrlIntParam(r, "type"))

	if err != nil {
		log.Printf("Adding card: %s", err)
	}

	Repply(w, currentRound)
}

func (b *Bingo) Draw(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.Rounds.Draw(GetUrlIntParam(r, "round")-1))
}

func (b *Bingo) ToggleNumber(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.Rounds.ToggleNumber(GetUrlIntParam(r, "round")-1, GetUrlIntParam(r, "card")-1, GetUrlIntParam(r, "number")))
}

func (b *Bingo) AddBingoCard(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.Rounds.AddCard(GetUrlIntParam(r, "round")-1))
}

func (b *Bingo) GetBingoCard(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.Rounds.GetCard(GetUrlIntParam(r, "round")-1, GetUrlIntParam(r, "card")-1))
}

func (b *Bingo) GetRound(w http.ResponseWriter, r *http.Request) {
	currentRound, err := b.Rounds.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Println(err)
	}

	Repply(w, currentRound)
}

func (b *Bingo) GetCardQR(w http.ResponseWriter, r *http.Request) {
	qr, _ := qr.Encode(GetQueryString(r, "url", ""), qr.L, qr.Auto)
	qrCode, _ := barcode.Scale(qr, 300, 300)
	png.Encode(w, qrCode)
}

func (b *Bingo) ToggleAutoplay(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.Rounds.ToggleAutoplay(GetUrlIntParam(r, "round")-1, GetUrlIntParam(r, "card")-1))
}

func (b *Bingo) UncheckNumbers(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.Rounds.UncheckNumber(GetUrlIntParam(r, "round")-1, GetUrlIntParam(r, "card")-1, GetUrlIntParam(r, "number")))
}

func GetUrlIntParam(r *http.Request, param string) int {
	if intParam, err := strconv.Atoi(GetUrlStringParam(r, param)); err == nil {
		return intParam
	}

	return 0
}

func GetQueryString(r *http.Request, key string, value string) string {
	query := r.URL.Query()

	if value, ok := query[key]; ok {
		return value[0]
	}

	return value
}

func GetUrlStringParam(r *http.Request, param string) string {
	vars := mux.Vars(r)
	return vars[param]
}

func Repply(w http.ResponseWriter, data any) {
	response, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func NewBingo() Bingo {
	return Bingo{}
}
