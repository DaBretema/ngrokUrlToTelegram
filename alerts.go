package main

import "log"

const _NoToken = `Token of telegram bot is not defined.

---* WINDOWS - On elevated powershell *---
[System.Environment]::SetEnvironmentVariable('ngrokUrlBot', '<BOT_TOKEN>', [System.EnvironmentVariableTarget]::User)

---* LINUX / MACOS - As your user *---
"export ngrokUrlBot = <BOT_TOKEN>\" >> ~/.profile
`

const _NgrokDown = "Ngrok public url not found. Maybe Ngrok is DOWN?"

const _TgSleep = "Please wakeup the bot before run me :)"

// -----

func recov(alertMsg string) {
	if r := recover(); r != nil {
		log.Fatalln(alertMsg)
	}
}
