package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"strings"
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

	irccon.AddCallback("PRIVMSG",
		handlePRIVMSG(irccon))

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

func handlePRIVMSG(irccon *irc.Connection) func(*irc.Event) {
	return func(event *irc.Event) {
		if strings.HasPrefix(event.Message(), swampling.Nick) {

			irccon.Privmsg("#lpil", "Hello!")
		}
		return
	}
}
