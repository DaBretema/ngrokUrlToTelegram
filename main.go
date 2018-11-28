package main

import (
	"fmt"
	"log"
)

var _uri = map[string]string{
	"ngrok":    "http://127.0.0.1:4040/api/tunnels/",
	"tgChatID": "https://api.telegram.org/bot%v/getUpdates",
	"tgMsg":    "https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v",
}

const token = "<BOT_TOKEN>"

func main() {
	tgSend(ngrokURL())
}

func ngrokURL() string {
	// Get reply of ngrok api
	var ngr ngrokReply
	ngr.FillFromURI(_uri["ngrok"])
	// Caputre public url of ngrok tunnel
	return ngr.Tunnels[0].PublicURL
}

func tgSend(msg string) {
	if r := recover(); r != nil {
		log.Printf("Please \"/start\" the bot if you haven't: %v\n", r)
	}
	// Get reply of telegram api by token
	var tgr tgUpdatesReply
	tgr.FillFromURI(fmt.Sprintf(_uri["tgChatID"], token))
	// Capture chat id
	chatID := tgr.Result[0].Message.Chat.ID
	// Send messagge
	doGetReq(fmt.Sprintf(_uri["tgMsg"], token, chatID, msg))
}
