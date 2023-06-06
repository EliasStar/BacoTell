package codec

import (
	"github.com/EliasStar/BacoTell/internal/proto/discordgopb"
	util "github.com/EliasStar/BacoTell/pkg/bacotell_util"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/protobuf/types/known/structpb"
)

// encodeApplicationCommand encodes a discordgo.ApplicationCommand into a discordgopb.ApplicationCommand.
func encodeApplicationCommand(command *discordgo.ApplicationCommand) *discordgopb.ApplicationCommand {
	if command == nil {
		return nil
	}

	return &discordgopb.ApplicationCommand{
		Id:            command.ID,
		ApplicationId: command.ApplicationID,
		GuildId:       command.GuildID,
		Version:       command.Version,

		Type: uint32(command.Type),

		Name:              command.Name,
		NameLocalizations: _encodeLocalizations(util.NilableVal[discordgo.Locale, string](command.NameLocalizations)),

		DefaultPermission:        command.DefaultPermission,
		DefaultMemberPermissions: command.DefaultMemberPermissions,
		DmPermission:             command.DMPermission,
		Nsfw:                     command.NSFW,

		Description:              command.Description,
		DescriptionLocalizations: _encodeLocalizations(util.NilableVal[discordgo.Locale, string](command.DescriptionLocalizations)),

		Options: encodeApplicationCommandOptions(command.Options),
	}
}

// decodeApplicationCommand decodes a discordgopb.ApplicationCommand into a discordgo.ApplicationCommand.
func decodeApplicationCommand(command *discordgopb.ApplicationCommand) *discordgo.ApplicationCommand {
	if command == nil {
		return nil
	}

	return &discordgo.ApplicationCommand{
		ID:            command.Id,
		ApplicationID: command.ApplicationId,
		GuildID:       command.GuildId,
		Version:       command.Version,

		Type: discordgo.ApplicationCommandType(command.Type),

		Name:              command.Name,
		NameLocalizations: util.NilablePtr[discordgo.Locale, string](_decodeLocalizations(command.NameLocalizations)),

		DefaultPermission:        command.DefaultPermission,
		DefaultMemberPermissions: command.DefaultMemberPermissions,
		DMPermission:             command.DmPermission,
		NSFW:                     command.Nsfw,

		Description:              command.Description,
		DescriptionLocalizations: util.NilablePtr[discordgo.Locale, string](_decodeLocalizations(command.DescriptionLocalizations)),

		Options: decodeApplicationCommandOptions(command.Options),
	}
}

// encodeApplicationCommandOption encodes a discordgo.ApplicationCommandOption into a discordgopb.ApplicationCommandOption.
func encodeApplicationCommandOption(option *discordgo.ApplicationCommandOption) *discordgopb.ApplicationCommandOption {
	if option == nil {
		return nil
	}

	return &discordgopb.ApplicationCommandOption{
		Type: uint32(option.Type),

		Name:              option.Name,
		NameLocalizations: _encodeLocalizations(option.NameLocalizations),

		Description:              option.Description,
		DescriptionLocalizations: _encodeLocalizations(option.DescriptionLocalizations),

		ChannelTypes: _encodeChannelTypes(option.ChannelTypes),
		Required:     option.Required,
		Options:      encodeApplicationCommandOptions(option.Options),

		Autocomplete: option.Autocomplete,
		Choices:      encodeApplicationCommandOptionChoices(option.Choices),

		MinValue: option.MinValue,
		MaxValue: option.MaxValue,

		MinLength: util.PtrConv[int, int32](option.MinLength),
		MaxLength: int32(option.MaxLength),
	}
}

// encodeApplicationCommandOptions encodes multiple discordgo.ApplicationCommandOption into multiple discordgopb.ApplicationCommandOption.
func encodeApplicationCommandOptions(options []*discordgo.ApplicationCommandOption) []*discordgopb.ApplicationCommandOption {
	if options == nil {
		return nil
	}

	result := make([]*discordgopb.ApplicationCommandOption, len(options))

	for i, option := range options {
		result[i] = encodeApplicationCommandOption(option)
	}

	return result
}

