package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

const exampleComponentCustomID = pluginID + "_component"

type ExampleComponent struct{}

var _ common.Component = ExampleComponent{}

// CustomID implements bacotell_common.Component.
func (c ExampleComponent) CustomID() (string, error) {
	return exampleComponentCustomID, nil
}

// Handle implements bacotell_common.Component.
func (c ExampleComponent) Handle(proxy common.HandleProxy) error {
	return proxy.Modal(exampleModalCustomID, "Example Modal", discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.TextInput{
				CustomID: exampleModalComponentCustomID,
				Label:    "Message",
				Style:    discordgo.TextInputShort,
			},
		},
	})
}
