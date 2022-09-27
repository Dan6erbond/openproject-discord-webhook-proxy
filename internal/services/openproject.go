package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/discord"
	"github.com/dan6erbond/openproject-discord-webhook-proxy/internal/openproject"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type OpenProjectService struct {
	logger *log.Logger
}

func (ops *OpenProjectService) ValidateSignature(body []byte, webhook map[string]interface{}, r *http.Request) error {
	signature := r.Header["X-Op-Signature"]
	h := hmac.New(sha1.New, []byte(webhook["secret"].(string)))

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

func (ops *OpenProjectService) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webhooks := viper.GetStringMap("webhooks")
	var webhook map[string]interface{}
	for name, w := range webhooks {
		if wh, ok := w.(map[string]interface{}); ok {
			path, ok := wh["path"].(string)
			if (ok && path != "" && path == vars["name"]) || name == vars["name"] {
				webhook = wh
				break
			}
		}
	}

	if webhook == nil {
		http.Error(w, "Couldn't find matching webhook", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/* err = ops.ValidateSignature(body, webhook, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} */

	var payload openproject.Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actions := webhook["actions"].([]interface{})
	for _, action := range actions {
		if payload.Action == action.(string) {
			switch payload.Action {
			case "work_package:created":
				ops.HandleWorkPackageCreatedWebhook(body, webhook, w, r)
				return
			default:
				http.Error(w, "Couldn't find action for webhook", http.StatusNotFound)
			}
		}
	}
}

func (ops *OpenProjectService) HandleWorkPackageCreatedWebhook(body []byte, webhook map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	var payload openproject.WorkPackageWebhookPayload
	err := json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	color, err := strconv.ParseInt(payload.WorkPackage.Embedded.Type.Color[1:], 16, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	openProjectBaseUrl, err := url.Parse(viper.GetString("openproject.baseurl"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projectUrl := url.URL{
		Scheme: openProjectBaseUrl.Scheme,
		Host:   openProjectBaseUrl.Host,
		Path:   fmt.Sprintf("/projects/%s", payload.WorkPackage.Embedded.Project.Identifier),
	}

	ops.logger.Println(viper.GetString("openproject.baseurl"))
	workPackageUrl := url.URL{
		Scheme: openProjectBaseUrl.Scheme,
		Host:   openProjectBaseUrl.Host,
		Path:   fmt.Sprintf("/projects/%s/work_packages/%d/activity", payload.WorkPackage.Embedded.Project.Identifier, payload.WorkPackage.ID),
	}

	webhookBuilder := discord.
		WebhookBuilder().
		Content("Work package created")
	embedBuilder := discord.EmbedBuilder().
		Author(discord.Author{
			Name:    payload.WorkPackage.Embedded.Project.Name,
			IconURL: payload.WorkPackage.Embedded.Author.Avatar,
			URL:     projectUrl.String(),
		}).
		Color(color).
		Title(fmt.Sprintf("%s: %s", payload.WorkPackage.Embedded.Type.Name, payload.WorkPackage.Subject)).
		URL(workPackageUrl.String()).
		Description(payload.WorkPackage.Description.Raw)
	if startDate, ok := payload.WorkPackage.StartDate.(string); ok {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  startDate,
			Inline: true,
		})
	}
	if derivedStartDate, ok := payload.WorkPackage.DerivedStartDate.(string); ok {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  derivedStartDate,
			Inline: true,
		})
	}
	if dueDate, ok := payload.WorkPackage.DueDate.(string); ok {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  dueDate,
			Inline: true,
		})
	}
	if derivedDueDate, ok := payload.WorkPackage.DerivedDueDate.(string); ok {
		embedBuilder.Field(discord.Field{
			Name:   "Start date",
			Value:  derivedDueDate,
			Inline: true,
		})
	}

	priorityEmojiMappings := viper.GetStringMapString("openproject.priorityemojimappings")

	if emoji, ok := priorityEmojiMappings[payload.WorkPackage.Embedded.Status.Name]; ok {
		embedBuilder.Field(discord.Field{
			Name:   "Priority",
			Value:  fmt.Sprintf("%s %s", emoji, payload.WorkPackage.Embedded.Priority.Name),
			Inline: true,
		})
	} else {
		embedBuilder.Field(discord.Field{
			Name:   "Priority",
			Value:  payload.WorkPackage.Embedded.Priority.Name,
			Inline: true,
		})
	}

	statusEmojiMappings := viper.GetStringMapString("openproject.statusemojimappings")

	if emoji, ok := statusEmojiMappings[payload.WorkPackage.Embedded.Status.Name]; ok {
		embedBuilder.Field(discord.Field{
			Name:   "Status",
			Value:  fmt.Sprintf("%s %s", emoji, payload.WorkPackage.Embedded.Status.Name),
			Inline: true,
		})
	} else {
		embedBuilder.Field(discord.Field{
			Name:   "Status",
			Value:  payload.WorkPackage.Embedded.Status.Name,
			Inline: true,
		})
	}

	embedBuilder.Field(discord.Field{
		Name:  "Responsible",
		Value: fmt.Sprintf("%s <%s>", payload.WorkPackage.Embedded.Responsible.Name, payload.WorkPackage.Embedded.Responsible.Email),
	})

	embedBuilder.Field(discord.Field{
		Name:  "Assignee",
		Value: fmt.Sprintf("%s <%s>", payload.WorkPackage.Embedded.Assignee.Name, payload.WorkPackage.Embedded.Assignee.Email),
	})

	embedBuilder.Field(discord.Field{
		Name:   "Remaining time",
		Value:  payload.WorkPackage.RemainingTime,
		Inline: false,
	})

	webhookBuilder.
		Embed(
			embedBuilder.Embed(),
		)

	content, err := json.Marshal(webhookBuilder.Webhook())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contentReader := bytes.NewReader(content)
	resp, err := http.Post(webhook["url"].(string), "application/json", contentReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if string(respBody) != "" {
		http.Error(w, string(respBody), http.StatusInternalServerError)
	}

	if resp.StatusCode != http.StatusCreated {
		http.Error(w, "Discord error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func NewOpenProjectService(logger *log.Logger) *OpenProjectService {
	logger.Print("Executing NewOpenProjectService.")
	ops := OpenProjectService{logger}
	return &ops
}
