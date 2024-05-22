package main

import (
	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/bingo"
	"github.com/inclunet/bin-go/pkg/braille"
	"github.com/inclunet/bin-go/pkg/server"
)

func main() {
	server.Logger.Info("Starting server...")

	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api").Subrouter()
	qr := r.PathPrefix("qr").Subrouter()
	ws := r.PathPrefix("ws").Subrouter()

	server.Logger.Info("Adding Bingo routes...")

	bingo.New(api).AddQrRoutes(qr).AddWsRoutes(ws)

	server.Logger.Info("Adding Braille routes...")

	_, err := braille.New(api)

	if err != nil {
		server.Logger.Error(err.Error())
	}

	server.Logger.Info("Adding file server...")

	server.AddFileServer(r)

	server.Logger.Info("Starting server...")

	server.Start(r)
}
