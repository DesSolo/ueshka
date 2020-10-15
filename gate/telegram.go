package gate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ueshka/ueshka"
)

// Telegram ...
type Telegram struct {
	Token  string
	ChatID string

	client *http.Client
}

// NewTelegram ...
func NewTelegram(token, chatID string) *Telegram {
	return &Telegram{
		Token:  token,
		ChatID: chatID,

		client: &http.Client{},
	}
}

// telegramMessagePayload https://core.telegram.org/bots/api#sendmessage
type telegramMessagePayload struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// Send ...
func (t *Telegram) Send(msg string) error {
	payload := telegramMessagePayload{
		ChatID:    t.ChatID,
		Text:      msg,
		ParseMode: "Markdown",
	}
	data, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.telegram.org/bot"+t.Token+"/sendMessage", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return nil
}

// RenderMessage ...
func (t *Telegram) RenderMessage(o *ueshka.Operation) string {
	return fmt.Sprintf("⚡️ `%s`: %s", o.Time, o.Name)
}
