package services

import (
	"fmt"
	"log"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/config"
	"github.com/spf13/viper"
)

type WebhookService struct {
	logger *log.Logger
}

func (ws *WebhookService) GetWebhook(webhookName string) (config.Webhook, error) {
	var webhooks []config.Webhook
	err := viper.UnmarshalKey("webhooks", &webhooks)
	if err != nil {
		return config.Webhook{}, err
	}
	var webhook config.Webhook
	found := false
	for _, w := range webhooks {
		if (w.Path != "" && w.Path == webhookName) || w.Name == webhookName {
			webhook = w
			found = true
			break
		}
	}

	if !found {
		return webhook, fmt.Errorf("couldn't find matching webhook for name %s", webhookName)
	}

	return webhook, nil
}

func NewWebhookService(logger *log.Logger) *WebhookService {
	logger.Print("Executing NewWebhookService.")
	ops := WebhookService{logger}
	return &ops
}
