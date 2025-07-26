package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/smnthjm08/go-social/internal/store"

	"github.com/smnthjm08/go-social/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr   string
	db     dbConfig
	env    string
	apiURL string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// request timeout
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		// docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		// r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/v1/swagger/doc.json")))

		// v1/posts
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)

				r.Get("/", app.getPostByIDHandler)
				r.Delete("/", app.deletePostByIDHandler)
				r.Put("/", app.updatePostByIDHandler)
			})

		})

		r.Route("/user", func(r chi.Router) {

			r.Route("/{userId}", func(r chi.Router) {
				r.Use(app.userContextMiddleware)

				r.Get("/", app.getUserHandler)
				// r.Delete("/", app.deletePostByIDHandler)
				// r.Put("/", app.updatePostByIDHandler)

				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
		})

		// v1/users/feed

	})

	return r

	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	// return mux

	// posts
	// users
	// auth

}

func (app *application) run(mux http.Handler) error {
	// mux := http.NewServeMux()
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return server.ListenAndServe()
}
