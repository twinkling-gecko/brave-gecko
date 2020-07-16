package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/twinkling-gecko/brave-gecko/pkg/command"
)

var (
	commands []command.Command
)

func Init() {
	register(command.Ping)
}

// 引数を元に指定コマンドを実行する
func Run(order string, session *discordgo.Session, event *discordgo.MessageCreate) {
	if order == "help" {
		session.ChannelMessageSend(event.ChannelID, genCommandHelp())
		return
	}

	for _, command := range commands {
		if command.Name == order {
			command.Handler(session, event)
			break
		}
	}
}

// コマンドを登録する
func register(command command.Command) {
	commands = append(commands, command)
}

// それっぽいヘルプを生成する
func genCommandHelp() string {
	var help = []string{"Yes, brave-gecko.", "```"}
	for _, command := range commands {
		help = append(help, command.Name+" "+command.Description)
	}
	help = append(help, "```")

	return strings.Join(help, "\n")
}
