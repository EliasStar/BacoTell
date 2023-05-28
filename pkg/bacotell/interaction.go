package bacotell

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionProxy interface {
	Defer(ephemeral, suppressEmbeds, tts bool) error
	Respond(message Response, ephemeral, suppressEmbeds, tts bool) error

	Followup(message Response, ephemeral, suppressEmbeds, tts bool) (string, error)

	Edit(id string, message Response) error
	Delete(id string) error

	//Modal(id, title string) error
	//Locale() string error
}

type Response struct {
	Content         string
	AllowedMentions discordgo.MessageAllowedMentions
	Components      []discordgo.MessageComponent
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
}
