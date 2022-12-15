package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var token = "BOT-TOKEN"
var Dg *discordgo.Session

func InitSession() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	err = dg.Open()
	if err != nil {
		fmt.Printf("Discord session couldn't opened: ", err)
		l.Fatalf("Discord session couldn't opened: ", err)
	}

	Dg = dg

}

func ConnectToDc() {
	InitSession()
	//Register messageCreate func

	Dg.AddHandler(messageCreate)

	Dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = Dg.Close()
	l.Printf("[INFO] Bot stopped")

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if strings.HasPrefix(m.Content, "!") {
		response := gpt(m.Content)

		_, err := s.ChannelMessageSend(m.ChannelID, response)

		if err != nil {
			fmt.Println("error sending DM message:", err)

		}

	}

}
