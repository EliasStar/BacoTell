package bacotell_common

import "github.com/bwmarrin/discordgo"

type Command interface {
	CommandData() (discordgo.ApplicationCommand, error)
	Execute(ExecuteProxy) error
	//Autocomplete(AutocompleteProxy) error TODO
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

// type AutocompleteProxy interface {} TODO
