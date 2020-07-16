package command

import (
	"net/http"
	"strconv"
	"strings"

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

var Status = Command{
	Name:        "status",
	Description: "第一引数のURIのステータスコードを返す",
	Handler:     statusHandler,
}

func pingHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	session.ChannelMessageSend(event.ChannelID, "Pong!")
}

func statusHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
  arg := strings.Split(event.Content, " ")[1]

	res, err := http.Get(arg)
	if err != nil {
		session.ChannelMessageSend(event.ChannelID, err.Error())
		return
	}

  result := strconv.Itoa(res.StatusCode)
	session.ChannelMessageSend(event.ChannelID, result)
}
