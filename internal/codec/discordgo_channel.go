package codec

import (
	"time"

	"github.com/EliasStar/BacoTell/internal/proto/discordgopb"
	util "github.com/EliasStar/BacoTell/pkg/bacotell_util"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func encodeChannel(channel *discordgo.Channel) *discordgopb.Channel {
	if channel == nil {
		return nil
	}

	var lastPinTimestamp *timestamppb.Timestamp
	if channel.LastPinTimestamp != nil {
		lastPinTimestamp = timestamppb.New(*channel.LastPinTimestamp)
	}

	return &discordgopb.Channel{
		Id:      channel.ID,
		GuildId: channel.GuildID,
		Name:    channel.Name,
		Topic:   channel.Topic,
		Type:    int32(channel.Type),

		LastMessageId:    channel.LastMessageID,
		LastPinTimestamp: lastPinTimestamp,

		MessageCount: int32(channel.MessageCount),
		MemberCount:  int32(channel.MemberCount),

		Nsfw:     channel.NSFW,
		Icon:     channel.Icon,
		Position: int32(channel.Position),
		Bitrate:  int32(channel.Bitrate),

		Recipients:           encodeUsers(channel.Recipients),
		Messages:             encodeMessages(channel.Messages),
		PermissionOverwrites: encodePermissionOverwrites(channel.PermissionOverwrites),

		UserLimit:        int32(channel.UserLimit),
		ParentId:         channel.ParentID,
		RateLimitPerUser: int32(channel.RateLimitPerUser),
		OwnerId:          channel.OwnerID,
		ApplicationId:    channel.ApplicationID,

		ThreadMetadata: encodeThreadMetadata(channel.ThreadMetadata),
		Member:         encodeThreadMember(channel.Member),
		Members:        encodeThreadMembers(channel.Members),

		Flags: int32(channel.Flags),

		AvailableTags: encodeForumTags(channel.AvailableTags),
		AppliedTags:   channel.AppliedTags,

		DefaultReactionEmoji:          encodeForumDefaultReaction(&channel.DefaultReactionEmoji),
		DefaultThreadRateLimitPerUser: int32(channel.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              util.PtrConv[discordgo.ForumSortOrderType, int32](channel.DefaultSortOrder),
		DefaultForumLayout:            int32(channel.DefaultForumLayout),
	}
}

func encodeChannels(channels []*discordgo.Channel) []*discordgopb.Channel {
	if channels == nil {
		return nil
	}

	result := make([]*discordgopb.Channel, len(channels))

	for i, channel := range channels {
		result[i] = encodeChannel(channel)
	}

	return result
}

func decodeChannel(channel *discordgopb.Channel) *discordgo.Channel {
	if channel == nil {
		return nil
	}

	var lastPinTimestamp *time.Time
	if channel.LastPinTimestamp != nil {
		lastPinTimestamp = util.Ptr(channel.LastPinTimestamp.AsTime())
	}

	return &discordgo.Channel{
		ID:      channel.Id,
		GuildID: channel.GuildId,
		Name:    channel.Name,
		Topic:   channel.Topic,
		Type:    discordgo.ChannelType(channel.Type),

		LastMessageID:    channel.LastMessageId,
		LastPinTimestamp: lastPinTimestamp,

		MessageCount: int(channel.MessageCount),
		MemberCount:  int(channel.MemberCount),

		NSFW:     channel.Nsfw,
		Icon:     channel.Icon,
		Position: int(channel.Position),
		Bitrate:  int(channel.Bitrate),

		Recipients:           decodeUsers(channel.Recipients),
		Messages:             decodeMessages(channel.Messages),
		PermissionOverwrites: decodePermissionOverwrites(channel.PermissionOverwrites),

		UserLimit:        int(channel.UserLimit),
		ParentID:         channel.ParentId,
		RateLimitPerUser: int(channel.RateLimitPerUser),
		OwnerID:          channel.OwnerId,
		ApplicationID:    channel.ApplicationId,

		ThreadMetadata: decodeThreadMetadata(channel.ThreadMetadata),
		Member:         decodeThreadMember(channel.Member),
		Members:        decodeThreadMembers(channel.Members),

		Flags: discordgo.ChannelFlags(channel.Flags),

		AvailableTags: decodeForumTags(channel.AvailableTags),
		AppliedTags:   channel.AppliedTags,

		DefaultReactionEmoji:          *decodeForumDefaultReaction(channel.DefaultReactionEmoji),
		DefaultThreadRateLimitPerUser: int(channel.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              util.PtrConv[int32, discordgo.ForumSortOrderType](channel.DefaultSortOrder),
		DefaultForumLayout:            discordgo.ForumLayout(channel.DefaultForumLayout),
	}
}

func decodeChannels(channels []*discordgopb.Channel) []*discordgo.Channel {
	if channels == nil {
		return nil
	}

	result := make([]*discordgo.Channel, len(channels))

	for i, channel := range channels {
		result[i] = decodeChannel(channel)
	}

	return result
}

func encodeMessage(message *discordgo.Message) *discordgopb.Message {
	if message == nil {
		return nil
	}

	var editedTimestamp *timestamppb.Timestamp
	if message.EditedTimestamp != nil {
		editedTimestamp = timestamppb.New(*message.EditedTimestamp)
	}

	return &discordgopb.Message{
		Id:        message.ID,
		ChannelId: message.ChannelID,
		GuildId:   message.GuildID,
		Content:   message.Content,

		Timestamp:       timestamppb.New(message.Timestamp),
		EditedTimestamp: editedTimestamp,
		MentionRoles:    message.MentionRoles,
		Tts:             message.TTS,
		MentionEveryone: message.MentionEveryone,
		Author:          encodeUser(message.Author),

		Attachments: encodeMessageAttachments(message.Attachments),
		Components:  encodeMessageComponents(message.Components),
		Embeds:      encodeMessageEmbeds(message.Embeds),
		Mentions:    encodeUsers(message.Mentions),
		Reactions:   encodeMessageReactions(message.Reactions),

		Pinned:          message.Pinned,
		Type:            int32(message.Type),
		WebhookId:       message.WebhookID,
		Member:          encodeMember(message.Member),
		MentionChannels: encodeChannels(message.MentionChannels),

		Activity:          encodeMessageActivity(message.Activity),
		Application:       encodeMessageApplication(message.Application),
		MessageReference:  encodeMessageReference(message.MessageReference),
		ReferencedMessage: encodeMessage(message.ReferencedMessage),
		Interaction:       encodeMessageInteraction(message.Interaction),

		Flags:        int32(message.Flags),
		Thread:       encodeChannel(message.Thread),
		StickerItems: encodeStickers(message.StickerItems),
	}
}

func encodeMessages(messages []*discordgo.Message) []*discordgopb.Message {
	if messages == nil {
		return nil
	}

	result := make([]*discordgopb.Message, len(messages))

	for i, message := range messages {
		result[i] = encodeMessage(message)
	}

	return result
}

func decodeMessage(message *discordgopb.Message) *discordgo.Message {
	if message == nil {
		return nil
	}

	var editedTimestamp *time.Time
	if message.EditedTimestamp != nil {
		editedTimestamp = util.Ptr(message.EditedTimestamp.AsTime())
	}

	return &discordgo.Message{
		ID:        message.Id,
		ChannelID: message.ChannelId,
		GuildID:   message.GuildId,
		Content:   message.Content,

		Timestamp:       message.Timestamp.AsTime(),
		EditedTimestamp: editedTimestamp,
		MentionRoles:    message.MentionRoles,
		TTS:             message.Tts,
		MentionEveryone: message.MentionEveryone,
		Author:          decodeUser(message.Author),

		Attachments: decodeMessageAttachments(message.Attachments),
		Components:  decodeMessageComponents(message.Components),
		Embeds:      decodeMessageEmbeds(message.Embeds),
		Mentions:    decodeUsers(message.Mentions),
		Reactions:   decodeMessageReactions(message.Reactions),

		Pinned:          message.Pinned,
		Type:            discordgo.MessageType(message.Type),
		WebhookID:       message.WebhookId,
		Member:          decodeMember(message.Member),
		MentionChannels: decodeChannels(message.MentionChannels),

		Activity:          decodeMessageActivity(message.Activity),
		Application:       decodeMessageApplication(message.Application),
		MessageReference:  decodeMessageReference(message.MessageReference),
		ReferencedMessage: decodeMessage(message.ReferencedMessage),
		Interaction:       decodeMessageInteraction(message.Interaction),

		Flags:        discordgo.MessageFlags(message.Flags),
		Thread:       decodeChannel(message.Thread),
		StickerItems: decodeStickers(message.StickerItems),
	}
}

func decodeMessages(messages []*discordgopb.Message) []*discordgo.Message {
	if messages == nil {
		return nil
	}

	result := make([]*discordgo.Message, len(messages))

	for i, message := range messages {
		result[i] = decodeMessage(message)
	}

	return result
}

func encodePermissionOverwrite(overwrite *discordgo.PermissionOverwrite) *discordgopb.PermissionOverwrite {
	if overwrite == nil {
		return nil
	}

	return &discordgopb.PermissionOverwrite{
		Id:    overwrite.ID,
		Type:  int32(overwrite.Type),
		Allow: overwrite.Allow,
		Deny:  overwrite.Deny,
	}
}

func encodePermissionOverwrites(overwrites []*discordgo.PermissionOverwrite) []*discordgopb.PermissionOverwrite {
	if overwrites == nil {
		return nil
	}

	result := make([]*discordgopb.PermissionOverwrite, len(overwrites))

	for i, overwrite := range overwrites {
		result[i] = encodePermissionOverwrite(overwrite)
	}

	return result
}

func decodePermissionOverwrite(overwrite *discordgopb.PermissionOverwrite) *discordgo.PermissionOverwrite {
	if overwrite == nil {
		return nil
	}

	return &discordgo.PermissionOverwrite{
		ID:    overwrite.Id,
		Type:  discordgo.PermissionOverwriteType(overwrite.Type),
		Allow: overwrite.Allow,
		Deny:  overwrite.Deny,
	}
}

func decodePermissionOverwrites(overwrites []*discordgopb.PermissionOverwrite) []*discordgo.PermissionOverwrite {
	if overwrites == nil {
		return nil
	}

	result := make([]*discordgo.PermissionOverwrite, len(overwrites))

	for i, overwrite := range overwrites {
		result[i] = decodePermissionOverwrite(overwrite)
	}

	return result
}

func encodeThreadMetadata(metadata *discordgo.ThreadMetadata) *discordgopb.ThreadMetadata {
	if metadata == nil {
		return nil
	}

	return &discordgopb.ThreadMetadata{
		Archived:            metadata.Archived,
		AutoArchiveDuration: int32(metadata.AutoArchiveDuration),
		ArchiveTimestamp:    timestamppb.New(metadata.ArchiveTimestamp),
		Locked:              metadata.Locked,
		Invitable:           metadata.Invitable,
	}
}

func decodeThreadMetadata(metadata *discordgopb.ThreadMetadata) *discordgo.ThreadMetadata {
	if metadata == nil {
		return nil
	}

	return &discordgo.ThreadMetadata{
		Archived:            metadata.Archived,
		AutoArchiveDuration: int(metadata.AutoArchiveDuration),
		ArchiveTimestamp:    metadata.ArchiveTimestamp.AsTime(),
		Locked:              metadata.Locked,
		Invitable:           metadata.Invitable,
	}
}

func encodeThreadMember(member *discordgo.ThreadMember) *discordgopb.ThreadMember {
	if member == nil {
		return nil
	}

	return &discordgopb.ThreadMember{
		Id:            member.ID,
		UserId:        member.UserID,
		JoinTimestamp: timestamppb.New(member.JoinTimestamp),
		Flags:         int32(member.Flags),
	}
}

func encodeThreadMembers(members []*discordgo.ThreadMember) []*discordgopb.ThreadMember {
	if members == nil {
		return nil
	}

	result := make([]*discordgopb.ThreadMember, len(members))

	for i, member := range members {
		result[i] = encodeThreadMember(member)
	}

	return result
}

func decodeThreadMember(member *discordgopb.ThreadMember) *discordgo.ThreadMember {
	if member == nil {
		return nil
	}

	return &discordgo.ThreadMember{
		ID:            member.Id,
		UserID:        member.UserId,
		JoinTimestamp: member.JoinTimestamp.AsTime(),
		Flags:         int(member.Flags),
	}
}

func decodeThreadMembers(members []*discordgopb.ThreadMember) []*discordgo.ThreadMember {
	if members == nil {
		return nil
	}

	result := make([]*discordgo.ThreadMember, len(members))

	for i, member := range members {
		result[i] = decodeThreadMember(member)
	}

	return result
}

func encodeForumTag(tag discordgo.ForumTag) *discordgopb.ForumTag {
	return &discordgopb.ForumTag{
		Id:        tag.ID,
		Name:      tag.Name,
		Moderated: tag.Moderated,
		EmojiId:   tag.EmojiID,
		EmojiName: tag.EmojiName,
	}
}

func encodeForumTags(tags []discordgo.ForumTag) []*discordgopb.ForumTag {
	if tags == nil {
		return nil
	}

	result := make([]*discordgopb.ForumTag, len(tags))

	for i, tag := range tags {
		result[i] = encodeForumTag(tag)
	}

	return result
}

func decodeForumTag(tag *discordgopb.ForumTag) discordgo.ForumTag {
	return discordgo.ForumTag{
		ID:        tag.Id,
		Name:      tag.Name,
		Moderated: tag.Moderated,
		EmojiID:   tag.EmojiId,
		EmojiName: tag.EmojiName,
	}
}

func decodeForumTags(tags []*discordgopb.ForumTag) []discordgo.ForumTag {
	if tags == nil {
		return nil
	}

	result := make([]discordgo.ForumTag, len(tags))

	for i, tag := range tags {
		result[i] = decodeForumTag(tag)
	}

	return result
}

func encodeForumDefaultReaction(reaction *discordgo.ForumDefaultReaction) *discordgopb.ForumDefaultReaction {
	if reaction == nil {
		return nil
	}

	return &discordgopb.ForumDefaultReaction{
		EmojiId:   reaction.EmojiID,
		EmojiName: reaction.EmojiName,
	}
}

func decodeForumDefaultReaction(reaction *discordgopb.ForumDefaultReaction) *discordgo.ForumDefaultReaction {
	if reaction == nil {
		return nil
	}

	return &discordgo.ForumDefaultReaction{
		EmojiID:   reaction.EmojiId,
		EmojiName: reaction.EmojiName,
	}
}
