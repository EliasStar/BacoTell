package codec

import (
	"bytes"
	"encoding/gob"

	"github.com/EliasStar/BacoTell/internal/proto/discordpb"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cast"
)

func encodeApplicationCommand(command discordgo.ApplicationCommand) *discordpb.ApplicationCommand {
	return &discordpb.ApplicationCommand{
		Type: cast.ToUint32(command.Type),

		Name:              command.Name,
		NameLocalizations: encodeLocalizations(command.NameLocalizations),

		Description:              command.Description,
		DescriptionLocalizations: encodeLocalizations(command.DescriptionLocalizations),

		Options: encodeOptions(command.Options),

		DefaultMemberPermissions: cast.ToInt64(command.DefaultMemberPermissions),

		Nsfw: cast.ToBool(command.NSFW),
	}
}

func decodeApplicationCommand(command *discordpb.ApplicationCommand) discordgo.ApplicationCommand {
	return discordgo.ApplicationCommand{
		Type: discordgo.ApplicationCommandType(command.Type),

		Name:              command.Name,
		NameLocalizations: decodeLocalizations(command.NameLocalizations),

		Description:              command.Description,
		DescriptionLocalizations: decodeLocalizations(command.DescriptionLocalizations),

		Options: decodeOptions(command.Options),

		DefaultMemberPermissions: &command.DefaultMemberPermissions,

		NSFW: &command.Nsfw,
	}
}

func encodeLocalizations(localizations *map[discordgo.Locale]string) map[string]string {
	result := make(map[string]string)

	if localizations == nil {
		return result
	}

	for locale, str := range *localizations {
		result[cast.ToString(locale)] = str
	}

	return result
}

func decodeLocalizations(localizations map[string]string) *map[discordgo.Locale]string {
	result := make(map[discordgo.Locale]string)

	for locale, str := range localizations {
		result[discordgo.Locale(locale)] = str
	}

	return &result
}

func encodeOptions(options []*discordgo.ApplicationCommandOption) []*discordpb.ApplicationCommandOption {
	result := make([]*discordpb.ApplicationCommandOption, len(options))

	for i, option := range options {
		result[i] = &discordpb.ApplicationCommandOption{
			Type: cast.ToUint32(option.Type),

			Name:              option.Name,
			NameLocalizations: encodeLocalizations(&option.NameLocalizations),

			Description:              option.Description,
			DescriptionLocalizations: encodeLocalizations(&option.DescriptionLocalizations),

			Required: option.Required,

			Choices: encodeChoices(option.Choices),
			Options: encodeOptions(option.Options),

			ChannelTypes: encodeChannelTypes(option.ChannelTypes),

			MinValue: cast.ToFloat64(option.MinValue),
			MaxValue: cast.ToFloat64(option.MaxValue),

			MinLength: cast.ToUint32(option.MinLength),
			MaxLength: cast.ToUint32(option.MaxLength),

			Autocomplete: option.Autocomplete,
		}
	}

	return result
}

func decodeOptions(options []*discordpb.ApplicationCommandOption) []*discordgo.ApplicationCommandOption {
	result := make([]*discordgo.ApplicationCommandOption, len(options))

	for i, option := range options {
		minLength := cast.ToInt(option.MinLength)

		result[i] = &discordgo.ApplicationCommandOption{
			Type: discordgo.ApplicationCommandOptionType(option.Type),

			Name:              option.Name,
			NameLocalizations: *decodeLocalizations(option.NameLocalizations),

			Description:              option.Description,
			DescriptionLocalizations: *decodeLocalizations(option.DescriptionLocalizations),

			Required: option.Required,

			Choices: decodeChoices(option.Choices),
			Options: decodeOptions(option.Options),

			ChannelTypes: decodeChannelTypes(option.ChannelTypes),

			MinValue: &option.MinValue,
			MaxValue: option.MaxValue,

			MinLength: &minLength,
			MaxLength: cast.ToInt(option.MaxLength),

			Autocomplete: option.Autocomplete,
		}
	}

	return result
}

func encodeChoices(choices []*discordgo.ApplicationCommandOptionChoice) []*discordpb.ApplicationCommandOptionChoice {
	result := make([]*discordpb.ApplicationCommandOptionChoice, len(choices))

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	for i, choice := range choices {
		buffer.Reset()
		enc.Encode(&choice.Value)

		result[i] = &discordpb.ApplicationCommandOptionChoice{
			Name:              choice.Name,
			NameLocalizations: encodeLocalizations(&choice.NameLocalizations),
			Value:             buffer.Bytes(), // TODO fix ref
		}
	}

	return result
}

func decodeChoices(choices []*discordpb.ApplicationCommandOptionChoice) []*discordgo.ApplicationCommandOptionChoice {
	result := make([]*discordgo.ApplicationCommandOptionChoice, len(choices))

	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)

	for i, choice := range choices {
		buffer.Reset()
		buffer.Write(choice.Value)

		var value any
		dec.Decode(&value)

		result[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:              choice.Name,
			NameLocalizations: *decodeLocalizations(choice.NameLocalizations),
			Value:             value,
		}
	}

	return result
}

func encodeChannelTypes(types []discordgo.ChannelType) []uint32 {
	result := make([]uint32, len(types))

	for i, typ := range types {
		result[i] = cast.ToUint32(typ)
	}

	return result
}

func decodeChannelTypes(types []uint32) []discordgo.ChannelType {
	result := make([]discordgo.ChannelType, len(types))

	for i, typ := range types {
		result[i] = discordgo.ChannelType(typ)
	}

	return result
}
