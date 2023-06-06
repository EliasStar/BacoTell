package bacotell_common

import "github.com/bwmarrin/discordgo"

// A Command represents a chat/message/user application command.
//
// It provides all information necessary to register the command to a guild,
// as well as the command handler implementation.
type Command interface {
	// Data returns a discordgo.ApplicationCommand with all required fields set.
	//
	// Using this BacoTell registers this command to all guilds so it can be executed by users.
	// To ensure that commands from different plugins do not conflict with each other the command name has to be prefixed with the plugin id.
	// While it is possible for this method to return an error, it is in most cases regarded as a programming error rather than typical behavior.
	Data() (discordgo.ApplicationCommand, error)

	// Execute gets called whenever a user runs this command.
	Execute(ExecuteProxy) error

	// Autocomplete gets called when a user is typing this command and the autocomplete flag was set for one of the command options.
	Autocomplete(AutocompleteProxy) error
}

// An ExecuteProxy provides methods for responding to the interaction as well as accessing interaction data and command options during command execution.
type ExecuteProxy interface {
	InteractionProxy

	// StringOption retrieves the value of a string option by name. Returns an error if it was not set or the value is not a string.
	StringOption(name string) (string, error)

	// IntegerOption retrieves the value of an integer option by name. Returns an error if it was not set or the value is not a integer.
	IntegerOption(name string) (int64, error)

	// NumberOption retrieves the value of a number option by name. Returns an error if it was not set or the value is not a number.
	NumberOption(name string) (float64, error)

	// BooleanOption retrieves the value of a boolean option by name. Returns an error if it was not set or the value is not a boolean.
	BooleanOption(name string) (bool, error)

	// UserOption retrieves the value of a user option by name. Returns an error if it was not set or the value is not a user.
	UserOption(name string) (*discordgo.User, error)

	// RoleOption retrieves the value of a role option by name. Returns an error if it was not set or the value is not a role.
	RoleOption(name string) (*discordgo.Role, error)

	// ChannelOption retrieves the value of a channel option by name. Returns an error if it was not set or the value is not a channel.
	ChannelOption(name string) (*discordgo.Channel, error)

	// AttachmentOption retrieves the value of an attachment option by name. Returns an error if it was not set or the value is not a attachment.
	AttachmentOption(name string) (*discordgo.MessageAttachment, error)
}

// An AutocompleteProxy provides methods for responding to the interaction as well as accessing the command option to be autocompleted.
type AutocompleteProxy interface {
	// Respond sends the user the autocomplete suggestions.
	Respond(choices ...*discordgo.ApplicationCommandOptionChoice) error

	// FocusedOption returns the name and current partial value of the option for which autocompletion is requested.
	FocusedOption() (string, any, error)
}
