package proxy

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

// The default implementation for bacotell_common.HandleProxy.
type handleProxy struct {
	interactionProxy
}

// handleProxy implements bacotell_common.HandleProxy.
var _ common.HandleProxy = handleProxy{}

// NewHandleProxy returns a new HandleProxy implementation.
func NewHandleProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.HandleProxy {
	return handleProxy{
		interactionProxy: interactionProxy{
			session:     session,
			interaction: interaction,
		},
	}
}

// Defer implements bacotell_common.InteractionProxy, bacotell_common.HandleProxy.
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

// Respond implements bacotell_common.InteractionProxy, bacotell_common.HandleProxy.
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

// ComponentType implements bacotell_common.HandleProxy.
func (p handleProxy) ComponentType() (discordgo.ComponentType, error) {
	return p.interaction.MessageComponentData().ComponentType, nil
}

// SelectedValues implements bacotell_common.HandleProxy.
func (p handleProxy) SelectedValues() ([]string, error) {
	return p.interaction.MessageComponentData().Values, nil
}
