package command

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Handler     func(*discordgo.Session, *discordgo.MessageCreate)
}

// TODO: コマンド増えてきたらファイルに分割する

var Ping = Command{
	Name:    "ping",
	Handler: pingHandler,
}

func pingHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	session.ChannelMessageSend(event.ChannelID, "Pong!")
}
