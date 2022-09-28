package pkg

import (
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/controllers"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/routes"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(
			NewLogger,
			services.NewOpenProjectService,
			services.NewWebhookService,
			services.NewDiscordService,
			controllers.NewWebhookController,
			NewMux,
		),
		fx.Invoke(routes.RegisterOpenProjectRoutes),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	return app
}
