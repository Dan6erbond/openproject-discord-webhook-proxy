package public

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func NewMux(lc fx.Lifecycle, logger *log.Logger) *mux.Router {
	logger.Print("Executing NewMux.")

	r := mux.NewRouter()
	server := &http.Server{
		Addr:    ":5001",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return r
}
