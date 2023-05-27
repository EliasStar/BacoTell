package interaction

import (
	"bytes"
	context "context"
	"io"
	"time"

	"github.com/EliasStar/BacoTell/internal/proto/discordpb"
	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cast"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type interactionProxyServer struct {
	providerpb.UnimplementedInteractionProxyServer

	impl provider.InteractionProxy
}

var _ providerpb.InteractionProxyServer = interactionProxyServer{}

// Defer implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Defer(_ context.Context, req *providerpb.DeferRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Defer(req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Respond implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Respond(_ context.Context, req *providerpb.RespondRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Respond(decodeResponse(req.Message), req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Followup implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Followup(_ context.Context, req *providerpb.FollowupRequest) (*providerpb.FollowupResponse, error) {
	id, err := s.impl.Followup(decodeResponse(req.Message), req.Ephemeral, req.SuppressEmbeds, req.Tts)
	if err != nil {
		return nil, err
	}

	return &providerpb.FollowupResponse{Id: id}, nil
}

// Edit implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Edit(_ context.Context, req *providerpb.EditRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Edit(req.Id, decodeResponse(req.Message))
}

// Delete implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Delete(_ context.Context, req *providerpb.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Delete(req.Id)
}

type interactionProxyClient struct {
	client providerpb.InteractionProxyClient
}

var _ provider.InteractionProxy = interactionProxyClient{}

