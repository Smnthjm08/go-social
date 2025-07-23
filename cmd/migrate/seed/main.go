package main

import (
	"log"

	"github.com/smnthjm08/go-social/internal/db"
	"github.com/smnthjm08/go-social/internal/env"
	"github.com/smnthjm08/go-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
