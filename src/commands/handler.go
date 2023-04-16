package commands

import (
	"dawn-bot/src/config"
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/utils"
	"github.com/bwmarrin/discordgo"
)

type commandHandler struct {
	s *discordgo.Session
	i *discordgo.InteractionCreate
}

func GlobalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	t, valid := Handlers[i.ApplicationCommandData().Name]
	if !valid {
		panic("impossible to find the slash command with the name " + i.ApplicationCommandData().Name)
	}
	t(commandHandler{s: s, i: i})
}

func configHandler(cm commandHandler) {
	options := cm.i.ApplicationCommandData().Options
	if len(options) != 2 {
		cm.respond("Il semblerait que vous avez oublié de remplir toutes les options...")
		return
	}
	values := optionsArrayToMap(options)

	var option *discordgo.ApplicationCommandInteractionDataOption
	var ok bool
	if option, ok = values["parameter"]; !ok {
		cm.respond("Impossible de trouver le paramètre à modifier :(")
		return
	}
	typ := option.IntValue()

	if option, ok = values["value"]; !ok {
		cm.respond("Impossible de trouver la valeur à entrer :(")
		return
	}
	value := option.StringValue()
	var param string

	conf, _ := config.GetConfig(cm.i.GuildID)

	switch typ {
	case 1:
		if !utils.IsChannelId(cm.s, value) {
			cm.respond("Impossible de trouver le salon avec l'ID `" + value + "`")
			return
		}
		param = "salon de bienvenue"
		conf.WelcomeChannelID = value
	default:
		cm.respond("Le paramètre entré n'existe pas :'(")
		return
	}
	postgres.Db.Save(&conf)
	cm.respond("Le paramètre `" + param + "` a bien été changé en `" + value + "`")
}

//func getConfigHandler(cm commandHandler) {
//	conf, nw := config.GetConfig(cm.i.GuildID)
//	if nw {
//		cm.respond("Vous n'avez pas encore défini de config sur ce serveur :3")
//	}
//}

func optionsArrayToMap(o []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(o))
	for _, opt := range o {
		optionMap[opt.Name] = opt
	}
	return optionMap
}

func (cm commandHandler) respond(c string) {
	err := cm.s.InteractionRespond(cm.i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
		},
	})
	utils.PanicError(err)
}
