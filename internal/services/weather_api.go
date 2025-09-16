package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pablopasquim/gopherbot/internal/models"
)

const apiURL = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

func WeatherMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	parts := strings.SplitN(m.Content, " ", 2)
	if len(parts) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Por favor, especifique uma cidade. Exemplo: `!tempo Curitiba`")
		return
	}

	city := strings.TrimSpace(parts[1])
	apiKey := os.Getenv("WEATHER_API")
	url := fmt.Sprintf(apiURL, city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao buscar os dados")
		log.Println("Erro no http.Get:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao ler a resposta")
		log.Println("Erro no ReadAll:", err)
		return
	}

	var weatherData models.WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao processar os dados do clima.")
		log.Println("Erro no json.Unmarshal:", err)
		return
	}

	if weatherData.Name == "" || len(weatherData.Weather) == 0 {
		s.ChannelMessageSend(m.ChannelID, "❌ Cidade não encontrada. Verifique o nome e tente novamente.")
		return
	}

	message := fmt.Sprintf(
		"🌍 Tempo em %s:\n```\n🌡️ Temperatura: %.1f°C\n💧 Umidade: %d%%\n🌤️ Clima: %s\n```",
		weatherData.Name,
		weatherData.Main.Temp,
		weatherData.Main.Humidity,
		weatherData.Weather[0].Description,
	)

	s.ChannelMessageSend(m.ChannelID, message)
}
