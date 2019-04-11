package swampling

import (
	irc "github.com/thoj/go-ircevent"
)

var Nick string = "swamplingbot"

type bot struct {
	context context
}

func (bot bot) HandleMessage(msg Message) {
	msg.Reaction().exec(bot.context)
}

type context interface {
	sendResponse(msg string)
}

type liveContext struct {
	chat chatConnection
}

func (ctx liveContext) sendResponse(msg string) {
	ctx.chat.sendResponse(msg)
}

type chatConnection interface {
	sendResponse(string)
}

// IRC

type ircChatConnection struct {
	channel string
	ircconn *irc.Connection
}

func (chat ircChatConnection) sendResponse(text string) {
	chat.ircconn.Privmsg(chat.channel, text)
}

func IrcBot(conn *irc.Connection, channel string) bot {
	return bot{context: liveContext{
		chat: ircChatConnection{
			channel: channel,
			ircconn: conn,
		},
	}}
}
