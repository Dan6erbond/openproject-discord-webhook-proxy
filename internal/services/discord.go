package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/config"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/discord"
)

type DiscordService struct {
	logger *log.Logger
}

func (ds *DiscordService) SendWebhook(webhook config.Webhook, webhookPayload discord.Webhook) error {
	content, err := json.Marshal(webhookPayload)
	if err != nil {
		return err
	}

	contentReader := bytes.NewReader(content)
	resp, err := http.Post(webhook.URL, "application/json", contentReader)
	if err != nil {
		return err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(respBody) != "" {
		return fmt.Errorf(string(respBody))
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		ds.logger.Println(resp.StatusCode)
		return fmt.Errorf("discord error")
	}

	return nil
}

func NewDiscordService(logger *log.Logger) *DiscordService {
	logger.Print("Executing NewDiscordService.")
	ds := DiscordService{logger}
	return &ds
}
