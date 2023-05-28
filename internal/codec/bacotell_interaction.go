package codec

import (
	"bytes"
	"context"
	"io"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	"github.com/EliasStar/BacoTell/internal/proto/discordpb"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cast"
)

type interactionProxyServer struct {
	bacotellpb.UnimplementedInteractionProxyServer

	impl bacotell.InteractionProxy
}

var _ bacotellpb.InteractionProxyServer = interactionProxyServer{}

// Defer implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Defer(_ context.Context, req *bacotellpb.DeferRequest) (*bacotellpb.DeferResponse, error) {
	return &bacotellpb.DeferResponse{}, s.impl.Defer(req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Respond implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Respond(_ context.Context, req *bacotellpb.RespondRequest) (*bacotellpb.RespondResponse, error) {
	return &bacotellpb.RespondResponse{}, s.impl.Respond(decodeResponse(req.Message), req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Followup implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Followup(_ context.Context, req *bacotellpb.FollowupRequest) (*bacotellpb.FollowupResponse, error) {
	id, err := s.impl.Followup(decodeResponse(req.Message), req.Ephemeral, req.SuppressEmbeds, req.Tts)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.FollowupResponse{Id: id}, nil
}

// Edit implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Edit(_ context.Context, req *bacotellpb.EditRequest) (*bacotellpb.EditResponse, error) {
	return &bacotellpb.EditResponse{}, s.impl.Edit(req.Id, decodeResponse(req.Message))
}

// Delete implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Delete(_ context.Context, req *bacotellpb.DeleteRequest) (*bacotellpb.DeleteResponse, error) {
	return &bacotellpb.DeleteResponse{}, s.impl.Delete(req.Id)
}

type interactionProxyClient struct {
	client bacotellpb.InteractionProxyClient
}

var _ bacotell.InteractionProxy = interactionProxyClient{}

// Defer implements bacotell.InteractionProxy
func (c interactionProxyClient) Defer(ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Defer(context.Background(), &bacotellpb.DeferRequest{
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Respond implements bacotell.InteractionProxy
func (c interactionProxyClient) Respond(message bacotell.Response, ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Respond(context.Background(), &bacotellpb.RespondRequest{
		Message:        encodeResponse(message),
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Followup implements bacotell.InteractionProxy
func (c interactionProxyClient) Followup(message bacotell.Response, ephemeral bool, suppressEmbeds bool, tts bool) (string, error) {
	res, err := c.client.Followup(context.Background(), &bacotellpb.FollowupRequest{
		Message:        encodeResponse(message),
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// Edit implements bacotell.InteractionProxy
func (c interactionProxyClient) Edit(id string, message bacotell.Response) error {
	_, err := c.client.Edit(context.Background(), &bacotellpb.EditRequest{
		Id:      id,
		Message: encodeResponse(message),
	})

	return err
}

// Delete implements bacotell.InteractionProxy
func (c interactionProxyClient) Delete(id string) error {
	_, err := c.client.Delete(context.Background(), &bacotellpb.DeleteRequest{Id: id})

	return err
}

// TODO

func encodeResponse(response bacotell.Response) *bacotellpb.Response {
	return &bacotellpb.Response{
		Content: response.Content,
		AllowedMentions: &discordpb.MessageAllowedMentions{
			Parse:       encodeParse(response.AllowedMentions.Parse),
			Roles:       response.AllowedMentions.Roles,
			Users:       response.AllowedMentions.Users,
			RepliedUser: response.AllowedMentions.RepliedUser,
		},
		Components: encodeComponents(response.Components),
		Embeds:     encodeEmbeds(response.Embeds),
		Files:      encodeFiles(response.Files),
	}
}

func encodeParse(allowedMentions []discordgo.AllowedMentionType) []string {
	result := make([]string, len(allowedMentions))

	for i, allowedMention := range allowedMentions {
		result[i] = cast.ToString(allowedMention)
	}

	return result
}

func encodeComponents(components []discordgo.MessageComponent) []*discordpb.MessageComponent {
	result := make([]*discordpb.MessageComponent, len(components))

	// for i, component := range components {
	// 	result[i] =
	// }

	return result
}

func encodeEmbeds(embeds []*discordgo.MessageEmbed) []*discordpb.MessageEmbed {
	result := make([]*discordpb.MessageEmbed, len(embeds))

	for i, embed := range embeds {
		result[i] = &discordpb.MessageEmbed{
			Url:         embed.URL,
			Type:        cast.ToString(embed.Type),
			Title:       embed.Title,
			Description: embed.Description,
			Timestamp:   embed.Timestamp,
			Color:       cast.ToUint32(embed.Color),
			Footer: &discordpb.MessageEmbedFooter{
				Text:         embed.Footer.Text,
				IconUrl:      embed.Footer.IconURL,
				ProxyIconUrl: embed.Footer.ProxyIconURL,
			},
			Image: &discordpb.MessageEmbedImage{
				Url:      embed.Image.URL,
				ProxyUrl: embed.Image.ProxyURL,
				Width:    cast.ToUint32(embed.Image.Width),
				Height:   cast.ToUint32(embed.Image.Height),
			},
			Thumbnail: &discordpb.MessageEmbedThumbnail{
				Url:      embed.Thumbnail.URL,
				ProxyUrl: embed.Thumbnail.ProxyURL,
				Width:    cast.ToUint32(embed.Thumbnail.Width),
				Height:   cast.ToUint32(embed.Thumbnail.Height),
			},
			Video: &discordpb.MessageEmbedVideo{
				Url:    embed.Video.URL,
				Width:  cast.ToUint32(embed.Video.Width),
				Height: cast.ToUint32(embed.Video.Height),
			},
			Provider: &discordpb.MessageEmbedProvider{
				Url:  embed.Provider.URL,
				Name: embed.Provider.Name,
			},
			Author: &discordpb.MessageEmbedAuthor{
				Url:          embed.Author.URL,
				Name:         embed.Author.Name,
				IconUrl:      embed.Author.IconURL,
				ProxyIconUrl: embed.Author.ProxyIconURL,
			},
			Fields: encodeFields(embed.Fields),
		}
	}

	return result
}

func encodeFields(fields []*discordgo.MessageEmbedField) []*discordpb.MessageEmbedField {
	result := make([]*discordpb.MessageEmbedField, len(fields))

	for i, field := range fields {
		result[i] = &discordpb.MessageEmbedField{
			Name:   field.Name,
			Value:  field.Value,
			Inline: field.Inline,
		}
	}

	return result
}

func encodeFiles(files []*discordgo.File) []*discordpb.File {
	result := make([]*discordpb.File, len(files))

	for i, file := range files {
		content, _ := io.ReadAll(file.Reader)

		result[i] = &discordpb.File{
			Name:        file.Name,
			ContentType: file.ContentType,
			Content:     content,
		}
	}

	return result
}

func decodeResponse(response *bacotellpb.Response) bacotell.Response {
	return bacotell.Response{
		Content: response.Content,
		AllowedMentions: discordgo.MessageAllowedMentions{
			Parse:       decodeParse(response.AllowedMentions.Parse),
			Roles:       response.AllowedMentions.Roles,
			Users:       response.AllowedMentions.Users,
			RepliedUser: response.AllowedMentions.RepliedUser,
		},
		Components: decodeComponents(response.Components),
		Embeds:     decodeEmbeds(response.Embeds),
		Files:      decodeFiles(response.Files),
	}
}

func decodeParse(allowedMentions []string) []discordgo.AllowedMentionType {
	result := make([]discordgo.AllowedMentionType, len(allowedMentions))

	for i, allowedMention := range allowedMentions {
		result[i] = discordgo.AllowedMentionType(allowedMention)
	}

	return result
}

func decodeComponents(components []*discordpb.MessageComponent) []discordgo.MessageComponent {
	result := make([]discordgo.MessageComponent, len(components))

	// for i, component := range components {
	// 	result[i] =
	// }

	return result
}

func decodeEmbeds(embeds []*discordpb.MessageEmbed) []*discordgo.MessageEmbed {
	result := make([]*discordgo.MessageEmbed, len(embeds))

	for i, embed := range embeds {
		result[i] = &discordgo.MessageEmbed{
			URL:         embed.Url,
			Type:        discordgo.EmbedType(embed.Type),
			Title:       embed.Title,
			Description: embed.Description,
			Timestamp:   embed.Timestamp,
			Color:       cast.ToInt(embed.Color),
			Footer: &discordgo.MessageEmbedFooter{
				Text:         embed.Footer.Text,
				IconURL:      embed.Footer.IconUrl,
				ProxyIconURL: embed.Footer.ProxyIconUrl,
			},
			Image: &discordgo.MessageEmbedImage{
				URL:      embed.Image.Url,
				ProxyURL: embed.Image.ProxyUrl,
				Width:    cast.ToInt(embed.Image.Width),
				Height:   cast.ToInt(embed.Image.Height),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:      embed.Thumbnail.Url,
				ProxyURL: embed.Thumbnail.ProxyUrl,
				Width:    cast.ToInt(embed.Thumbnail.Width),
				Height:   cast.ToInt(embed.Thumbnail.Height),
			},
			Video: &discordgo.MessageEmbedVideo{
				URL:    embed.Video.Url,
				Width:  cast.ToInt(embed.Video.Width),
				Height: cast.ToInt(embed.Video.Height),
			},
			Provider: &discordgo.MessageEmbedProvider{
				URL:  embed.Provider.Url,
				Name: embed.Provider.Name,
			},
			Author: &discordgo.MessageEmbedAuthor{
				URL:          embed.Author.Url,
				Name:         embed.Author.Name,
				IconURL:      embed.Author.IconUrl,
				ProxyIconURL: embed.Author.ProxyIconUrl,
			},
			Fields: decodeFields(embed.Fields),
		}
	}

	return result
}

func decodeFields(fields []*discordpb.MessageEmbedField) []*discordgo.MessageEmbedField {
	result := make([]*discordgo.MessageEmbedField, len(fields))

	for i, field := range fields {
		result[i] = &discordgo.MessageEmbedField{
			Name:   field.Name,
			Value:  field.Value,
			Inline: field.Inline,
		}
	}

	return result
}

func decodeFiles(files []*discordpb.File) []*discordgo.File {
	result := make([]*discordgo.File, len(files))

	for i, file := range files {
		result[i] = &discordgo.File{
			Name:        file.Name,
			ContentType: file.ContentType,
			Reader:      bytes.NewReader(file.Content),
		}
	}

	return result
}
