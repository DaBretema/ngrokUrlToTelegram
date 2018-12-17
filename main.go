package main

import "os"

const (
	ngTunnels = "http://127.0.0.1:4040/api/tunnels/"
	tgGetID   = "https://api.telegram.org/bot%v/getUpdates"
	tgSend    = "https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v"
)

func main() {

	// Get telegram token from os env
	token := os.Getenv("ngrokUrlBot")
	if token == "" {
		errxit(_TgNoToken)
	}

	// Get ngrok tunnel url
	ng := newNg()

	// Send using telegram
	tg := newTg(token, ng.url)
	tg.send()
}
