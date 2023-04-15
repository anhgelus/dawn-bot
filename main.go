package main

import (
	"dawn-bot/src/config"
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

func main() {
	var err error
	utils.GlobalPath, err = os.Executable()
	utils.PanicError(err)

	generateConfigs()
	loadConfigs()

	token := os.Args[1]
	client, err := discordgo.New("Bot " + token)
	utils.PanicError(err)

	client.Identify.Intents = discordgo.IntentsGuildMessages

	err = client.Open()
	utils.PanicError(err)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	client.Close()
}

func generateConfigs() {
	err := os.Mkdir("config", 0666)
	utils.PanicError(err)
	if !utils.FileExist(utils.FilePath("/config/databases.toml")) {
		file, err := os.Create("config/databases.toml")
		utils.PanicError(err)
		defer file.Close()

		content, err := configs.ReadFile("resources/config/databases.toml")
		utils.PanicError(err)

		_, err = file.Write(content)
		utils.PanicError(err)
	}
}

func loadConfigs() {
	databases := config.LoadAndImportDatabaseConfig()
	databases.Postgres.Connect()
	databases.Redis.Connect()
}
