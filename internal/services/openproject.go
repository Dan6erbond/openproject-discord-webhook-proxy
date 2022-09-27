package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type OpenProjectService struct {
	logger *log.Logger
}

func (ops *OpenProjectService) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webhooks := viper.GetStringMap("webhooks")
	for name, w := range webhooks {
		if webhook, ok := w.(map[string]interface{}); ok {
			path, ok := webhook["path"].(string)
			if (ok && path != "" && path == vars["name"]) || name == vars["name"] {
				ops.logger.Printf("Received webhook %s payload", vars["name"])
			}
		}
	}
}

func NewOpenProjectService(logger *log.Logger) *OpenProjectService {
	logger.Print("Executing NewOpenProjectService.")
	ops := OpenProjectService{logger}
	return &ops
}
