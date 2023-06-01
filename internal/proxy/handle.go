package proxy

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

type handleProxy struct {
	interactionProxy
}

var _ common.HandleProxy = handleProxy{}

func NewHandleProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.HandleProxy {
	return handleProxy{
		interactionProxy: interactionProxy{
			session:     session,
			interaction: interaction,
		},
	}
}

// Defer implements bacotell_common.HandleProxy
func (p handleProxy) Defer(ephemeral bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
		Data: &discordgo.InteractionResponseData{
			Flags: flags,
		},
	})
}

// Respond implements bacotell_common.HandleProxy
func (p handleProxy) Respond(message common.Response, ephemeral bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if message.SuppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:         message.Content,
			TTS:             message.TTS,
			AllowedMentions: &message.AllowedMentions,
			Components:      message.Components,
			Embeds:          message.Embeds,
			Files:           message.Files,
			Flags:           flags,
		},
	})
}
