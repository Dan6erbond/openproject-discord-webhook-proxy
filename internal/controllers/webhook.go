package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/openproject"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type WebhookController struct {
	logger             *zap.Logger
	openProjectService *services.OpenProjectService
	webhookService     *services.WebhookService
	discordService     *services.DiscordService
}

func (wc *WebhookController) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wc.logger.Info("Received payload for webhook", zap.String("name", vars["name"]))
	webhook, err := wc.webhookService.GetWebhook(vars["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wc.logger.Error("Error reading body", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/* err = wc.openProjectService.ValidateSignature(body, webhook, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} */

	var payload openproject.Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		wc.logger.Error("Error unmarshaling JSON", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	run := false
	if len(webhook.Actions) > 0 {
		for _, action := range webhook.Actions {
			if payload.Action == action {
				run = true
			}
		}
	} else {
		run = true
	}

	if run {
		switch payload.Action {
		case "work_package:created", "work_package:updated":
			var payload openproject.WorkPackageWebhookPayload
			err := json.Unmarshal(body, &payload)
			if err != nil {
				wc.logger.Error("Error unmarshaling JSON", zap.String("error", err.Error()))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			webhookPayload, err := wc.openProjectService.GetWorkPackagePayload(payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = wc.discordService.SendWebhook(webhook, webhookPayload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			wc.logger.Info("Couldn't find handler for webhook", zap.String("webhook", vars["name"]), zap.String("action", payload.Action))
			http.Error(w, "Couldn't find action for webhook", http.StatusNotFound)
		}
	}
}

func NewWebhookController(logger *zap.Logger, openProjectService *services.OpenProjectService, webhookService *services.WebhookService, discordService *services.DiscordService) *WebhookController {
	logger.Info("Executing NewWebhookController.")
	ops := WebhookController{logger, openProjectService, webhookService, discordService}
	return &ops
}
