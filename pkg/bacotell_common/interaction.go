package bacotell_common

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionProxy interface {
	Defer(ephemeral bool) error
	Respond(message Response, ephemeral bool) error
	Modal(customId, title string, components ...discordgo.MessageComponent) error

	Followup(message Response, ephemeral bool) (string, error)

	Edit(id string, message Response) error
	Delete(id string) error

	GuildID() (string, error)
	ChannelID() (string, error)

	UserLocale() (discordgo.Locale, error)
	GuildLocale() (discordgo.Locale, error)

	User() (*discordgo.User, error)
	Member() (*discordgo.Member, error)

	Message() (*discordgo.Message, error)

	Permissions() (int64, error)
}

type Response struct {
	Content         string
	SuppressEmbeds  bool
	TTS             bool
	AllowedMentions discordgo.MessageAllowedMentions
	Components      []discordgo.MessageComponent
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
}
