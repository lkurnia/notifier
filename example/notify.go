package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"lkurnia/notifier/lib/request"
)

func main() {
	_ = godotenv.Load()

	sampleMessage := "this is a sample message"
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	notifier := request.NewTelegramRequest(apiKey, chatId)
	err := notifier.SendMessage(sampleMessage)
	if err != nil {
		fmt.Println("Notifier Request Error : ", err.Error())
	}
}
