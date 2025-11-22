package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"lkurnia/notifier/client"
)

func main() {
	_ = godotenv.Load()

	sampleMessage := "this is a sample message"
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	notifier := client.Notifier{}
	notifier.SetupTelegram(apiKey, chatId)
	err := notifier.Telegram.SendMessage(sampleMessage)
	if err != nil {
		fmt.Println("Notifier Request Error : ", err.Error())
	}
}
