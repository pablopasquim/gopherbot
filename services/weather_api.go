package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const apiURL = "https://wttr.in/%s?0"

func WeatherMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	parts := strings.SplitN(m.Content, " ", 2)
	if len(parts) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Por favor, especifique uma cidade. Exemplo: `!tempo Curitiba`")
		return
	}

	city := strings.TrimSpace(parts[1])
	url := fmt.Sprintf(apiURL, city)

	resp, err := http.Get(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao buscar os dados")
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao ler a respota")
		log.Fatal(err)
		return
	}

	message := string(body)
	s.ChannelMessageSend(m.ChannelID, message)
}
