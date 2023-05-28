package codec

import (
	"time"

	"github.com/EliasStar/BacoTell/internal/proto/discordpb"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cast"
)

func encodeUser(user *discordgo.User) *discordpb.User {
	return &discordpb.User{
		Id:            user.ID,
		Username:      user.Username,
		Discriminator: user.Discriminator,
		Email:         user.Email,

		Avatar: user.Avatar,
		Banner: user.Banner,

		AccentColor: cast.ToUint32(user.AccentColor),
		Locale:      user.Locale,

		Bot:        user.Bot,
		System:     user.System,
		MfaEnabled: user.MFAEnabled,
		Verified:   user.Verified,

		PremiumType: cast.ToUint32(user.PremiumType),
		Flags:       cast.ToUint64(user.Flags),
		PublicFlags: cast.ToUint64(user.PublicFlags),
	}
}

func encodeRole(role *discordgo.Role) *discordpb.Role {
	return &discordpb.Role{
		Id:   role.ID,
		Name: role.Name,

		Color: cast.ToUint32(role.Color),

		Managed:     role.Managed,
		Mentionable: role.Mentionable,
		Hoist:       role.Hoist,

		Position:    cast.ToUint32(role.Position),
		Permissions: role.Permissions,
	}
}

func encodeChannel(channel *discordgo.Channel) *discordpb.Channel {
	lastPinTimestamp := ""
	if channel.LastPinTimestamp != nil {
		timestamp, err := channel.LastPinTimestamp.MarshalText()
		if err != nil {
			lastPinTimestamp = cast.ToString(timestamp)
		}
	}

	archiveTimestamp, _ := channel.ThreadMetadata.ArchiveTimestamp.MarshalText()

	var defaultSortOrder *uint32
	if channel.DefaultSortOrder != nil {
		order := cast.ToUint32(*channel.DefaultSortOrder)
		defaultSortOrder = &order
	}

	return &discordpb.Channel{
		Id:               channel.ID,
		GuildId:          channel.GuildID,
		Name:             channel.Name,
		Topic:            channel.Topic,
		Type:             cast.ToUint32(channel.Type),
		LastMessageId:    channel.LastMessageID,
		LastPinTimestamp: lastPinTimestamp,
		MessageCount:     cast.ToUint32(channel.MessageCount),
		MemberCount:      cast.ToUint32(channel.MemberCount),
		Nsfw:             channel.NSFW,
		Icon:             channel.Icon,
		Position:         cast.ToUint32(channel.Position),
		Bitrate:          cast.ToUint32(channel.Bitrate),

		// TODO
		// Recipients: 			 []*discordpb.User{},
		// Messages:             []*discordpb.Message{},
		// PermissionOverwrites: []*discordpb.Overwrite{},

		UserLimit:        cast.ToUint32(channel.UserLimit),
		ParentId:         channel.ParentID,
		RateLimitPerUser: cast.ToUint32(channel.RateLimitPerUser),
		OwnerId:          channel.OwnerID,
		ApplicationId:    channel.ApplicationID,
		ThreadMetadata: &discordpb.ThreadMetadata{
			Archived:            channel.ThreadMetadata.Archived,
			AutoArchiveDuration: cast.ToUint32(channel.ThreadMetadata.AutoArchiveDuration),
			ArchiveTimestamp:    cast.ToString(archiveTimestamp),
			Locked:              channel.ThreadMetadata.Locked,
			Invitable:           channel.ThreadMetadata.Invitable,
		},

		// TODO
		// Member:        &discordgo.ThreadMember{},
		// Members:       []*discordgo.ThreadMember{},
		// Flags:         discordgo.ChannelFlags(channel.Flags),
		// AvailableTags: []discordgo.ForumTag{},

		AppliedTags: channel.AppliedTags,
		DefaultReactionEmoji: &discordpb.DefaultReaction{
			EmojiId:   channel.DefaultReactionEmoji.EmojiID,
			EmojiName: channel.DefaultReactionEmoji.EmojiName,
		},
		DefaultThreadRateLimitPerUser: cast.ToUint32(channel.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            cast.ToUint32(channel.DefaultForumLayout),
	}
}

func encodeAttachment(attachment *discordgo.MessageAttachment) *discordpb.Attachment {
	return &discordpb.Attachment{
		Id:       attachment.ID,
		Filename: attachment.Filename,

		Url:         attachment.URL,
		ProxyUrl:    attachment.ProxyURL,
		ContentType: attachment.ContentType,
		Size:        cast.ToUint32(attachment.Size),

		Height: cast.ToUint32(attachment.Height),
		Width:  cast.ToUint32(attachment.Width),

		Ephemeral: attachment.Ephemeral,
	}
}

func decodeUser(user *discordpb.User) *discordgo.User {
	return &discordgo.User{
		ID:            user.Id,
		Username:      user.Username,
		Discriminator: user.Discriminator,
		Email:         user.Email,

		Avatar: user.Avatar,
		Banner: user.Banner,

		AccentColor: cast.ToInt(user.AccentColor),
		Locale:      user.Locale,

		Bot:        user.Bot,
		System:     user.System,
		MFAEnabled: user.MfaEnabled,
		Verified:   user.Verified,

		PremiumType: cast.ToInt(user.PremiumType),
		Flags:       cast.ToInt(user.Flags),
		PublicFlags: discordgo.UserFlags(user.PublicFlags),
	}
}

func decodeRole(role *discordpb.Role) *discordgo.Role {
	return &discordgo.Role{
		ID:   role.Id,
		Name: role.Name,

		Color: cast.ToInt(role.Color),

		Managed:     role.Managed,
		Mentionable: role.Mentionable,
		Hoist:       role.Hoist,

		Position:    cast.ToInt(role.Position),
		Permissions: role.Permissions,
	}
}

func decodeChannel(channel *discordpb.Channel) *discordgo.Channel {
	var lastPinTimestamp *time.Time
	if channel.LastPinTimestamp != "" {
		lastPinTimestamp.UnmarshalText([]byte(channel.LastPinTimestamp))
	}

	var archiveTimestamp time.Time
	archiveTimestamp.UnmarshalText([]byte(channel.ThreadMetadata.ArchiveTimestamp))

	var defaultSortOrder *discordgo.ForumSortOrderType
	if channel.DefaultSortOrder != nil {
		order := discordgo.ForumSortOrderType(*channel.DefaultSortOrder)
		defaultSortOrder = &order
	}

	return &discordgo.Channel{
		ID:               channel.Id,
		GuildID:          channel.GuildId,
		Name:             channel.Name,
		Topic:            channel.Topic,
		Type:             discordgo.ChannelType(channel.Type),
		LastMessageID:    channel.LastMessageId,
		LastPinTimestamp: lastPinTimestamp,
		MessageCount:     cast.ToInt(channel.MessageCount),
		MemberCount:      cast.ToInt(channel.MemberCount),
		NSFW:             channel.Nsfw,
		Icon:             channel.Icon,
		Position:         cast.ToInt(channel.Position),
		Bitrate:          cast.ToInt(channel.Bitrate),

		// TODO
		// Recipients:           []*discordgo.User{},
		// Messages:             []*discordgo.Message{},
		// PermissionOverwrites: []*discordgo.PermissionOverwrite{},

		UserLimit:        cast.ToInt(channel.UserLimit),
		ParentID:         channel.ParentId,
		RateLimitPerUser: cast.ToInt(channel.RateLimitPerUser),
		OwnerID:          channel.OwnerId,
		ApplicationID:    channel.ApplicationId,
		ThreadMetadata: &discordgo.ThreadMetadata{
			Archived:            channel.ThreadMetadata.Archived,
			AutoArchiveDuration: cast.ToInt(channel.ThreadMetadata.AutoArchiveDuration),
			ArchiveTimestamp:    archiveTimestamp,
			Locked:              channel.ThreadMetadata.Locked,
			Invitable:           channel.ThreadMetadata.Invitable,
		},

		// TODO
		// Member:        &discordgo.ThreadMember{},
		// Members:       []*discordgo.ThreadMember{},
		// Flags:         discordgo.ChannelFlags(channel.Flags),
		// AvailableTags: []discordgo.ForumTag{},

		AppliedTags: channel.AppliedTags,
		DefaultReactionEmoji: discordgo.ForumDefaultReaction{
			EmojiID:   channel.DefaultReactionEmoji.EmojiId,
			EmojiName: channel.DefaultReactionEmoji.EmojiName,
		},
		DefaultThreadRateLimitPerUser: cast.ToInt(channel.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            discordgo.ForumLayout(channel.DefaultForumLayout),
	}
}

func decodeAttachment(attachment *discordpb.Attachment) *discordgo.MessageAttachment {
	return &discordgo.MessageAttachment{
		ID:       attachment.Id,
		Filename: attachment.Filename,

		URL:         attachment.Url,
		ProxyURL:    attachment.ProxyUrl,
		ContentType: attachment.ContentType,
		Size:        cast.ToInt(attachment.Size),

		Height: cast.ToInt(attachment.Height),
		Width:  cast.ToInt(attachment.Width),

		Ephemeral: attachment.Ephemeral,
	}
}
