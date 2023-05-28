package discord

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/EliasStar/BacoTell/internal/common"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	logger   = common.GetLogger().Named("discord")
	handlers = []any{
		onConnect,
		onDisconnect,
		onReady,

		interactionOnConnect,
		interactionOnGuildCreate,
		interactionOnInteractionCreate,
	}
)

func Connect() {
	session, err := discordgo.New("Bot " + viper.GetString(common.ConfigBotToken))
	if err != nil {
		panic(err)
	}

	session.UserAgent = "BacoTell " + common.Version

	logger.Info("registering handlers")
	for _, handler := range handlers {
		session.AddHandler(handler)
	}

	logger.Info("opening session")
	if session.Open() != nil {
		panic(err)
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	logger.Info("closing session")
	if session.Close() != nil {
		panic(err)
	}
}

func onConnect(*discordgo.Session, *discordgo.Connect) {
	logger.Info("connected to discord gateway")
}

func onDisconnect(*discordgo.Session, *discordgo.Disconnect) {
	logger.Info("disconnected from discord gateway")
}

func onReady(_ *discordgo.Session, evt *discordgo.Ready) {
	logger.Info("logged in", "username", evt.User.Username, "tag", evt.User.Discriminator)
}
