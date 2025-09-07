package errors

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed messages.json
var messagesFS embed.FS

type messageDef struct {
	StatusCode  int               `json:"statusCode"`
	MessageCode string            `json:"messageCode"`
	Message     map[string]string `json:"message"`
}

var messages map[string]messageDef

func init() {
	data, err := messagesFS.ReadFile("messages.json")
	if err != nil {
		panic(fmt.Sprintf("failed to load messages.json: %v", err))
	}
	if err := json.Unmarshal(data, &messages); err != nil {
		panic(fmt.Sprintf("invalid messages.json: %v", err))
	}
}
