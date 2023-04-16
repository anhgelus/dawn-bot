package utils

import "github.com/bwmarrin/discordgo"

func IsChannelId(s *discordgo.Session, id string) bool {
	_, err := s.Channel(id)
	return err == nil
}
