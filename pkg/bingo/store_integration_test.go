package bingo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
	connection, err := pool.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, migrationErr := connection.Conn().PgConn().Exec(ctx, string(migration)).ReadAll()
	connection.Release()
	if migrationErr != nil {
		t.Fatal(migrationErr)
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

	roundID, err := uuid.Parse(round.ID)
	if err != nil {
		t.Fatal(err)
	}
	expectedDrawTime := time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC)
	if _, err := pool.Exec(ctx, `
		UPDATE bingo_draws
		SET created_at = $1
		WHERE round_id = $2 AND number = $3
	`, expectedDrawTime, roundID, drawn); err != nil {
		t.Fatal(err)
	}
	if err := store.SaveRound(ctx, &round); err != nil {
		t.Fatal(err)
	}
	var persistedDrawTime time.Time
	if err := pool.QueryRow(ctx, `
		SELECT created_at
		FROM bingo_draws
		WHERE round_id = $1 AND number = $2
	`, roundID, drawn).Scan(&persistedDrawTime); err != nil {
		t.Fatal(err)
	}
	if !persistedDrawTime.Equal(expectedDrawTime) {
		t.Fatalf("draw timestamp changed to %s, want %s", persistedDrawTime, expectedDrawTime)
	}

	loaded, err := store.LoadRounds(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(loaded) != 1 {
		t.Fatalf("loaded %d rounds, want 1", len(loaded))
	}
	if len(loaded[0].Cards) != 2 {
		t.Fatalf("loaded %d cards, want 2", len(loaded[0].Cards))
	}
	if loaded[0].ID != round.ID || loaded[0].Cards[1].ID != round.Cards[1].ID {
		t.Fatal("round or card UUID did not survive the database round trip")
	}
	if !loaded[0].Cards[0].IsChecked(drawn) {
		t.Fatalf("drawn number %d was not restored", drawn)
	}
}
