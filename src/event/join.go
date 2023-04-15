package event

import (
	"dawn-bot/src/config"
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/utils"
	"github.com/bwmarrin/discordgo"
)

func HandleJoin(s *discordgo.Session, j *discordgo.GuildMemberAdd) {
	member := j.User
	user := postgres.User{
		DiscordID: member.ID,
		Name:      member.Username,
		XP:        0,
	}
	postgres.Db.Create(&user)
	_, err := s.ChannelMessageSend(config.Config.WelcomeChannelID, "Bienvenue "+member.Mention()+" sur Dawn City !")
	utils.PanicError(err)
}
