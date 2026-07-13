package database

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultMaxConnections int32 = 5

//go:embed migrations/*.sql
var migrationFiles embed.FS

// Open creates a PostgreSQL pool from DATABASE_URL. A missing URL keeps the
// application in its legacy in-memory mode, which is useful for local usage.
func Open(ctx context.Context) (*pgxpool.Pool, error) {
	databaseURL := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseURL == "" {
		if strings.EqualFold(strings.TrimSpace(os.Getenv("REQUIRE_DATABASE")), "true") {
			return nil, errors.New("DATABASE_URL is required")
		}
		return nil, nil
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse DATABASE_URL: %w", err)
	}

	config.MinConns = 0
	config.MaxConns = getMaxConnections()
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	if err := migrate(ctx, pool); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func getMaxConnections() int32 {
	value := strings.TrimSpace(os.Getenv("DB_MAX_CONNS"))
	if value == "" {
		return defaultMaxConnections
	}

	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil || parsed < 1 {
		return defaultMaxConnections
	}

	return int32(parsed)
}

func migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if pool == nil {
		return errors.New("database pool is required")
	}

	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read database migrations: %w", err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	connection, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquire migration connection: %w", err)
	}
	defer connection.Release()

	if _, err := connection.Exec(ctx, `SELECT pg_advisory_lock(731946210)`); err != nil {
		return fmt.Errorf("lock database migrations: %w", err)
	}
	defer connection.Exec(context.Background(), `SELECT pg_advisory_unlock(731946210)`) //nolint:errcheck

	if _, err := connection.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`); err != nil {
		return fmt.Errorf("create schema migrations table: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		var alreadyApplied bool
		err := connection.QueryRow(
			ctx,
			`SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`,
			entry.Name(),
		).Scan(&alreadyApplied)
		if err != nil {
			return fmt.Errorf("check migration %s: %w", entry.Name(), err)
		}
		if alreadyApplied {
			continue
		}

		sql, err := migrationFiles.ReadFile("migrations/" + entry.Name())
		if err != nil {
			return fmt.Errorf("read migration %s: %w", entry.Name(), err)
		}

		tx, err := connection.Begin(ctx)
		if err != nil {
			return fmt.Errorf("begin migration %s: %w", entry.Name(), err)
		}

		if _, err := tx.Exec(ctx, string(sql)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("apply migration %s: %w", entry.Name(), err)
		}
		if _, err := tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, entry.Name()); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("record migration %s: %w", entry.Name(), err)
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("commit migration %s: %w", entry.Name(), err)
		}
	}

	return nil
}
