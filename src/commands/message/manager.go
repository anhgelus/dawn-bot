package message

import (
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

const (
	Prefix = "!"
)

type Command struct {
	Command string
	Args    uint
	Handler func(*discordgo.Session, *discordgo.MessageCreate)
}

var (
	joinCommand = Command{
		Command: "join",
		Args:    1,
		Handler: join,
	}

	Commands    = [1]Command{joinCommand}
	CommandsMap = map[string]Command{}
	inited      = false
)

func init() {
	if inited {
		return
	}
	for _, v := range Commands {
		CommandsMap[v.Command] = v
	}
	inited = true
}

// GlobalHandler handles every message command
func GlobalHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	c := msg.Content
	split := strings.Split(c, " ")

	var command Command
	var ok bool
	if command, ok = CommandsMap[split[0]]; !ok {
		_, err := s.ChannelMessageSend(msg.ChannelID, "Cette commande n'existe pas :(\nUtilise `!help` pour obtenir l'aide !")
		if err != nil {
			println("MessageCommand not found, when send sending the information an error occurred:", err.Error())
		}
	}
	if len(split) != calculateArgs(command) {
		_, err := s.ChannelMessageSend(msg.ChannelID, command.Command+" a besoin de "+strconv.Itoa(int(command.Args))+" argument.s !")
		if err != nil {
			println("Command found without good information, when send sending this information an error occurred:", err.Error())
		}
	}
	command.Handler(s, m)
}

func join(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	c := msg.Content
	split := strings.Split(c, " ")
	d := split[1]
	user := postgres.User{DiscordID: m.Author.ID}
	district := postgres.District{}

	db := postgres.Db
	db.First(&user)
	db.First(&district, d)

	var users []postgres.User
	db.Where("district_id = $1", district.ID).Table("users").Find(&users)

	if len(users) >= int(district.Max) {
		_, err := s.ChannelMessageSend(m.ChannelID, "Ce quartier est déjà plein à craquer :'3")
		utils.PanicError(err)
	}

	user.District = district
	db.Save(&user)

	err := s.GuildMemberRoleAdd(msg.GuildID, user.DiscordID, district.RoleID)
	utils.PanicError(err)
}

func calculateArgs(c Command) int {
	return int(c.Args + 1)
}
