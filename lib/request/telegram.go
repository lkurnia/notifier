package request

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const TELEGRAM_BASE_URL = "https://api.telegram.org"
const TELEGRAM_SEND_MESSAGE = "/sendMessage"

type TelegramRequest struct {
	BaseRequest
	apiKey string
	ChatId string `json:"chat_id"`
}

type TelegramSendMessageBody struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewTelegramRequest(apiKey string, chatId string) TelegramRequest {
	return TelegramRequest{
		BaseRequest: BaseRequest{BaseUrl: TELEGRAM_BASE_URL},
		apiKey:      apiKey,
		ChatId:      chatId,
	}
}

func (r *TelegramRequest) getBaseUrl() string {
	return r.BaseUrl + "/bot" + r.apiKey
}

func (r *TelegramRequest) getSendMessageBody(message string) ([]byte, error) {
	jsonData, err := json.Marshal(TelegramSendMessageBody{
		ChatId: r.ChatId,
		Text:   message,
	})
	if err != nil {
		log.Println("Error generating SendMessageBody. Error : ", err.Error())
		return []byte{}, err
	}

	return jsonData, nil
}

func (r *TelegramRequest) SendMessage(message string) error {
	payload, err := r.getSendMessageBody(message)
	if err != nil {
		log.Println("Error getting payload for SendMessage request. Error : ", err.Error())
		return err
	}

	// Create Http Request
	req, err := http.NewRequest("POST", r.getBaseUrl()+TELEGRAM_SEND_MESSAGE, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