// decodeApplicationCommandOption decodes a discordgopb.ApplicationCommandOption into a discordgo.ApplicationCommandOption.
func decodeApplicationCommandOption(option *discordgopb.ApplicationCommandOption) *discordgo.ApplicationCommandOption {
	if option == nil {
		return nil
	}

	return &discordgo.ApplicationCommandOption{
		Type: discordgo.ApplicationCommandOptionType(option.Type),

		Name:              option.Name,
		NameLocalizations: _decodeLocalizations(option.NameLocalizations),

		Description:              option.Description,
		DescriptionLocalizations: _decodeLocalizations(option.DescriptionLocalizations),

		ChannelTypes: _decodeChannelTypes(option.ChannelTypes),
		Required:     option.Required,
		Options:      decodeApplicationCommandOptions(option.Options),

		Autocomplete: option.Autocomplete,
		Choices:      decodeApplicationCommandOptionChoices(option.Choices),

		MinValue: option.MinValue,
		MaxValue: option.MaxValue,

		MinLength: util.PtrConv[int32, int](option.MinLength),
		MaxLength: int(option.MaxLength),
	}
}

// decodeApplicationCommandOptions decodes multiple discordgopb.ApplicationCommandOption into multiple discordgo.ApplicationCommandOption.
func decodeApplicationCommandOptions(options []*discordgopb.ApplicationCommandOption) []*discordgo.ApplicationCommandOption {
	if options == nil {
		return nil
	}

	result := make([]*discordgo.ApplicationCommandOption, len(options))

	for i, option := range options {
		result[i] = decodeApplicationCommandOption(option)
	}

	return result
}

// encodeApplicationCommandOptionChoice encodes a discordgo.ApplicationCommandOptionChoice into a discordgopb.ApplicationCommandOptionChoice.
func encodeApplicationCommandOptionChoice(choice *discordgo.ApplicationCommandOptionChoice) *discordgopb.ApplicationCommandOptionChoice {
	if choice == nil {
		return nil
	}

	value, _ := structpb.NewValue(choice.Value)

	return &discordgopb.ApplicationCommandOptionChoice{
		Name:              choice.Name,
		NameLocalizations: _encodeLocalizations(choice.NameLocalizations),
		Value:             value,
	}
}

// encodeApplicationCommandOptionChoices encodes multiple discordgo.ApplicationCommandOptionChoice into multiple discordgopb.ApplicationCommandOptionChoice.
func encodeApplicationCommandOptionChoices(choices []*discordgo.ApplicationCommandOptionChoice) []*discordgopb.ApplicationCommandOptionChoice {
	if choices == nil {
		return nil
	}

	result := make([]*discordgopb.ApplicationCommandOptionChoice, len(choices))

	for i, choice := range choices {
		result[i] = encodeApplicationCommandOptionChoice(choice)
	}

	return result
}

// decodeApplicationCommandOptionChoice decodes a discordgopb.ApplicationCommandOptionChoice into a discordgo.ApplicationCommandOptionChoice.
func decodeApplicationCommandOptionChoice(choice *discordgopb.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOptionChoice {
	if choice == nil {
		return nil
	}

	return &discordgo.ApplicationCommandOptionChoice{
		Name:              choice.Name,
		NameLocalizations: _decodeLocalizations(choice.NameLocalizations),
		Value:             choice.Value.AsInterface(),
	}
}

// decodeApplicationCommandOptionChoices decodes multiple discordgopb.ApplicationCommandOptionChoice into multiple discordgo.ApplicationCommandOptionChoice.
func decodeApplicationCommandOptionChoices(choices []*discordgopb.ApplicationCommandOptionChoice) []*discordgo.ApplicationCommandOptionChoice {
	if choices == nil {
		return nil
	}

	result := make([]*discordgo.ApplicationCommandOptionChoice, len(choices))

	for i, choice := range choices {
		result[i] = decodeApplicationCommandOptionChoice(choice)
	}

	return result
}
