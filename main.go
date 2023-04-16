package main

import (
	"dawn-bot/src/commands"
	"dawn-bot/src/config"
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/event"
	"dawn-bot/src/utils"
	"embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

//go:embed resources/config
var configs embed.FS

const (
	intents = discordgo.IntentsAll
)

func main() {
	var err error
	utils.GlobalPath, err = os.Executable()
	utils.PanicError(err)

	// generate static configs file
	generateConfigs()
	loadConfigs()
	// Migrate the database
	postgres.Migrate()

	// Generate the websocket
	if len(os.Args) != 2 {
		println("You must provide only the token")
		return
	}
	token := os.Args[1]
	client, err := discordgo.New("Bot " + token)
	utils.PanicError(err)

	// Handle classical event
	client.AddHandler(event.HandleJoin)

	// Add used intents
	client.Identify.Intents = discordgo.MakeIntent(intents)

	// Connect the bot
	err = client.Open()
	commands.Init(client)
	client.AddHandler(commands.GlobalHandler)

	utils.PanicError(err)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.* and delete the slash commands
	commands.Remove(client)
	client.Close()
}

func generateConfigs() {
	err := os.Mkdir("config", 0777)
	if err != nil {
		println(err.Error())
	}
	if !utils.FileExist("config/databases.toml") {
		println("the file do not exist")
		file, err := os.Create("config/databases.toml")
		utils.PanicError(err)

		content, err := configs.ReadFile("resources/config/databases.toml")
		utils.PanicError(err)

		_, err = file.Write(content)
		utils.PanicError(err)

		err = file.Close()
		utils.PanicError(err)
	}
}

func loadConfigs() {
	databases := config.LoadAndImportDatabaseConfig()
	databases.Postgres.Connect()
	databases.Redis.Connect()
}
