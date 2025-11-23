package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"io"
)

const TELEGRAM_BASE_URL = "https://api.telegram.org"
const TELEGRAM_SEND_MESSAGE = "/sendMessage"
const TELEGRAM_SEND_PHOTO = "/sendPhoto"

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

func (r *TelegramRequest) getSendPhotoBody(fileDir string, filename string) (*bytes.Buffer, string, error) {
	filePath := path.Join(fileDir, filename)

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file. Error : " + err.Error())
		return &bytes.Buffer{}, "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error creating form file. Error : " + err.Error())
		return &bytes.Buffer{}, "", err
	}

	io.Copy(part, file)
	writer.Close()

	return body, writer.FormDataContentType(), nil
}

func (r *TelegramRequest) SendPhoto(fileDir string, filename string) error {
	body, contentType, err := r.getSendPhotoBody(fileDir, filename)
	if err != nil {
		log.Println("Error getting payload for SendPhoto request. Error : ", err.Error())
		return err
	}

	// Create Http Request
	req, err := http.NewRequest("POST", r.getBaseUrl() + TELEGRAM_SEND_PHOTO + "?chat_id=" + r.ChatId, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Debug line
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}

	log.Println("### Sending Photo ###")
	log.Println("Response Status:", resp.Status)
	log.Println("Response Body:", string(bdy))

	return nil
}
