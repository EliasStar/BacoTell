package provider

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionProvider interface {
	GetPrefix() (string, error)
	GetApplicationCommands() ([]Command, error)
	GetMessageComponents() ([]Component, error)
	// GetModals()
}

type Command interface {
	CommandData() (discordgo.ApplicationCommand, error)
	Execute(InteractionProxy) error
	//Autocomplete(proxy) error
}

type Component interface {
	CustomId() (string, error)
	Handle(InteractionProxy) error
}

type InteractionProxy interface {
	Defer(ephemeral, suppressEmbeds, tts bool) error
	Respond(message string, ephemeral, suppressEmbeds, tts bool) error

	Followup(message string, ephemeral, suppressEmbeds, tts bool) (string, error)

	Edit(id string, message string) error
	Delete(id string) error
}

type interactionProvider struct {
	prefix     string
	commands   []Command
	components []Component
}

var _ InteractionProvider = interactionProvider{}

func NewInteractionProvider(prefix string, commands []Command, components []Component) InteractionProvider {
	return interactionProvider{
		prefix:     prefix,
		commands:   commands,
		components: components,
	}
}

func (i interactionProvider) GetPrefix() (string, error) {
	return i.prefix, nil
}

func (i interactionProvider) GetApplicationCommands() ([]Command, error) {
	return i.commands, nil
}

func (i interactionProvider) GetMessageComponents() ([]Component, error) {
	return i.components, nil
}
