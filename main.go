package main

import (
	"fmt"
	"log"
	"os"

	"github.com/glucktek/gfc-d-bot/pkgs/bot"
)

func main() {

	fmt.Println("Starting bot service...")
	token := os.Getenv("DISCORD_BOT_TOKEN")
	adminRole := os.Getenv("DISCORD_ADMIN_ROLE") // Role ID that is allowed to use the bot
	guildID := os.Getenv("DISCORD_GUILD_ID")

	//Pre flight check
	if token == "" || adminRole == "" || guildID == "" {
		log.Fatal("Missing required environment variables")
	}

	// Initialize Discord Bot
	b, err := bot.New(token, guildID, adminRole)
	if err != nil {
		log.Fatal(err)
	}
	//Run the bot
	if err := b.Start(); err != nil {
		log.Fatal(err)
	}

}
