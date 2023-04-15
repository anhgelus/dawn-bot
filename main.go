package main

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

func main() {
	token := os.Args[1]
	client, err := discordgo.New("Bot " + token)
}
