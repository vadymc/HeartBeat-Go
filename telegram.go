package main

import (
	"fmt"

	tg "github.com/vadymc/telegram-client-go/v2"
)

var telegramClient *tg.TelegramClient

func init() {
	telegramClient = tg.NewTelegramClient()
}

func sendTelegram(event []byte) {
	title, body := ParseEvent(event)
	text := fmt.Sprintf("[%v] %v", title, body)
	telegramClient.SendMessage(text)
	fmt.Printf("Sent to telegram %v\n", text)
}
