package bingo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Bingo struct {
	Rounds []Round
}

func (b *Bingo) AddRound(w http.ResponseWriter, r *http.Request) {
	round := NewRound(len(b.Rounds)+1, GetUrlIntParam(r, "type"))
	card := round.AddBingoCard()
	b.Rounds = append(b.Rounds, round)
	Repply(w, card)
}

func (b *Bingo) CheckNumber(w http.ResponseWriter, r *http.Request) {
	card := b.Rounds[GetUrlIntParam(r, "round")-1].Cards[GetUrlIntParam(r, "card")-1].CheckNumber(GetUrlIntParam(r, "number"))
	Repply(w, card)
}

func (b *Bingo) AddBingoCard(w http.ResponseWriter, r *http.Request) {
	card := b.Rounds[GetUrlIntParam(r, "round")-1].AddBingoCard()
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

func GetUrlIntParam(r *http.Request, param string) int {
	if intParam, err := strconv.Atoi(GetUrlStringParam(r, param)); err == nil {
		return intParam
	}

	return 0
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
