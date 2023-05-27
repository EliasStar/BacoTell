package discord

import (
	"errors"

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
func (p interactionProxy) Respond(message provider.Response, ephemeral bool, suppressEmbeds bool, tts bool) error {
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
			Content:         message.Content,
			AllowedMentions: &message.AllowedMentions,
			Components:      message.Components,
			Embeds:          message.Embeds,
			Files:           message.Files,
			Flags:           flags,
			TTS:             tts,
		},
	})
}

// Followup implements provider.InteractionProxy
func (p interactionProxy) Followup(message provider.Response, ephemeral bool, suppressEmbeds bool, tts bool) (string, error) {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if suppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	msg, err := p.session.FollowupMessageCreate(p.interaction, true, &discordgo.WebhookParams{
		Content:         message.Content,
		AllowedMentions: &message.AllowedMentions,
		Components:      message.Components,
		Embeds:          message.Embeds,
		Files:           message.Files,
		Flags:           flags,
		TTS:             tts,
	})

	if err != nil {
		return "", err
	}

	return msg.ID, nil
}

// Edit implements provider.InteractionProxy
func (p interactionProxy) Edit(id string, message provider.Response) (err error) {
	msg := &discordgo.WebhookEdit{
		Content:         &message.Content,
		AllowedMentions: &message.AllowedMentions,
		Components:      &message.Components,
		Embeds:          &message.Embeds,
		Files:           message.Files,
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
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionString {
		return "", errors.New("option is not a string")
	}

	return option.StringValue(), nil
}

// IntegerOption implements provider.ExecuteProxy
func (p executeProxy) IntegerOption(name string) (int64, error) {
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionInteger {
		return 0, errors.New("option is not an integer")
	}

	return option.IntValue(), nil
}

// NumberOption implements provider.ExecuteProxy
func (p executeProxy) NumberOption(name string) (float64, error) {
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionNumber {
		return 0, errors.New("option is not a number")
	}

	return option.FloatValue(), nil
}

// BooleanOption implements provider.ExecuteProxy
func (p executeProxy) BooleanOption(name string) (bool, error) {
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionBoolean {
		return false, errors.New("option is not a boolean")
	}

	return option.BoolValue(), nil
}

// UserOption implements provider.ExecuteProxy
func (p executeProxy) UserOption(name string) (*discordgo.User, error) {
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionUser && option.Type != discordgo.ApplicationCommandOptionMentionable {
		return nil, errors.New("option is not a user nor a mentionable")
	}

	return option.UserValue(p.session), nil
}

// RoleOption implements provider.ExecuteProxy
func (p executeProxy) RoleOption(name string) (*discordgo.Role, error) {
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionRole && option.Type != discordgo.ApplicationCommandOptionMentionable {
		return nil, errors.New("option is not a role nor a mentionable")
	}

	return option.RoleValue(p.session, p.interaction.GuildID), nil
}

// ChannelOption implements provider.ExecuteProxy
func (p executeProxy) ChannelOption(name string) (*discordgo.Channel, error) {
	option := findOption(p.interaction.ApplicationCommandData().Options, name)
	if option.Type != discordgo.ApplicationCommandOptionChannel {
		return nil, errors.New("option is not a channel")
	}

	return option.ChannelValue(p.session), nil
}

// AttachmentOption implements provider.ExecuteProxy
func (p executeProxy) AttachmentOption(name string) (*discordgo.MessageAttachment, error) {
	data := p.interaction.ApplicationCommandData()
	option := findOption(data.Options, name)
	if option.Type != discordgo.ApplicationCommandOptionAttachment {
		return nil, errors.New("option is not an attachment")
	}

	id, ok := option.Value.(string)
	if !ok {
		return nil, errors.New("cannot get attachment")
	}

	attachment, ok := data.Resolved.Attachments[id]
	if !ok {
		return nil, errors.New("cannot get attachment")
	}

	return attachment, nil
}

func findOption(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, opt := range options {
		if opt.Name == name {
			return opt
		}

		if opt.Options != nil {
			return findOption(opt.Options, name)
		}
	}

	return nil
}

type handleProxy struct {
	interactionProxy
}

var _ provider.HandleProxy = handleProxy{}

// Defer implements provider.HandleProxy
func (p handleProxy) Defer(ephemeral bool, suppressEmbeds bool, tts bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if suppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
		Data: &discordgo.InteractionResponseData{
			Flags: flags,
			TTS:   tts,
		},
	})
}

// Respond implements provider.HandleProxy
func (p handleProxy) Respond(message provider.Response, ephemeral bool, suppressEmbeds bool, tts bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if suppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:         message.Content,
			AllowedMentions: &message.AllowedMentions,
			Components:      message.Components,
			Embeds:          message.Embeds,
			Files:           message.Files,
			Flags:           flags,
			TTS:             tts,
		},
	})
}
