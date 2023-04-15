package event

import (
	"dawn-bot/src/commands/message"
	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c := m.Message.Content
	if c[:1] == message.Prefix {
		message.GlobalHandler(s, m)
	}
}
