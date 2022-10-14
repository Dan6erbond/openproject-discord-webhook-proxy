package pkg

import (
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/controllers"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/middleware"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/routes"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewMux,
			services.NewStorageService,
			services.NewOpenProjectService,
			services.NewWebhookService,
			services.NewDiscordService,
			controllers.NewWebhookController,
		),
		fx.Invoke(
			routes.RegisterOpenProjectRoutes,
			middleware.RegisterRequestMiddleware,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	return app
}
