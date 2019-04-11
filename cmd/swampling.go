package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lpil/swampling/pkg/swampling"
	irc "github.com/thoj/go-ircevent"
)

const channel = "#lpil"
const serverssl = "irc.freenode.net:7000"

func main() {
	irccon := irc.IRC(swampling.Nick, "IRCTestSSL")
	irccon.VerboseCallbackHandler = false

	irccon.Debug = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Join channel
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})

	bot := swampling.IrcBot(irccon, channel)

	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
		msg := swampling.Message{Text: event.Message(), From: event.Nick}
		bot.HandleMessage(msg)
	})

	// Connect
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	go irccon.Loop()

	defer fmt.Println("\nGoodbye")

	waitForExitSignal()
}

// Other

func waitForExitSignal() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
