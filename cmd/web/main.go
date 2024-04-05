package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/dhinogz/requestwatcher/db"
	"github.com/dhinogz/requestwatcher/manager"
	"github.com/dhinogz/requestwatcher/server"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type config struct {
	port  string
	env   string
	dbDsn string
}

func loadConfig() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, err
	}
	var cfg config
	cfg.port = os.Getenv("PORT")
	cfg.env = os.Getenv("ENV")
	cfg.dbDsn = os.Getenv("PSQL_DSN")

	return cfg, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := loadConfig()
	if err != nil {
		logger.Error("could not load env variables, will exit", "err", err)
		os.Exit(1)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.dbDsn)
	if err != nil {
		logger.Error("could not connect to database, will exit", "err", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	q := db.New(conn)

	manager := manager.New()

	svr := server.New(
		server.WithPort(cfg.port),
		server.WithLogger(logger),
		server.WithStore(q),
		server.WithManager(manager),
	)

	logger.Info("starting server", "port", cfg.port, "env", cfg.env)
	log.Fatal(svr.ListenAndServe())
}
