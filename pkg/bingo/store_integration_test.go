package bingo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestPostgresStoreRoundTrip(t *testing.T) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("TEST_DATABASE_URL is not configured")
	}

	ctx := context.Background()
	admin, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	defer admin.Close()

	schema := "bingo_test_" + strings.ReplaceAll(uuid.NewString(), "-", "")
	if _, err := admin.Exec(ctx, fmt.Sprintf(`CREATE SCHEMA "%s"`, schema)); err != nil {
		t.Fatal(err)
	}
	defer admin.Exec(ctx, fmt.Sprintf(`DROP SCHEMA "%s" CASCADE`, schema)) //nolint:errcheck

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	config.ConnConfig.RuntimeParams["search_path"] = schema
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	migration, err := os.ReadFile(filepath.Join("..", "database", "migrations", "001_bingo.sql"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, string(migration)); err != nil {
		t.Fatal(err)
	}

	round := NewRound(&Bingo{}, 75)
	if _, err := round.AddCard(); err != nil {
		t.Fatal(err)
	}
	drawn := round.Draw().LastNumber
	round.ToggleNumberForAll(drawn)

	store := NewPostgresStore(pool)
	if err := store.SaveRound(ctx, &round); err != nil {
		t.Fatal(err)
	}
	loaded, err := store.LoadRounds(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(loaded) != 1 || len(loaded[0].Cards) != 2 {
		t.Fatalf("loaded %d rounds and %d cards", len(loaded), len(loaded[0].Cards))
	}
	if loaded[0].ID != round.ID || loaded[0].Cards[1].ID != round.Cards[1].ID {
		t.Fatal("round or card UUID did not survive the database round trip")
	}
	if !loaded[0].Cards[0].IsChecked(drawn) {
		t.Fatalf("drawn number %d was not restored", drawn)
	}
}
