package commands

import (
	"dawn-bot/src/utils"
	"github.com/bwmarrin/discordgo"
)

var (
	configCommand = &discordgo.ApplicationCommand{
		Name:        "config",
		Description: "Config le bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "parameter",
				Description: "Paramètre à modifier",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "welcome-channel",
						Value: 1,
					},
				},
				Required: true,
			},
			{
				Name:        "value",
				Description: "Entrez la valeur attendue ici (s'il s'agit d'un rôle/salon, entrez un ID)",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}
	RegisteredCommands = make([]*discordgo.ApplicationCommand, 1)
	Commands           = [1]*discordgo.ApplicationCommand{configCommand}
	Handlers           = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		configCommand.Name: ConfigHandler,
	}
)

func Init(s *discordgo.Session) {
	for i, c := range Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", c)
		utils.PanicError(err)
		RegisteredCommands[i] = cmd
	}
}

func Remove(s *discordgo.Session) {
	for i, c := range RegisteredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", c.ID)
		utils.PanicError(err)
		RegisteredCommands[i] = nil
	}
}
