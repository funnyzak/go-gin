package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WeComPayload struct {
	Key     string
	Message string
}

type WeCom struct {
	Payload WeComPayload
}

func (w WeCom) Send(title string, message string) error {
	key := w.Payload.Key

	sendMessageUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)
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
