package bacotell_common

import (
	"github.com/bwmarrin/discordgo"
)

// An InteractionProxy provides methods for responding to the interaction as well as accessing interaction data.
type InteractionProxy interface {
	// Defer defers the response to the interaction which shows a possibly ephemeral loading message to the user.
	Defer(ephemeral bool) error

	// Respond responds to the user with the specified message and ephemerality.
	Respond(message Response, ephemeral bool) error

	// Modal shows a modal to the user with the specified custom id, title, and components.
	Modal(customId, title string, components ...discordgo.MessageComponent) error

	// Followup sends a follow-up message to the user with the specified message and ephemerality. The message id of the follow-up message is returned.
	Followup(message Response, ephemeral bool) (string, error)

	// Edit edits the message with the specified id. If id is the empty string it edits the original interaction response.
	Edit(id string, message Response) error

	// Delete deletes the message with the specified id. If id is the empty string it deletes the original interaction response.
	Delete(id string) error

	// GuildID returns the ID of the guild where the interaction occurred, if it was sent from a guild channel. Otherwise an empty string.
	GuildID() (string, error)

	// ChannelID returns the ID of the channel where the interaction occurred.
	ChannelID() (string, error)

	// UserLocale returns the locale of the user who triggered the interaction.
	UserLocale() (discordgo.Locale, error)

	// GuildLocale returns the locale of the guild where the interaction occurred, if it was sent from a guild channel. Otherwise it returns the user's locale.
	GuildLocale() (discordgo.Locale, error)

	// User returns the user who triggered the interaction, if it was sent from a DM channel. Otherwise nil.
	User() (*discordgo.User, error)

	// Member returns the member who triggered the interaction, if it was sent from a guild channel. Otherwise nil.
	Member() (*discordgo.Member, error)

	// Message returns the message on which interaction was used. Only available for message application commands and component, otherwise it will be nil.
	Message() (*discordgo.Message, error)

	// Permissions returns a bitwise set of permissions the bot has within the channel the interaction was sent from.
	Permissions() (int64, error)
}

// Response represents a response message for an interaction.
type Response struct {
	// Content is the text content of the message.
	Content string

	// SuppressEmbeds specifies whether to suppress generated embeds in the message for e.g. links.
	SuppressEmbeds bool

	// TTS specifies whether the message should be spoken to the user.
	TTS bool

	// AllowedMentions specifies which mentions should be parsed by Discord.
	AllowedMentions discordgo.MessageAllowedMentions

	// Components specifies the message components to be included on the response message.
	Components []discordgo.MessageComponent

	// Embeds specify the message embeds to be included in the response message.
	Embeds []*discordgo.MessageEmbed

	// Files specify the files to be attached to the response message.
	Files []*discordgo.File
}
