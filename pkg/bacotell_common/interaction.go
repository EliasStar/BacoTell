package bacotell_common

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionProxy interface {
	Defer(ephemeral bool) error
	Respond(message Response, ephemeral bool) error

	Followup(message Response, ephemeral bool) (string, error)

	Edit(id string, message Response) error
	Delete(id string) error

	//Modal(id, title string) error TODO
	//Locale() string error TODO
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
