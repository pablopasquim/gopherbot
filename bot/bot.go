package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token = "MTQxNjUyNzMzODkyMTAwNTI0MA.GF6U9X.WPnvOgJuuJhFU4LdeDLeND0GhEiSIuEHJomzsg"

func Start() {
	// Cria uma nova sessão do Discord.
	sess, err := discordgo.New("Bot " + Token)
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
