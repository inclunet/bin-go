package main

import (
	"context"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/bingo"
	"github.com/inclunet/bin-go/pkg/braille"
	"github.com/inclunet/bin-go/pkg/database"
	"github.com/inclunet/bin-go/pkg/server"
	"github.com/inclunet/bin-go/pkg/tictac"
)

func main() {
	server.Logger.Info("Starting server...")

	r := mux.NewRouter().StrictSlash(true)
	r.Use(server.MetricsAndLoggingMiddleware)
	api := r.PathPrefix("/api").Subrouter()
	qr := r.PathPrefix("/qr").Subrouter()
	ws := r.PathPrefix("/ws").Subrouter()

	server.Logger.Info("Adding metrics server...")

	server.AddMetricsServer(r)

	server.Logger.Info("Adding Bingo routes...")

	pool, err := database.Open(context.Background())
	if err != nil {
		if strings.EqualFold(os.Getenv("REQUIRE_DATABASE"), "true") {
			server.Logger.Error("Database startup failed", "error", err)
			os.Exit(1)
		}
		server.Logger.Warn("Database startup failed; bingo persistence is disabled", "error", err)
		pool = nil
	}
	if pool != nil {
		defer pool.Close()
		server.Logger.Info("PostgreSQL persistence enabled")
	} else if err == nil {
		server.Logger.Warn("DATABASE_URL is not configured; bingo persistence is disabled")
	}

	var bingoStore bingo.Store
	if pool != nil {
		bingoStore = bingo.NewPostgresStore(pool)
	}

	bingoGame, err := bingo.NewWithStore(api, bingoStore)
	if err != nil {
		server.Logger.Error("Bingo startup failed", "error", err)
		os.Exit(1)
	}
	bingoGame.AddQrRoutes(qr).AddWsRoutes(ws)

	server.Logger.Info("Adding Braille routes...")

	_, err = braille.New(api)

	if err != nil {
		server.Logger.Error(err.Error())
	}

	server.Logger.Info("Adding file server...")

	server.Logger.Info("Adding TicTac routes...")
	tictac.New(api).AddWsRoutes(ws)

	server.AddFileServer(r)

	server.Logger.Info("Starting server...")

	server.Start(r)
}
