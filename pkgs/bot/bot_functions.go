package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/glucktek/gfc-d-bot/pkgs/lightsail"
)

type Bot struct {
	Session   *discordgo.Session
	Token     string
	GuildID   string
	AdminRole string // The role ID that's allowed to use the bot
}

func New(token, guildID, adminRole string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	return &Bot{
		Session:   dg,
		Token:     token,
		GuildID:   guildID,
		AdminRole: adminRole,
	}, nil
}

// HasRequiredRole checks if the user has the admin role
func (b *Bot) HasRequiredRole(member *discordgo.Member) bool {
	for _, roleID := range member.Roles {
		if roleID == b.AdminRole {
			return true
		}
	}
	return false
}

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "gfcbot",
		Description: "GFC Bot commands",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "server",
				Description: "Manage the server",
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "start",
						Description: "Start the server",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
					{
						Name:        "stop",
						Description: "Stop the server",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
					{
						Name:        "reboot",
						Description: "Reboot the server",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
					{
						Name:        "status",
						Description: "Get server status",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
					{
						Name:        "check-website",
						Description: "Check if the website is returning a 200 success code",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
				},
			},
			{
				Name:        "bot",
				Description: "Bot management commands",
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "status",
						Description: "Check bot status",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
				},
			},
		},
	},
}

func (b *Bot) handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check for required role
	if !b.HasRequiredRole(i.Member) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "🚫 You need the required role to use this bot!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// Initialize the Lightsail client
	client, err := lightsail.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// HACK: Using hardcaded instance name for now
	instanceName := "GreaterFaithChurchSite"

	data := i.ApplicationCommandData()

	if data.Name == "gfcbot" {
		group := data.Options[0].Name
		subcommand := data.Options[0].Options[0].Name

		switch group {
		case "server":
			switch subcommand {
			case "start":
				//Start instance
				if err := client.StartInstance(ctx, instanceName); err != nil {
					fmt.Println("Error starting instance:", err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "❌ Failed to start the server: " + err.Error(),
						},
					})
					return
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "✅ Server start command successfully sent!",
					},
				})
				return

			case "stop":
				if err := client.StopInstance(ctx, instanceName); err != nil {
					fmt.Println("Error stopping instance:", err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "❌ Failed to stop the server: " + err.Error(),
						},
					})
					return
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "✅ Server stop command successfully sent, Please wait 5 minutes!",
					},
				})
				return

			case "reboot":
				if err := client.RebootInstance(ctx, instanceName); err != nil {
					fmt.Println("Error rebooting instance:", err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "❌ Failed to reboot the server: " + err.Error(),
						},
					})
					return
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "✅ Server rebooted successfully, please wait 5 mins and check status again!",
					},
				})
				return

			case "status":
				// Get status
				state, err := client.GetInstanceState(ctx, instanceName)
				if err != nil {
					log.Fatal(err)
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Current Instance state: " + state,
					},
				})

			case "check-website":
				// Check website status
				websiteURL := "https://greaterfaithchurch.org"
				resp, err := http.Get(websiteURL)
				if err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "❌ Failed to reach the website: \n" + err.Error(),
						},
					})
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == 200 {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "✅ Website is up\nreturned a 200 success code!",
						},
					})
				} else {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: fmt.Sprintf("⚠️ Website returned a non-200 status code: \n%d", resp.StatusCode),
						},
					})
				}
			}

		case "bot":
			switch subcommand {
			case "status":
				fmt.Println("Status command recieved: Bot is running normally")
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "🟢 Bot is running normally!",
					},
				})
			}
		}
	}
}

func (b *Bot) registerCommands() error {
	for _, cmd := range commands {
		_, err := b.Session.ApplicationCommandCreate(b.Session.State.User.ID, b.GuildID, cmd)
		if err != nil {
			return fmt.Errorf("error creating command %s: %w", cmd.Name, err)
		}
	}
	return nil
}

func (b *Bot) removeCommands() error {
	registeredCommands, err := b.Session.ApplicationCommands(b.Session.State.User.ID, b.GuildID)
	if err != nil {
		return fmt.Errorf("error getting registered commands: %w", err)
	}

	for _, cmd := range registeredCommands {
		err := b.Session.ApplicationCommandDelete(b.Session.State.User.ID, b.GuildID, cmd.ID)
		if err != nil {
			return fmt.Errorf("error removing command %s: %w", cmd.Name, err)
		}
	}
	return nil
}

func (b *Bot) Start() error {
	b.Session.AddHandler(b.handleCommands)

	if err := b.Session.Open(); err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}

	if err := b.registerCommands(); err != nil {
		return fmt.Errorf("error registering commands: %w", err)
	}

	fmt.Printf("Bot is running. Press CTRL+C to exit.\n")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := b.removeCommands(); err != nil {
		fmt.Printf("Error removing commands: %v\n", err)
	}

	return b.Session.Close()
}
