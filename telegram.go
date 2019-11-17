package main

import (
	"fmt"

	"github.com/vadymc/telegram-client-go/v2"
)

var telegramClient *telegram.TelegramClient

func init() {
	telegramClient = telegram.NewTelegramClient()
}

func sendTelegram(event []byte) {
	title, body := ParseEvent(event)
	text := fmt.Sprintf("[%v] %v", title, body)
	telegramClient.SendMessage("HeartBeat", text)
	fmt.Printf("Sent to telegram %v\n", text)
}
