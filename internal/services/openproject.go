package services

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/config"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/discord"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/openproject"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type OpenProjectService struct {
	logger *zap.Logger
}

func (ops *OpenProjectService) ValidateSignature(body []byte, webhook config.Webhook, r *http.Request) error {
	signature := r.Header["X-Op-Signature"]
	h := hmac.New(sha1.New, []byte(webhook.Secret))

	jsonString, err := json.Marshal(string(body))
	if err != nil {
		return err
	}

	h.Write(jsonString)

	if signature[0] != fmt.Sprintf("sha1=%x", h.Sum(nil)) {
		return fmt.Errorf("signatures don't match")
	}

	return nil
}

func (ops *OpenProjectService) GetWorkPackagePayload(payload openproject.WorkPackageWebhookPayload) (discord.Webhook, error) {
	color, err := strconv.ParseInt(payload.WorkPackage.Embedded.Type.Color[1:], 16, 64)
	if err != nil {
		return discord.Webhook{}, err
	}

	openProjectBaseUrl, err := url.Parse(viper.GetString("openproject.baseurl"))
	if err != nil {
		return discord.Webhook{}, err
	}

	projectUrl := url.URL{
		Scheme: openProjectBaseUrl.Scheme,
		Host:   openProjectBaseUrl.Host,
		Path:   fmt.Sprintf("/projects/%s", payload.WorkPackage.Embedded.Project.Identifier),
	}

	workPackageUrl := url.URL{
		Scheme: openProjectBaseUrl.Scheme,
		Host:   openProjectBaseUrl.Host,
		Path:   fmt.Sprintf("/projects/%s/work_packages/%d/activity", payload.WorkPackage.Embedded.Project.Identifier, payload.WorkPackage.ID),
	}

	webhookBuilder := discord.WebhookBuilder()

	ops.logger.Info("Building embed", zap.String("action", payload.Action))
	if payload.Action == "work_package:created" {
		webhookBuilder.Content("Work package created")
	} else {
		webhookBuilder.Content("Work package updated")
	}

	embedBuilder := discord.EmbedBuilder().
		Author(discord.Author{
			Name:    payload.WorkPackage.Embedded.Project.Name,
			IconURL: payload.WorkPackage.Embedded.Author.Avatar,
			URL:     projectUrl.String(),
		}).
		Color(color).
		Title(fmt.Sprintf("%s: %s", payload.WorkPackage.Embedded.Type.Name, payload.WorkPackage.Subject)).
		URL(workPackageUrl.String())
	if payload.WorkPackage.Description.Raw != "" {
		embedBuilder.
			Description(payload.WorkPackage.Description.Raw)
	}
	if startDate, ok := payload.WorkPackage.StartDate.(string); ok && startDate != "" {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  startDate,
			Inline: true,
		})
	}
	if derivedStartDate, ok := payload.WorkPackage.DerivedStartDate.(string); ok && derivedStartDate != "" {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  derivedStartDate,
			Inline: true,
		})
	}
	if dueDate, ok := payload.WorkPackage.DueDate.(string); ok && dueDate != "" {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  dueDate,
			Inline: true,
		})
	}
	if derivedDueDate, ok := payload.WorkPackage.DerivedDueDate.(string); ok && derivedDueDate != "" {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  derivedDueDate,
			Inline: true,
		})
	}

	priorityEmojiMappings := viper.GetStringMapString("openproject.priorityemojimappings")
	priorityName := payload.WorkPackage.Embedded.Priority.Name

	if priorityName != "" {
		if emoji, ok := priorityEmojiMappings[priorityName]; ok {
			embedBuilder.Field(discord.Field{
				Name:   "Priority",
				Value:  fmt.Sprintf("%s %s", emoji, priorityName),
				Inline: true,
			})
		} else {
			embedBuilder.Field(discord.Field{
				Name:   "Priority",
				Value:  priorityName,
				Inline: true,
			})
		}
	}

	statusEmojiMappings := viper.GetStringMapString("openproject.statusemojimappings")
	statusName := payload.WorkPackage.Embedded.Status.Name

	if statusName != "" {
		if emoji, ok := statusEmojiMappings[statusName]; ok {
			embedBuilder.Field(discord.Field{
				Name:   "Status",
				Value:  fmt.Sprintf("%s %s", emoji, statusName),
				Inline: true,
			})
		} else {
			embedBuilder.Field(discord.Field{
				Name:   "Status",
				Value:  statusName,
				Inline: true,
			})
		}
	}

	if payload.WorkPackage.Embedded.Responsible.Name != "" {
		embedBuilder.Field(discord.Field{
			Name:  "Responsible",
			Value: fmt.Sprintf("%s <%s>", payload.WorkPackage.Embedded.Responsible.Name, payload.WorkPackage.Embedded.Responsible.Email),
		})
	} else {
		embedBuilder.Field(discord.Field{
			Name:  "Responsible",
			Value: "None",
		})
	}

	if payload.WorkPackage.Embedded.Assignee.Name != "" {
		embedBuilder.Field(discord.Field{
			Name:  "Assignee",
			Value: fmt.Sprintf("%s <%s>", payload.WorkPackage.Embedded.Assignee.Name, payload.WorkPackage.Embedded.Assignee.Email),
		})
	} else {
		embedBuilder.Field(discord.Field{
			Name:  "Assignee",
			Value: "None",
		})
	}

	remainingTime := payload.WorkPackage.RemainingTime
	if remainingTime != "" {
		embedBuilder.Field(discord.Field{
			Name:   "Remaining time",
			Value:  remainingTime,
			Inline: false,
		})
	}

	webhookBuilder.
		Embed(
			embedBuilder.Embed(),
		)

	return webhookBuilder.Webhook(), nil
}

func NewOpenProjectService(logger *zap.Logger) *OpenProjectService {
	logger.Info("Executing NewOpenProjectService.")
	ops := OpenProjectService{logger}
	return &ops
}
