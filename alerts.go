package main

import (
	"fmt"
	"os"
)

const _NgrokDown = "Ngrok public url not found. Maybe Ngrok is DOWN?"
const _TgBadToken = "Telegram request error. Maybe BAD token?"
const _TgSleep = "Please wakeup the bot before run me :)"
const _TgNoToken = `Token of telegram bot is not defined.

---* WINDOWS - On elevated powershell *---
[System.Environment]::SetEnvironmentVariable('ngrokUrlBot', '<BOT_TOKEN>', [System.EnvironmentVariableTarget]::User)

---* LINUX / MACOS - As your user *---
"export ngrokUrlBot = <BOT_TOKEN>\" >> ~/.profile
`

// -----

func errxit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(2)
}
func recov(alertMsg string) {
	if r := recover(); r != nil {
		errxit(alertMsg)
	}
}
