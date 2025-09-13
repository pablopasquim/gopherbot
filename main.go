package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// IMPORTANTE: Substitua esta variável pelo seu token.
// Nunca a compartilhe em repositórios públicos!
var BotToken = "SEU_TOKEN_DO_BOT_AQUI"

func main() {
	// Cria uma nova sessão do Discord.
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Erro ao criar a sessão do Discord:", err)
		return
	}

	// Registra uma função para lidar com o evento de 'mensagem criada'.
	dg.AddHandler(messageCreate)

	// Abre a conexão com o Discord.
	err = dg.Open()
	if err != nil {
		fmt.Println("Erro ao abrir a conexão:", err)
		return
	}

	fmt.Println("Bot está online! Pressione CTRL+C para sair.")

	// Mantém o programa rodando até receber um sinal de interrupção (CTRL+C).
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Fecha a sessão do Discord de forma limpa.
	dg.Close()
}

// Esta função é ativada toda vez que uma nova mensagem é enviada em um canal.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignora mensagens do próprio bot para evitar loops infinitos.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Se a mensagem for "!ping", o bot responde com "Pong!".
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
