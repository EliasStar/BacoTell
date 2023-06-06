package proxy

import (
	"errors"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

// The default implementation for bacotell_common.ExecuteProxy.
type executeProxy struct {
	interactionProxy
}

// executeProxy implements bacotell_common.ExecuteProxy.
var _ common.ExecuteProxy = executeProxy{}

// NewExecuteProxy returns a new ExecuteProxy implementation.
func NewExecuteProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.ExecuteProxy {
	return executeProxy{
		interactionProxy: interactionProxy{
			session:     session,
			interaction: interaction,
		},
	}
}

// StringOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) StringOption(name string) (string, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return "", errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionString {
		return "", errors.New("option is not a string")
	}

	return option.StringValue(), nil
}

// IntegerOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) IntegerOption(name string) (int64, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return 0, errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionInteger {
		return 0, errors.New("option is not an integer")
	}

	return option.IntValue(), nil
}

// NumberOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) NumberOption(name string) (float64, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return 0, errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionNumber {
		return 0, errors.New("option is not a number")
	}

	return option.FloatValue(), nil
}

// BooleanOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) BooleanOption(name string) (bool, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return false, errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionBoolean {
		return false, errors.New("option is not a boolean")
	}

	return option.BoolValue(), nil
}

// UserOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) UserOption(name string) (*discordgo.User, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return nil, errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionUser && option.Type != discordgo.ApplicationCommandOptionMentionable {
		return nil, errors.New("option is not a user nor a mentionable")
	}

	return option.UserValue(p.session), nil
}

// RoleOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) RoleOption(name string) (*discordgo.Role, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return nil, errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionRole && option.Type != discordgo.ApplicationCommandOptionMentionable {
		return nil, errors.New("option is not a role nor a mentionable")
	}

	return option.RoleValue(p.session, p.interaction.GuildID), nil
}

// ChannelOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) ChannelOption(name string) (*discordgo.Channel, error) {
	option := _findOption(p.interaction.ApplicationCommandData().Options, name)

	if option == nil {
		return nil, errors.New("option was not set")
	}

	if option.Type != discordgo.ApplicationCommandOptionChannel {
		return nil, errors.New("option is not a channel")
	}

	return option.ChannelValue(p.session), nil
}

// AttachmentOption implements bacotell_common.ExecuteProxy.
func (p executeProxy) AttachmentOption(name string) (*discordgo.MessageAttachment, error) {
	data := p.interaction.ApplicationCommandData()
	option := _findOption(data.Options, name)

	if option == nil {
		return nil, errors.New("option was not set")
	}

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

// _findOption returns the option with the specified name if it exists.
func _findOption(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, opt := range options {
		if opt.Name == name {
			return opt
		}

		if opt.Options != nil {
			return _findOption(opt.Options, name)
		}
	}

	return nil
}
