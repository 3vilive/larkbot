package larkbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/3vilive/larkbot/message"
)

type WebhookResp struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

type Bot struct {
	webhook     string
	secretToken string
	httpClient  *http.Client
}

func NewBot(webhook string) *Bot {
	return &Bot{
		webhook:    webhook,
		httpClient: &http.Client{},
	}
}

func NewBotWithSecretToken(webhook string, token string) *Bot {
	bot := NewBot(webhook)
	bot.secretToken = token

	return bot
}

func (b *Bot) requestWebhook(payload map[string]interface{}) error {
	if b.secretToken != "" {
		unixNow := time.Now().Unix()
		signature, err := GenerateSignature(b.secretToken, unixNow)
		if err != nil {
			return err
		}

		payload["sign"] = signature
		payload["timestamp"] = strconv.FormatInt(unixNow, 10)
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := b.httpClient.Post(b.webhook, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var webhookResp WebhookResp
	if err := json.Unmarshal(respBody, &webhookResp); err != nil {
		return err
	}

	if webhookResp.StatusCode != 0 {
		return fmt.Errorf("unexpected response: %s", webhookResp.StatusMessage)
	}

	return nil
}

func (b *Bot) SendTextMessage(text string) error {
	payload := message.BuildTextMessagePayload(text)
	return b.requestWebhook(payload)
}
