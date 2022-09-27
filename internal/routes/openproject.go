package routes

import (
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/controllers"
	"github.com/gorilla/mux"
)

func RegisterOpenProjectRoutes(router *mux.Router, wc *controllers.WebhookController) {
	router.HandleFunc("/webhooks/{name}", wc.HandleWebhook).Methods("POST")
}
