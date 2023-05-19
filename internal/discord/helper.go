package discord

import (
	"github.com/bwmarrin/discordgo"
)

func Defer(session *discordgo.Session, interaction *discordgo.Interaction) error {
	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
}

func DeferEphemeral(session *discordgo.Session, interaction *discordgo.Interaction) error {
	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: 1 << 6,
		},
	})
}

func Respond(session *discordgo.Session, interaction *discordgo.Interaction, content string) error {
	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func RespondEphemeral(session *discordgo.Session, interaction *discordgo.Interaction, content string) error {
	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   1 << 6,
		},
	})
}

func Followup(session *discordgo.Session, interaction *discordgo.Interaction, content string) (*discordgo.Message, error) {
	return session.FollowupMessageCreate(interaction, true, &discordgo.WebhookParams{
		Content: content,
	})
}

func FollowupEphemeral(session *discordgo.Session, interaction *discordgo.Interaction, content string) (*discordgo.Message, error) {
	return session.FollowupMessageCreate(interaction, true, &discordgo.WebhookParams{
		Content: content,
		Flags:   1 << 6,
	})
}

type Message struct {
	Content         string
	Components      []discordgo.MessageComponent
	Embeds          []discordgo.MessageEmbed
	AllowedMentions discordgo.MessageAllowedMentions
	Files           []discordgo.File
}
