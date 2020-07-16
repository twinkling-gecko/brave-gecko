package command

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string // コマンド名
	Description string // 書いておくとhelpに出る
	Handler     func(*discordgo.Session, *discordgo.MessageCreate)
}

// TODO: コマンド増えてきたらファイルに分割する

var Ping = Command{
	Name:        "ping",
	Description: "灼熱の卓球娘",
	Handler:     pingHandler,
}

func pingHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	session.ChannelMessageSend(event.ChannelID, "Pong!")
}
