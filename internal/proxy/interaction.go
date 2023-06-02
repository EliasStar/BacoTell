package proxy

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

type interactionProxy struct {
	session     *discordgo.Session
	interaction *discordgo.Interaction
}

var _ common.InteractionProxy = interactionProxy{}

func NewInteractionProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.InteractionProxy {
	return interactionProxy{
		session:     session,
		interaction: interaction,
	}
}

// Defer implements bacotell_common.InteractionProxy
func (p interactionProxy) Defer(ephemeral bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: flags,
		},
	})
}

// Respond implements bacotell_common.InteractionProxy
func (p interactionProxy) Respond(message common.Response, ephemeral bool) error {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if message.SuppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
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

// Modal implements bacotell_common.InteractionProxy
func (p interactionProxy) Modal(customId string, title string, components ...discordgo.MessageComponent) error {
	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID:   customId,
			Title:      title,
			Components: components,
		},
	})
}

// Followup implements bacotell_common.InteractionProxy
func (p interactionProxy) Followup(message common.Response, ephemeral bool) (string, error) {
	var flags discordgo.MessageFlags

	if ephemeral {
		flags |= discordgo.MessageFlagsEphemeral
	}

	if message.SuppressEmbeds {
		flags |= discordgo.MessageFlagsSuppressEmbeds
	}

	msg, err := p.session.FollowupMessageCreate(p.interaction, true, &discordgo.WebhookParams{
		Content:         message.Content,
		TTS:             message.TTS,
		AllowedMentions: &message.AllowedMentions,
		Components:      message.Components,
		Embeds:          message.Embeds,
		Files:           message.Files,
		Flags:           flags,
	})

	if err != nil {
		return "", err
	}

	return msg.ID, nil
}

// Edit implements bacotell_common.InteractionProxy
func (p interactionProxy) Edit(id string, message common.Response) (err error) {
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

// Delete implements bacotell_common.InteractionProxy
func (p interactionProxy) Delete(id string) error {
	if id == "" {
		return p.session.InteractionResponseDelete(p.interaction)
	}

	return p.session.FollowupMessageDelete(p.interaction, id)
}

// GuildID implements bacotell_common.InteractionProxy
func (p interactionProxy) GuildID() (string, error) {
	return p.interaction.GuildID, nil
}

// ChannelID implements bacotell_common.InteractionProxy
func (p interactionProxy) ChannelID() (string, error) {
	return p.interaction.ChannelID, nil
}

// UserLocale implements bacotell_common.InteractionProxy
func (p interactionProxy) UserLocale() (discordgo.Locale, error) {
	return p.interaction.Locale, nil
}

// GuildLocale implements bacotell_common.InteractionProxy
func (p interactionProxy) GuildLocale() (discordgo.Locale, error) {
	if p.interaction.GuildLocale != nil {
		return *p.interaction.GuildLocale, nil
	}

	return p.interaction.Locale, nil
}

// User implements bacotell_common.InteractionProxy
func (p interactionProxy) User() (*discordgo.User, error) {
	return p.interaction.User, nil
}

// Member implements bacotell_common.InteractionProxy
func (p interactionProxy) Member() (*discordgo.Member, error) {
	return p.interaction.Member, nil
}

// Message implements bacotell_common.InteractionProxy
func (p interactionProxy) Message() (*discordgo.Message, error) {
	return p.interaction.Message, nil
}

// Permissions implements bacotell_common.InteractionProxy
func (p interactionProxy) Permissions() (int64, error) {
	return p.interaction.AppPermissions, nil
}
