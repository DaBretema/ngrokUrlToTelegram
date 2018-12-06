package main

import (
	"os"
)

const (
	ngTunnels = "http://127.0.0.1:4040/api/tunnels/"
	tgGetID   = "https://api.telegram.org/bot%v/getUpdates"
	tgSend    = "https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v"
)

var trys = 0

func main() {

	// 1.- Get ngrok url
	ng := ngFromURI(ngTunnels)
	ngrokURL := ng.Tunnels[0].PublicURL

	// 2.- Get token from os env
	token := os.Getenv("ngrokUrlBot")
	if token == "" {
		errxit(_TgNoToken)
	}

	// 3.- Send :D
	tg := newTg(token, ngrokURL)
	tg.send()
}
