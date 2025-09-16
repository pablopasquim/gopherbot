package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/pablopasquim/gopherbot/internal/services"
)

func Start() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal(err)
	}

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "hello" {
			send, err := s.ChannelMessageSend(m.ChannelID, "Hello, World!")
			if err != nil {
				log.Fatal(err)
			}
			log.Println(send)
		}

		if strings.HasPrefix(m.Content, "!tempo") {
			services.WeatherMessage(s, m)
			return
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sess *discordgo.Session) {
		err := sess.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sess)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Interrupt)
	<-sc
}
