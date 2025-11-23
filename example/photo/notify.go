package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/lkurnia/notifier/client"
)

func main() {
	_ = godotenv.Load()

	sampleDir := "example/photo"
	samplePhoto := "randomqr.png"
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	notifier := client.Notifier{}
	notifier.SetupTelegram(apiKey, chatId)
	err := notifier.Telegram.SendPhoto(sampleDir, samplePhoto)
	if err != nil {
		fmt.Println("Notifier Request Error : ", err.Error())
	}
}
