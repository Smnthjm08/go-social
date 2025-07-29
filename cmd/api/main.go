// starting point

package main

import (
	"log"
	"os"
	"time"

	"github.com/smnthjm08/go-social/internal/db"
	"github.com/smnthjm08/go-social/internal/env"
	"github.com/smnthjm08/go-social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			go-social API
//	@description	API for the GopherSocial, a social network built on go
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@BasePath	/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {
	cfg := config{
		// addr: ":8000",
		addr:   env.GetString("ADDR", ":8000"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3,
		},
	}

	// logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		// log.Panic(err)
		logger.Fatal(err)
	}

	defer db.Close()
	log.Printf("db connected...")
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
