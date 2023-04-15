package main

import (
	"dawn-bot/src/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	token := os.Args[1]
	client, err := discordgo.New("Bot " + token)
	utils.PanicError(err)

	client.Identify.Intents = discordgo.IntentsGuildMessages

	err = client.Open()
	utils.PanicError(err)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	client.Close()
}
