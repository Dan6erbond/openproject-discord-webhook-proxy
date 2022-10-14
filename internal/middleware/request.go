package middleware

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/services"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RegisterRequestMiddleware(lc fx.Lifecycle, logger *zap.Logger, router *mux.Router, storageService services.StorageService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						logger.Sugar().Error(err)
					}
					err = storageService.SaveRequest(body, r.URL.Path)
					if err != nil {
						logger.Sugar().Error(err)
					}
					r.Body = io.NopCloser(bytes.NewBuffer(body))
					next.ServeHTTP(w, r)
				})
			})
			return nil
		},
	})
}
