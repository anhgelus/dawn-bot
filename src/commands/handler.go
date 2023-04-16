package commands

import (
	"dawn-bot/src/config"
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/utils"
	"github.com/bwmarrin/discordgo"
	"time"
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

func getConfigHandler(cm commandHandler) {
	guildID := cm.i.GuildID
	conf, nw := config.GetConfig(guildID)
	if nw {
		cm.respond("Vous n'avez pas encore défini de config sur ce serveur :3")
		return
	}
	guild, err := cm.s.Guild(guildID)
	utils.PanicError(err)
	embed := discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Configuration du serveur " + guild.Name,
		Description: "Détails :",
		Color:       15418782,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Dawn Bot © 2023 - AGPLv3",
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: cm.i.Member.Nick,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
	welcomeField := discordgo.MessageEmbedField{
		Name:   "Salon de bienvenue",
		Inline: true,
	}

	var fields []*discordgo.MessageEmbedField
	if conf.WelcomeChannelID != "" {
		welcome, err := cm.s.Channel(cm.i.ChannelID)
		utils.PanicError(err)
		welcomeField.Value = welcome.Mention()
		fields = append(fields, &welcomeField)
	}
	if len(fields) == 0 {
		embed.Description = "Vous n'avez aucune configuration particulière"
	} else {
		embed.Fields = fields
	}
	cm.respondWithEmbed(&embed)
}

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

func (cm commandHandler) respondWithEmbed(e *discordgo.MessageEmbed) {
	err := cm.s.InteractionRespond(cm.i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{e},
		},
	})
	utils.PanicError(err)
}