// Defer implements provider.InteractionProxy
func (c interactionProxyClient) Defer(ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Defer(context.Background(), &providerpb.DeferRequest{
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Respond implements provider.InteractionProxy
func (c interactionProxyClient) Respond(message provider.Response, ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Respond(context.Background(), &providerpb.RespondRequest{
		Message:        encodeResponse(message),
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Followup implements provider.InteractionProxy
func (c interactionProxyClient) Followup(message provider.Response, ephemeral bool, suppressEmbeds bool, tts bool) (string, error) {
	res, err := c.client.Followup(context.Background(), &providerpb.FollowupRequest{
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

// Edit implements provider.InteractionProxy
func (c interactionProxyClient) Edit(id string, message provider.Response) error {
	_, err := c.client.Edit(context.Background(), &providerpb.EditRequest{
		Id:      id,
		Message: encodeResponse(message),
	})

	return err
}

// Delete implements provider.InteractionProxy
func (c interactionProxyClient) Delete(id string) error {
	_, err := c.client.Delete(context.Background(), &providerpb.DeleteRequest{
		Id: id,
	})

	return err
}

type executeProxyServer struct {
	providerpb.UnimplementedExecuteProxyServer
	interactionProxyServer

	impl provider.ExecuteProxy
}

var _ providerpb.InteractionProxyServer = executeProxyServer{}
var _ providerpb.ExecuteProxyServer = executeProxyServer{}

// StringOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) StringOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.StringOptionResponse, error) {
	val, err := s.impl.StringOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.StringOptionResponse{
		Value: val,
	}, nil
}

// IntegerOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) IntegerOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.IntegerOptionResponse, error) {
	val, err := s.impl.IntegerOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.IntegerOptionResponse{
		Value: val,
	}, nil
}

// NumberOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) NumberOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.NumberOptionResponse, error) {
	val, err := s.impl.NumberOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.NumberOptionResponse{
		Value: val,
	}, nil
}

// BooleanOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) BooleanOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.BooleanOptionResponse, error) {
	val, err := s.impl.BooleanOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.BooleanOptionResponse{
		Value: val,
	}, nil
}

// UserOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) UserOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.UserOptionResponse, error) {
	val, err := s.impl.UserOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.UserOptionResponse{
		Value: &discordpb.User{
			Id:            val.ID,
			Username:      val.Username,
			Discriminator: val.Discriminator,
			Email:         val.Email,

			Avatar: val.Avatar,
			Banner: val.Banner,

			AccentColor: cast.ToUint32(val.AccentColor),
			Locale:      val.Locale,

			Bot:        val.Bot,
			System:     val.System,
			MfaEnabled: val.MFAEnabled,
			Verified:   val.Verified,

			PremiumType: cast.ToUint32(val.PremiumType),
			Flags:       cast.ToUint64(val.Flags),
			PublicFlags: cast.ToUint64(val.PublicFlags),
		},
	}, nil
}

// RoleOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) RoleOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.RoleOptionResponse, error) {
	val, err := s.impl.RoleOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.RoleOptionResponse{
		Value: &discordpb.Role{
			Id:   val.ID,
			Name: val.Name,

			Color: cast.ToUint32(val.Color),

			Managed:     val.Managed,
			Mentionable: val.Mentionable,
			Hoist:       val.Hoist,

			Position:    cast.ToUint32(val.Position),
			Permissions: val.Permissions,
		},
	}, nil
}

// ChannelOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) ChannelOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.ChannelOptionResponse, error) {
	val, err := s.impl.ChannelOption(req.Name)
	if err != nil {
		return nil, err
	}

	lastPinTimestamp := ""
	if val.LastPinTimestamp != nil {
		timestamp, err := val.LastPinTimestamp.MarshalText()
		if err != nil {
			lastPinTimestamp = cast.ToString(timestamp)
		}
	}

	archiveTimestamp, _ := val.ThreadMetadata.ArchiveTimestamp.MarshalText()

	var defaultSortOrder *uint32
	if val.DefaultSortOrder != nil {
		order := cast.ToUint32(*val.DefaultSortOrder)
		defaultSortOrder = &order
	}

	return &providerpb.ChannelOptionResponse{
		Value: &discordpb.Channel{
			Id:               val.ID,
			GuildId:          val.GuildID,
			Name:             val.Name,
			Topic:            val.Topic,
			Type:             cast.ToUint32(val.Type),
			LastMessageId:    val.LastMessageID,
			LastPinTimestamp: lastPinTimestamp,
			MessageCount:     cast.ToUint32(val.MessageCount),
			MemberCount:      cast.ToUint32(val.MemberCount),
			Nsfw:             val.NSFW,
			Icon:             val.Icon,
			Position:         cast.ToUint32(val.Position),
			Bitrate:          cast.ToUint32(val.Bitrate),

			// TODO
			// Recipients: 			 []*discordpb.User{},
			// Messages:             []*discordpb.Message{},
			// PermissionOverwrites: []*discordpb.Overwrite{},

			UserLimit:        cast.ToUint32(val.UserLimit),
			ParentId:         val.ParentID,
			RateLimitPerUser: cast.ToUint32(val.RateLimitPerUser),
			OwnerId:          val.OwnerID,
			ApplicationId:    val.ApplicationID,
			ThreadMetadata: &discordpb.ThreadMetadata{
				Archived:            val.ThreadMetadata.Archived,
				AutoArchiveDuration: cast.ToUint32(val.ThreadMetadata.AutoArchiveDuration),
				ArchiveTimestamp:    cast.ToString(archiveTimestamp),
				Locked:              val.ThreadMetadata.Locked,
				Invitable:           val.ThreadMetadata.Invitable,
			},

			// TODO
			// Member:        &discordgo.ThreadMember{},
			// Members:       []*discordgo.ThreadMember{},
			// Flags:         discordgo.ChannelFlags(val.Flags),
			// AvailableTags: []discordgo.ForumTag{},

			AppliedTags: val.AppliedTags,
			DefaultReactionEmoji: &discordpb.DefaultReaction{
				EmojiId:   val.DefaultReactionEmoji.EmojiID,
				EmojiName: val.DefaultReactionEmoji.EmojiName,
			},
			DefaultThreadRateLimitPerUser: cast.ToUint32(val.DefaultThreadRateLimitPerUser),
			DefaultSortOrder:              defaultSortOrder,
			DefaultForumLayout:            cast.ToUint32(val.DefaultForumLayout),
		},
	}, nil
}

// AttachmentOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) AttachmentOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.AttachmentOptionResponse, error) {
	val, err := s.impl.AttachmentOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.AttachmentOptionResponse{
		Value: &discordpb.Attachment{
			Id:       val.ID,
			Filename: val.Filename,

			Url:         val.URL,
			ProxyUrl:    val.ProxyURL,
			ContentType: val.ContentType,
			Size:        cast.ToUint32(val.Size),

			Height: cast.ToUint32(val.Height),
			Width:  cast.ToUint32(val.Width),

			Ephemeral: val.Ephemeral,
		},
	}, nil
}

type executeProxyClient struct {
	interactionProxyClient

	client providerpb.ExecuteProxyClient
}

var _ provider.InteractionProxy = executeProxyClient{}
var _ provider.ExecuteProxy = executeProxyClient{}

// StringOption implements provider.ExecuteProxy
func (c executeProxyClient) StringOption(name string) (string, error) {
	res, err := c.client.StringOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return res.Value, nil
}

// IntegerOption implements provider.ExecuteProxy
func (c executeProxyClient) IntegerOption(name string) (int64, error) {
	res, err := c.client.IntegerOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// NumberOption implements provider.ExecuteProxy
func (c executeProxyClient) NumberOption(name string) (float64, error) {
	res, err := c.client.NumberOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// BooleanOption implements provider.ExecuteProxy
func (c executeProxyClient) BooleanOption(name string) (bool, error) {
	res, err := c.client.BooleanOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return false, err
	}

	return res.Value, nil
}

// UserOption implements provider.ExecuteProxy
func (c executeProxyClient) UserOption(name string) (*discordgo.User, error) {
	res, err := c.client.UserOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return &discordgo.User{
		ID:            res.Value.Id,
		Username:      res.Value.Username,
		Discriminator: res.Value.Discriminator,
		Email:         res.Value.Email,

		Avatar: res.Value.Avatar,
		Banner: res.Value.Banner,

		AccentColor: cast.ToInt(res.Value.AccentColor),
		Locale:      res.Value.Locale,

		Bot:        res.Value.Bot,
		System:     res.Value.System,
		MFAEnabled: res.Value.MfaEnabled,
		Verified:   res.Value.Verified,

		PremiumType: cast.ToInt(res.Value.PremiumType),
		Flags:       cast.ToInt(res.Value.Flags),
		PublicFlags: discordgo.UserFlags(res.Value.PublicFlags),
	}, nil
}

// RoleOption implements provider.ExecuteProxy
func (c executeProxyClient) RoleOption(name string) (*discordgo.Role, error) {
	res, err := c.client.RoleOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return &discordgo.Role{
		ID:   res.Value.Id,
		Name: res.Value.Name,

		Color: cast.ToInt(res.Value.Color),

		Managed:     res.Value.Managed,
		Mentionable: res.Value.Mentionable,
		Hoist:       res.Value.Hoist,

		Position:    cast.ToInt(res.Value.Position),
		Permissions: res.Value.Permissions,
	}, nil
}

// ChannelOption implements provider.ExecuteProxy
func (c executeProxyClient) ChannelOption(name string) (*discordgo.Channel, error) {
	res, err := c.client.ChannelOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	var lastPinTimestamp *time.Time
	if res.Value.LastPinTimestamp != "" {
		lastPinTimestamp.UnmarshalText([]byte(res.Value.LastPinTimestamp))
	}

	var archiveTimestamp time.Time
	archiveTimestamp.UnmarshalText([]byte(res.Value.ThreadMetadata.ArchiveTimestamp))

	var defaultSortOrder *discordgo.ForumSortOrderType
	if res.Value.DefaultSortOrder != nil {
		order := discordgo.ForumSortOrderType(*res.Value.DefaultSortOrder)
		defaultSortOrder = &order
	}

	return &discordgo.Channel{
		ID:               res.Value.Id,
		GuildID:          res.Value.GuildId,
		Name:             res.Value.Name,
		Topic:            res.Value.Topic,
		Type:             discordgo.ChannelType(res.Value.Type),
		LastMessageID:    res.Value.LastMessageId,
		LastPinTimestamp: lastPinTimestamp,
		MessageCount:     cast.ToInt(res.Value.MessageCount),
		MemberCount:      cast.ToInt(res.Value.MemberCount),
		NSFW:             res.Value.Nsfw,
		Icon:             res.Value.Icon,
		Position:         cast.ToInt(res.Value.Position),
		Bitrate:          cast.ToInt(res.Value.Bitrate),

		// TODO
		// Recipients:           []*discordgo.User{},
		// Messages:             []*discordgo.Message{},
		// PermissionOverwrites: []*discordgo.PermissionOverwrite{},

		UserLimit:        cast.ToInt(res.Value.UserLimit),
		ParentID:         res.Value.ParentId,
		RateLimitPerUser: cast.ToInt(res.Value.RateLimitPerUser),
		OwnerID:          res.Value.OwnerId,
		ApplicationID:    res.Value.ApplicationId,
		ThreadMetadata: &discordgo.ThreadMetadata{
			Archived:            res.Value.ThreadMetadata.Archived,
			AutoArchiveDuration: cast.ToInt(res.Value.ThreadMetadata.AutoArchiveDuration),
			ArchiveTimestamp:    archiveTimestamp,
			Locked:              res.Value.ThreadMetadata.Locked,
			Invitable:           res.Value.ThreadMetadata.Invitable,
		},

		// TODO
		// Member:        &discordgo.ThreadMember{},
		// Members:       []*discordgo.ThreadMember{},
		// Flags:         discordgo.ChannelFlags(res.Value.Flags),
		// AvailableTags: []discordgo.ForumTag{},

		AppliedTags: res.Value.AppliedTags,
		DefaultReactionEmoji: discordgo.ForumDefaultReaction{
			EmojiID:   res.Value.DefaultReactionEmoji.EmojiId,
			EmojiName: res.Value.DefaultReactionEmoji.EmojiName,
		},
		DefaultThreadRateLimitPerUser: cast.ToInt(res.Value.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            discordgo.ForumLayout(res.Value.DefaultForumLayout),
	}, nil
}

// AttachmentOption implements provider.ExecuteProxy
func (c executeProxyClient) AttachmentOption(name string) (*discordgo.MessageAttachment, error) {
	res, err := c.client.AttachmentOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return &discordgo.MessageAttachment{
		ID:       res.Value.Id,
		Filename: res.Value.Filename,

		URL:         res.Value.Url,
		ProxyURL:    res.Value.ProxyUrl,
		ContentType: res.Value.ContentType,
		Size:        cast.ToInt(res.Value.Size),

		Height: cast.ToInt(res.Value.Height),
		Width:  cast.ToInt(res.Value.Width),

		Ephemeral: res.Value.Ephemeral,
	}, nil
}

type handleProxyServer struct {
	providerpb.UnimplementedHandleProxyServer
	interactionProxyServer

	impl provider.HandleProxy
}

var _ providerpb.InteractionProxyServer = handleProxyServer{}
var _ providerpb.HandleProxyServer = handleProxyServer{}

type handleProxyClient struct {
	interactionProxyClient

	client providerpb.HandleProxyClient
}

var _ provider.InteractionProxy = handleProxyClient{}
var _ provider.HandleProxy = handleProxyClient{}

func encodeResponse(response provider.Response) *providerpb.Response {
	return &providerpb.Response{
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

func decodeResponse(response *providerpb.Response) provider.Response {
	return provider.Response{
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
