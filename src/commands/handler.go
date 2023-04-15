package commands

import (
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/utils"
	"github.com/bwmarrin/discordgo"
)

func GlobalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	t, valid := Handlers[i.ApplicationCommandData().ID]
	if !valid {
		panic("impossible to find the slash command with the id " + i.ID)
	}
	t(s, i)
}

func ConfigHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) != 2 {
		respondOnError(s, i, "Il semblerait que vous avez oublié de remplir toutes les options...")
		return
	}
	values := optionsArrayToMap(options)

	var option *discordgo.ApplicationCommandInteractionDataOption
	var ok bool
	if option, ok = values["parameter"]; !ok {
		respondOnError(s, i, "Impossible de trouver le paramètre à modifier :(")
		return
	}
	typ := option.IntValue()

	if option, ok = values["value"]; !ok {
		respondOnError(s, i, "Impossible de trouver la valeur à entrer :(")
		return
	}
	value := option.StringValue()

	switch typ {
	case 1:
		postgres.ConfigDB.WelcomeChannelID = value
	default:
		respondOnError(s, i, "Le paramètre entré n'existe pas :'(")
		return
	}
	postgres.Db.Save(&postgres.ConfigDB)
}

func optionsArrayToMap(o []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(o))
	for _, opt := range o {
		optionMap[opt.Name] = opt
	}
	return optionMap
}

func respondOnError(s *discordgo.Session, i *discordgo.InteractionCreate, c string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
		},
	})
	utils.PanicError(err)
}
