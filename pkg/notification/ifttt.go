package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IFTTTPayload struct {
	Key   string
	Event string
}

type IFTTT struct {
	Payload IFTTTPayload
}

func (i IFTTT) Send(title string, message string) error {
	key := i.Payload.Key
	event := i.Payload.Event
	value1 := title
	value2 := message

	sendMessageUrl := fmt.Sprintf("https://maker.ifttt.com/trigger/%s/with/key/%s", event, key)
	sendMessageBody := map[string]string{
		"value1": value1,
		"value2": value2,
	}
	jsonBody, _ := json.Marshal(sendMessageBody)
	resp, err := http.Post(sendMessageUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
