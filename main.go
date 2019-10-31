package main

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var file []byte

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	f, err := ioutil.ReadFile("bam.png")
	if err != nil {
		log.Panic(err)
	}
	file = f

	client.AddHandler(onMessage)

	err = client.Open()
	if err != nil {
		log.Panic(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = client.Close()
	if err != nil {
		log.Panic(err)
	}
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!bam") {
		_, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Files: []*discordgo.File{
				&discordgo.File{
					Name:        "bam.png",
					ContentType: "image/png",
					Reader:      bytes.NewReader(file),
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
