package main

import (
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
)

type TestCommand struct{}

var _ provider.Command = TestCommand{}

func (TestCommand) CommandData() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "test_cmd",
		Description: "Test command from test plugin.",
	}, nil
}

func (TestCommand) Execute(provider.InteractionProxy) error {
	logger.Info("execute command")
	return nil
}

func (TestCommand) Autocomplete(provider.InteractionProxy) error {
	logger.Info("execute command")
	return nil
}
