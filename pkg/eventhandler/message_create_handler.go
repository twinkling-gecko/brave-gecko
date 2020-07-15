package eventhandler

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/twinkling-gecko/brave-gecko/pkg/router"
)

var (
	prefix = "/"
)

func MessageCreateHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	// 発言者がこのBotだった場合は相手をしない
	if session.State.User.ID == event.Author.ID {
		return
	}

	// Contentがprefixから始まる場合、コマンドルーターに実行依頼
	if strings.Index(event.Content, prefix) == 0 {
		order := strings.Replace(
			strings.Split(event.Content, " ")[0],
			prefix, "", 1,
		)

		router.Run(order, session, event)
	}
}
