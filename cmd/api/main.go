// starting point

package main

import (
	"log"
	"os"

	"github.com/smnthjm08/go-social/internal/db"
	"github.com/smnthjm08/go-social/internal/env"
	"github.com/smnthjm08/go-social/internal/store"
)

func main() {
	cfg := config{
		// addr: ":8000",
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
