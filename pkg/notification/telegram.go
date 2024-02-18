package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramPayload struct {
	BotToken string
	ChatId   string
}

type Telegram struct {
	Payload TelegramPayload
}

func (t Telegram) Send(title string, message string) error {
	botToken := t.Payload.BotToken
	chatId := t.Payload.ChatId

	text := fmt.Sprintf("%s\n%s", title, message)

	sendMessageUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	sendMessageBody := map[string]string{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "HTML",
	}
	jsonBody, _ := json.Marshal(sendMessageBody)
	resp, err := http.Post(sendMessageUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
