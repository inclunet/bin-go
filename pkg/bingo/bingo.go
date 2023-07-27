package bingo

import (
	"encoding/json"
	"image/png"
	"net/http"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
)

type Bingo struct {
	Rounds []Round
}

func (b *Bingo) AddRound(w http.ResponseWriter, r *http.Request) {
	round := NewRound(len(b.Rounds)+1, GetUrlIntParam(r, "type"))
	card := round.AddCard()
	b.Rounds = append(b.Rounds, round)
	Repply(w, card)
}

func (b *Bingo) CheckNumber(w http.ResponseWriter, r *http.Request) {
	roundID := GetUrlIntParam(r, "round") - 1
	cardID := GetUrlIntParam(r, "card") - 1
	number := GetUrlIntParam(r, "number")
	if b.Rounds[roundID].Cards[0].IsChecked(number) || cardID == 0 {
		Repply(w, b.Rounds[roundID].Cards[cardID].CheckNumber(number))
	} else {
		Repply(w, b.Rounds[roundID].Cards[cardID])
	}
}

func (b *Bingo) AddBingoCard(w http.ResponseWriter, r *http.Request) {
	card := b.Rounds[GetUrlIntParam(r, "round")-1].AddCard()
	Repply(w, card)
}

func (b *Bingo) GetBingoCard(w http.ResponseWriter, r *http.Request) {
	round := b.Rounds[GetUrlIntParam(r, "round")-1]
	Repply(w, round.GetBingoCard(GetUrlIntParam(r, "card")))
}

func (b *Bingo) GetRound(w http.ResponseWriter, r *http.Request) {
	round := b.Rounds[GetUrlIntParam(r, "round")-1]
	Repply(w, round)
}

func (b *Bingo) GetCardQR(w http.ResponseWriter, r *http.Request) {
	qr, _ := qr.Encode(GetQueryString(r, "url", ""), qr.L, qr.Auto)
	qrCode, _ := barcode.Scale(qr, 300, 300)
	png.Encode(w, qrCode)
}

func (b *Bingo) uncheckNumbers(w http.ResponseWriter, r *http.Request) {
	roundID := GetUrlIntParam(r, "round") - 1
	cardID := GetUrlIntParam(r, "card") - 1
	number := GetUrlIntParam(r, "number")
	if b.Rounds[roundID].Cards[0].IsChecked(number) && cardID == 0 {
		Repply(w, b.Rounds[roundID].uncheckNumber(number).Cards[cardID])
	} else {
		Repply(w, b.Rounds[roundID].Cards[cardID])
	}
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
