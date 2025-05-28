package main

import (
	"github.com/VoltealProductions/the-royal-archives/internal/db"
	"github.com/VoltealProductions/the-royal-archives/internal/env"
	"go.uber.org/zap"
)

const version = "0.1.0"

func main() {
	cfg := config{
		addr:        env.GetString("ADDR", ":3030"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:3030"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:2929"),
		env:         env.GetString("ENV", "development"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Main DB
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established.")

	// Cache

	// Rate Limiter

	// Mailer

	// Authenticator

	app := &app{
		config: cfg,
		logger: logger,
	}

	// Metrics collected

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
