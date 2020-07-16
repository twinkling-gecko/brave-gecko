package command

import (
	"net/http"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string // コマンド名
	Description string // 書いておくとhelpに出る
	Handler     func(*discordgo.Session, *discordgo.MessageCreate)
}

type MegurusayResponse struct {
	Message string `json:message`
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

var Megurusay = Command{
	Name:        "megurusay",
	Description: "めぐる、Discordで喋れるようになっちゃいました！",
	Handler:     megurusayHandler,
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

func megurusayHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	apiUrl := "http://say.tatikaze.com/meguru"

	res, err := http.Get(apiUrl)
	if err != nil {
		session.ChannelMessageSend(event.ChannelID, err.Error())
		return
	}
	defer res.Body.Close()

	var megurusayResponse MegurusayResponse
	if err:= json.NewDecoder(res.Body).Decode(&megurusayResponse); err != nil {
		session.ChannelMessageSend(event.ChannelID, err.Error())
		return
	}

	session.ChannelMessageSend(event.ChannelID, megurusayResponse.Message)
}
