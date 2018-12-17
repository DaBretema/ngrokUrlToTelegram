package main

import "fmt"

type tgram struct {
	token     string
	chatID    int
	msg       string
	idURI     string
	sendTries int
}

func newTg(token, msg string) *tgram {
	return &tgram{
		token:     token,
		msg:       msg,
		idURI:     fmt.Sprintf(tgGetID, token),
		sendTries: 0,
	}
}

func (tg *tgram) send() {

	// Exit after some tries
	if tg.sendTries >= _Tries {
		errxit(_TgSleep)
	}
	// If fails try again (sometimes bot is sleeped)
	tg.sendTries++
	defer recovWithFunc(tg.send)
	// 1.- Get chat id from telegram api
	tg.chatID = tgFromURI(tg.idURI).Result[0].Message.Chat.ID
	// 2.- Send ngrok url to telegram bot
	hGet(fmt.Sprintf(tgSend, tg.token, tg.chatID, tg.msg))
}
