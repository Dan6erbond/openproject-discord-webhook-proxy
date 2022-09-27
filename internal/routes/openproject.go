package routes

import (
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/services"
	"github.com/gorilla/mux"
)

func RegisterOpenProjectRoutes(router *mux.Router, ops *services.OpenProjectService) {
	router.HandleFunc("/webhooks/{name}", ops.HandleWebhook).Methods("POST")
}
