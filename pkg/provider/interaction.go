package provider

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionProvider interface {
	Prefix() (string, error)
	ApplicationCommands() ([]Command, error)
	MessageComponents() ([]Component, error)
	// Modals() ([]Modal, error)
}

type Command interface {
	CommandData() (discordgo.ApplicationCommand, error)
	Execute(ExecuteProxy) error
	//Autocomplete(AutocompleteProxy) error
}

type Component interface {
	CustomId() (string, error)
	Handle(HandleProxy) error
}

// type Modal interface {
// 	CustomId() (string, error)
// 	Submit(SubmitProxy) error
// }

type InteractionProxy interface {
	Defer(ephemeral, suppressEmbeds, tts bool) error
	Respond(message Response, ephemeral, suppressEmbeds, tts bool) error

	Followup(message Response, ephemeral, suppressEmbeds, tts bool) (string, error)

	Edit(id string, message Response) error
	Delete(id string) error

	//Modal(id, title string) error

	//Locale() string error
}

type ExecuteProxy interface {
	InteractionProxy

	StringOption(name string) (string, error)
	IntegerOption(name string) (int64, error)
	NumberOption(name string) (float64, error)
	BooleanOption(name string) (bool, error)

	UserOption(name string) (*discordgo.User, error)
	RoleOption(name string) (*discordgo.Role, error)
	ChannelOption(name string) (*discordgo.Channel, error)

	AttachmentOption(name string) (*discordgo.MessageAttachment, error)
}

type HandleProxy interface {
	InteractionProxy
}

// type AutocompleteProxy interface {
// }

// type SubmitProxy interface {
// 	InteractionProxy
// }

type Response struct {
	Content         string
	AllowedMentions discordgo.MessageAllowedMentions
	Components      []discordgo.MessageComponent
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
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

func (i interactionProvider) Prefix() (string, error) {
	return i.prefix, nil
}

func (i interactionProvider) ApplicationCommands() ([]Command, error) {
	return i.commands, nil
}

func (i interactionProvider) MessageComponents() ([]Component, error) {
	return i.components, nil
}
