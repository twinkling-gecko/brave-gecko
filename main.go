package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/twinkling-gecko/brave-gecko/pkg/eventhandler"
	"github.com/twinkling-gecko/brave-gecko/pkg/router"
)

var (
	Token = os.Getenv("DISCORD_TOKEN")
)

func main() {
	app, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	router.Init()
	app.AddHandler(eventhandler.MessageCreateHandler)

	err = app.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	app.Close()
}
