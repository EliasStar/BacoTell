package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

const exampleCommandName = pluginID + "_command"
const exampleCommandOptionMessage = "message"

type ExampleCommand struct{}

var _ common.Command = ExampleCommand{}

// Data implements bacotell_common.Command.
func (c ExampleCommand) Data() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        exampleCommandName,
		Description: "Example Command",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         exampleCommandOptionMessage,
				Description:  "Example Option",
				Autocomplete: true,
			},
		},
	}, nil
}

// Execute implements bacotell_common.Command.
func (c ExampleCommand) Execute(proxy common.ExecuteProxy) error {
	logger.Info("Executing ExampleCommand!")

	value, err := proxy.StringOption(exampleCommandOptionMessage)
	reply := "User wrote: " + value
	if err != nil {
		reply = "User wrote nothing!"
	}

	return proxy.Respond(
		common.Response{
			Content: reply,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.Button{
						CustomID: exampleComponentCustomID,
						Label:    "Example Button",
						Style:    discordgo.PrimaryButton,
					},
				}},
			},
		},
		false,
	)
}

// Autocomplete implements bacotell_common.Command.
func (c ExampleCommand) Autocomplete(proxy common.AutocompleteProxy) error {
	logger.Info("Autocompleting ExampleCommand!")

	_, value, err := proxy.FocusedOption()
	if err != nil {
		return nil
	}

	return proxy.Respond(
		&discordgo.ApplicationCommandOptionChoice{
			Name:  value.(string) + " and " + value.(string),
			Value: value.(string) + " and " + value.(string),
		},
	)
}
