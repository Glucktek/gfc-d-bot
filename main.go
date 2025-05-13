package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/glucktek/gfc-d-bot/pkgs/bot"
	"github.com/glucktek/gfc-d-bot/pkgs/lightsail"
)

func main() {

	fmt.Println("Starting bot service...")
	token := os.Getenv("DISCORD_BOT_TOKEN")
	adminRole := os.Getenv("DISCORD_ADMIN_ROLE") // Role ID that is allowed to use the bot
	guildID := os.Getenv("DISCORD_GUILD_ID")

	if token == "" || len(adminRole) == 0 || guildID == "" {
		log.Fatal("Missing required environment variables")
	}

	b, err := bot.New(token, guildID, adminRole)
	if err != nil {
		log.Fatal(err)
	}

	if err := b.Start(); err != nil {
		log.Fatal(err)
	}

	// Initialize the Lightsail client
	client, err := lightsail.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// HACK: Using hardcaded instance name for now
	instanceName := "GreaterFaithChurchSite"

	// Get status
	state, err := client.GetInstanceState(ctx, instanceName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Instance state: %s\n", state)

	// Start instance
	// if err := client.StartInstance(ctx, instanceName); err != nil {
	// 	log.Fatal(err)
	// }

	// Stop instance
	// if err := client.StopInstance(ctx, instanceName); err != nil {
	// 	log.Fatal(err)
	// }

	// // Reboot instance
	// if err := client.RebootInstance(ctx, instanceName); err != nil {
	// 	log.Fatal(err)
	// }

}
