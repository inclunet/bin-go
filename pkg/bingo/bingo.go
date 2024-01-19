package bingo

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
)

type Bingo struct {
	Rounds []Round
}

func (b *Bingo) AddCardsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Adding bingo card for round %d...", GetUrlIntParam(r, "round"))

	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	card := round.AddCard()

	log.Printf("Added bingo card %d for round %d", card.Card, GetUrlIntParam(r, "round"))

	Repply(w, card)
}

func (b *Bingo) AddRoundsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding round", GetUrlIntParam(r, "round"), GetUrlIntParam(r, "type"))

	round := NewRound(*b, GetUrlIntParam(r, "type"))

	if GetUrlIntParam(r, "round") > 0 {
		log.Printf("Forwarding %d round players to the %d new round...", GetUrlIntParam(r, "round"), round.Round)

		oldRound, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

		if err != nil {
			log.Printf("Error on get old round: %v", err)
		}

		if players := oldRound.SetNextRound(round.Round); players > 0 {
			log.Printf("Forwarded %d players from round %d to %d new round.", players, oldRound.Round, round.Round)
		}
	}

	card, err := round.GetCard(0)

	if err != nil {
		log.Printf("Error on get main card: %v", err)
	}

	b.Rounds = append(b.Rounds, round)

	log.Printf("Added round %d", round.Round)

	Repply(w, card)
}

func (b *Bingo) DrawHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Drawing number for round %d...", GetUrlIntParam(r, "round"))

	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	card, err := round.GetCard(0)

	if err != nil {
		log.Printf("Error on get card: %v", err)
	}

	number := card.Draw()

	log.Printf("Drawed number %d for round %d", number, GetUrlIntParam(r, "round"))

	checked := round.CheckNumberForAll(number)

	log.Printf("Checked number %d for %d cards into round %d", number, checked, round.Round)

	Repply(w, card)
}

func (b *Bingo) GetCardsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting bingo card %d for round %d...", GetUrlIntParam(r, "card"), GetUrlIntParam(r, "round"))

	Round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	card, err := Round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		log.Printf("Error on get card: %v", err)
	}

	log.Printf("Got bingo card %d for round %d", card.Card, Round.Round)

	Repply(w, card)
}

func (b *Bingo) GetCardsQRHandler(w http.ResponseWriter, r *http.Request) {
	qr, _ := qr.Encode(GetQueryString(r, "url", ""), qr.L, qr.Auto)
	qrCode, _ := barcode.Scale(qr, 300, 300)
	png.Encode(w, qrCode)
}

func (b *Bingo) GetRound(round int) (*Round, error) {
	if round < 0 || round >= len(b.Rounds) || len(b.Rounds) == 0 {
		return nil, fmt.Errorf("Round %d not found", round)
	}

	return &b.Rounds[round], nil
}

func (b *Bingo) GetRoundsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting round %d...", GetUrlIntParam(r, "round"))

	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	Repply(w, round)
}

func (b *Bingo) LiveHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting live for round %d...", GetUrlIntParam(r, "round"))

	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	card, err := round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		log.Printf("Error on get card: %v", err)
	}

	round.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := round.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Error on upgrade connection: %v", err)
	}

	if !card.SetConn(conn) {
		log.Printf("Error on set connection: %v", err)
	}

	card.UpdateCard()
}

func (b *Bingo) ToggleCardsAutoplayHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Toggling autoplay for card %d on round %d...", GetUrlIntParam(r, "card"), GetUrlIntParam(r, "round"))

	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	card, err := round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		log.Printf("Error on get card: %v", err)
	}

	card.ToggleAutoplay()

	Repply(w, card)
}

func (b *Bingo) ToggleNumbersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Toggling number %d for card %d on round %d...", GetUrlIntParam(r, "number"), GetUrlIntParam(r, "card"), GetUrlIntParam(r, "round"))

	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		log.Printf("Error on get round: %v", err)
	}

	mainCard, err := round.GetCard(0)

	if err != nil {
		log.Printf("Error on get main card: %v", err)
	}

	card, err := round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		log.Printf("Error on get card: %v", err)
	}

	if card.ToggleNumber(*mainCard, GetUrlIntParam(r, "number")) {
		log.Printf("Number %d for card %d on round %d toggled", GetUrlIntParam(r, "number"), GetUrlIntParam(r, "card"), GetUrlIntParam(r, "round"))

		if card.Card == 1 {
			checked, unchecked := round.ToggleNumberForAll(GetUrlIntParam(r, "number"))

			log.Printf("Number %d checked for %d cards and unchecked for %d cards into round %d", GetUrlIntParam(r, "number"), checked, unchecked, round.Round)
		}
	}

	Repply(w, card)
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
