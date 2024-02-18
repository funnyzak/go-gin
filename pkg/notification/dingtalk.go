package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DingTalkPayload struct {
	Webhook string
	Message string
}

type DingTalk struct {
	Payload DingTalkPayload
}

func (d DingTalk) Send(title string, message string) error {
	webhook := d.Payload.Webhook

	sendMessageUrl := webhook
	sendMessageBody := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("%s\n%s", title, message),
		},
	}
	jsonBody, _ := json.Marshal(sendMessageBody)
	resp, err := http.Post(sendMessageUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
