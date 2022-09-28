package middleware

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/services"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func RegisterRequestMiddleware(lc fx.Lifecycle, logger *log.Logger, router *mux.Router, storageService services.StorageService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						logger.Println(err)
					}
					err = storageService.SaveRequest(body, r.URL.Path)
					if err != nil {
						logger.Println(err)
					}
					r.Body = io.NopCloser(bytes.NewBuffer(body))
					next.ServeHTTP(w, r)
				})
			})
			return nil
		},
	})
}
