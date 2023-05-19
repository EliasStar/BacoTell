package interaction

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/EliasStar/BacoTell/internal/proto/discordpb"
	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type commandServer struct {
	providerpb.UnimplementedCommandServer

	impl provider.Command
}

var _ providerpb.CommandServer = commandServer{}

// CommandData implements providerpb.CommandServer
func (s commandServer) CommandData(context.Context, *emptypb.Empty) (*providerpb.CommandDataResponse, error) {
	data, err := s.impl.CommandData()
	if err != nil {
		return nil, err
	}

	return &providerpb.CommandDataResponse{
		Data: &discordpb.CommandData{
			Type: uint32(data.Type),

			Name:              data.Name,
			NameLocalizations: encodeLocalizations(*data.NameLocalizations),

			Description:              data.Description,
			DescriptionLocalizations: encodeLocalizations(*data.DescriptionLocalizations),

			Options: encodeOptions(data.Options),

			DefaultMemberPermissions: *data.DefaultMemberPermissions,

			Nsfw: *data.NSFW,
		},
	}, nil
}

// Execute implements providerpb.CommandServer
func (s commandServer) Execute(context.Context, *providerpb.ExecuteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Execute(nil)
}

type commandClient struct {
	client providerpb.CommandClient
}

var _ provider.Command = commandClient{}

// CommandData implements provider.Command
func (c commandClient) CommandData() (discordgo.ApplicationCommand, error) {
	res, err := c.client.CommandData(context.Background(), &emptypb.Empty{})
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}

	return discordgo.ApplicationCommand{
		Type: discordgo.ApplicationCommandType(res.Data.Type),

		Name:              res.Data.Name,
		NameLocalizations: decodeLocalizations(res.Data.NameLocalizations),

		Description:              res.Data.Description,
		DescriptionLocalizations: decodeLocalizations(res.Data.DescriptionLocalizations),

		Options: decodeOptions(res.Data.Options),

		DefaultMemberPermissions: &res.Data.DefaultMemberPermissions,

		NSFW: &res.Data.Nsfw,
	}, nil
}

// Execute implements provider.Command
func (c commandClient) Execute(provider.InteractionProxy) error {
	_, err := c.client.Execute(context.Background(), &providerpb.ExecuteRequest{})
	return err
}

func encodeLocalizations(localizations map[discordgo.Locale]string) map[string]string {
	result := make(map[string]string)

	for locale, str := range localizations {
		result[string(locale)] = str
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

func encodeOptions(options []*discordgo.ApplicationCommandOption) []*discordpb.CommandOptionData {
	result := make([]*discordpb.CommandOptionData, len(options))

	for i, option := range options {
		result[i] = &discordpb.CommandOptionData{
			Type: uint32(option.Type),

			Name:              option.Name,
			NameLocalizations: encodeLocalizations(option.NameLocalizations),

			Description:              option.Description,
			DescriptionLocalizations: encodeLocalizations(option.DescriptionLocalizations),

			Required: option.Required,

			Choices: encodeChoices(option.Choices),
			Options: encodeOptions(option.Options),

			ChannelTypes: encodeChannelTypes(option.ChannelTypes),

			MinValue: *option.MinValue,
			MaxValue: option.MaxValue,

			MinLength: uint32(*option.MinLength),
			MaxLength: uint32(option.MaxLength),

			Autocomplete: option.Autocomplete,
		}
	}

	return result
}

func decodeOptions(options []*discordpb.CommandOptionData) []*discordgo.ApplicationCommandOption {
	result := make([]*discordgo.ApplicationCommandOption, len(options))

	for i, option := range options {
		minLength := int(option.MinLength)

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
			MaxLength: int(option.MaxLength),

			Autocomplete: option.Autocomplete,
		}
	}

	return result
}

func encodeChoices(choices []*discordgo.ApplicationCommandOptionChoice) []*discordpb.CommandOptionChoiceData {
	result := make([]*discordpb.CommandOptionChoiceData, len(choices))

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	for i, choice := range choices {
		buffer.Reset()
		enc.Encode(choice.Value)

		result[i] = &discordpb.CommandOptionChoiceData{
			Name:              choice.Name,
			NameLocalizations: encodeLocalizations(choice.NameLocalizations),
			Value:             buffer.Bytes(),
		}
	}

	return result
}

func decodeChoices(choices []*discordpb.CommandOptionChoiceData) []*discordgo.ApplicationCommandOptionChoice {
	result := make([]*discordgo.ApplicationCommandOptionChoice, len(choices))

	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)

	for i, choice := range choices {
		buffer.Reset()
		buffer.Write(choice.Value)

		var value any
		dec.Decode(value)

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
		result[i] = uint32(typ)
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
