package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	Token   string
}

func New(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}
	return &Bot{
		Session: dg,
		Token:   token,
	}, nil
}

func (b *Bot) Start() error {
	b.Session.AddHandler(b.echoHandler)

	err := b.Session.Open()
	if err != nil {
		return fmt.Errorf("error opening Discord session: %w", err)
	}
	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for a signal to quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	b.Session.Close()
	return nil
}

func (b *Bot) echoHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Echo command: !echo <message>
	if len(m.Content) > 6 && m.Content[:6] == "!echo " {
		msg := m.Content[6:]
		s.ChannelMessageSend(m.ChannelID, msg)
	}
}
