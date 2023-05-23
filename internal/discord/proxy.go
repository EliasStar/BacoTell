package discord

import (
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
)

type interactionProxy struct {
	session     *discordgo.Session
	interaction *discordgo.Interaction
}

var _ provider.InteractionProxy = interactionProxy{}

// Defer implements provider.InteractionProxy
func (p interactionProxy) Defer(ephemeral bool, suppressEmbeds bool, tts bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if suppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: flags,
			TTS:   tts,
		},
	})
}

// Respond implements provider.InteractionProxy
func (p interactionProxy) Respond(message string, ephemeral bool, suppressEmbeds bool, tts bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if suppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   flags,
			TTS:     tts,
		},
	})
}

// Followup implements provider.InteractionProxy
func (p interactionProxy) Followup(message string, ephemeral bool, suppressEmbeds bool, tts bool) (string, error) {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if suppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	msg, err := p.session.FollowupMessageCreate(p.interaction, true, &discordgo.WebhookParams{
		Content: message,
		Flags:   flags,
		TTS:     tts,
	})

	if err != nil {
		return "", err
	}

	return msg.ID, nil
}

// Edit implements provider.InteractionProxy
func (p interactionProxy) Edit(id string, message string) (err error) {
	msg := &discordgo.WebhookEdit{
		Content: &message,
	}

	if id == "" {
		_, err = p.session.InteractionResponseEdit(p.interaction, msg)
	} else {
		_, err = p.session.FollowupMessageEdit(p.interaction, id, msg)
	}

	return
}

// Delete implements provider.InteractionProxy
func (p interactionProxy) Delete(id string) error {
	if id == "" {
		return p.session.InteractionResponseDelete(p.interaction)
	}

	return p.session.FollowupMessageDelete(p.interaction, id)
}

type executeProxy struct {
	interactionProxy
}

var _ provider.ExecuteProxy = executeProxy{}

// StringOption implements provider.ExecuteProxy
func (p executeProxy) StringOption(name string) (string, error) {
	panic("unimplemented")
}

// IntegerOption implements provider.ExecuteProxy
func (p executeProxy) IntegerOption(name string) (int64, error) {
	panic("unimplemented")
}

// NumberOption implements provider.ExecuteProxy
func (p executeProxy) NumberOption(name string) (float64, error) {
	panic("unimplemented")
}

// BooleanOption implements provider.ExecuteProxy
func (p executeProxy) BooleanOption(name string) (bool, error) {
	panic("unimplemented")
}

// UserOption implements provider.ExecuteProxy
func (executeProxy) UserOption(name string) (discordgo.User, error) {
	panic("unimplemented")
}

// RoleOption implements provider.ExecuteProxy
func (executeProxy) RoleOption(name string) (discordgo.Role, error) {
	panic("unimplemented")
}

// ChannelOption implements provider.ExecuteProxy
func (executeProxy) ChannelOption(name string) (discordgo.Channel, error) {
	panic("unimplemented")
}

// AttachmentOption implements provider.ExecuteProxy
func (executeProxy) AttachmentOption(name string) (discordgo.MessageAttachment, error) {
	panic("unimplemented")
}

type handleProxy struct {
	interactionProxy
}

var _ provider.HandleProxy = handleProxy{}
