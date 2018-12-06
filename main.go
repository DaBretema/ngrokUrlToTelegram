package main

import (
	"fmt"
	"os"
)

const ngTunnels = "http://127.0.0.1:4040/api/tunnels/"
const tgGetID = "https://api.telegram.org/bot%v/getUpdates"
const tgSend = "https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v"

func main() {
	defer recov(_TgSleep)

	// 1.- Get ngrok url
	var ng ngModel
	ng.FillFromURI(ngTunnels)
	ngrokURL := ng.Tunnels[0].PublicURL

	// 2.- Get token from os env
	token := os.Getenv("ngrokUrlBot")
	if token == "" {
		errxit(_TgNoToken)
	}

	// 3.- Get chat id from telegram api
	var tgr tgModel
	tgr.FillFromURI(fmt.Sprintf(tgGetID, token))
	chatID := tgr.Result[0].Message.Chat.ID

	// 4.- Send ngrok url to telegram bot
	hGet(fmt.Sprintf(tgSend, token, chatID, ngrokURL))
}
