package codec

import (
	"bytes"
	"io"

	"github.com/EliasStar/BacoTell/internal/proto/discordgopb"
	"github.com/bwmarrin/discordgo"
)

// encodeMessageAllowedMention encodes a discordgo.MessageAllowedMentions into a discordgopb.MessageAllowedMentions.
func encodeMessageAllowedMention(mention *discordgo.MessageAllowedMentions) *discordgopb.MessageAllowedMentions {
	if mention == nil {
		return nil
	}

	return &discordgopb.MessageAllowedMentions{
		Parse:       _encodeParse(mention.Parse),
		Roles:       mention.Roles,
		Users:       mention.Users,
		RepliedUser: mention.RepliedUser,
	}
}

// decodeMessageAllowedMention decodes a discordgopb.MessageAllowedMentions into a discordgo.MessageAllowedMentions.
func decodeMessageAllowedMention(mention *discordgopb.MessageAllowedMentions) *discordgo.MessageAllowedMentions {
	if mention == nil {
		return nil
	}

	return &discordgo.MessageAllowedMentions{
		Parse:       _decodeParse(mention.Parse),
		Roles:       mention.Roles,
		Users:       mention.Users,
		RepliedUser: mention.RepliedUser,
	}
}

// encodeMessageAttachment encodes a discordgo.MessageAttachment into a discordgopb.MessageAttachment.
func encodeMessageAttachment(attachment *discordgo.MessageAttachment) *discordgopb.MessageAttachment {
	if attachment == nil {
		return nil
	}

	return &discordgopb.MessageAttachment{
		Id:          attachment.ID,
		Url:         attachment.URL,
		ProxyUrl:    attachment.ProxyURL,
		Filename:    attachment.Filename,
		ContentType: attachment.ContentType,

		Height: int32(attachment.Height),
		Width:  int32(attachment.Width),
		Size:   int32(attachment.Size),

		Ephemeral: attachment.Ephemeral,
	}
}

// encodeMessageAttachments encodes multiple discordgo.MessageAttachment into multiple discordgopb.MessageAttachment.
func encodeMessageAttachments(attachments []*discordgo.MessageAttachment) []*discordgopb.MessageAttachment {
	if attachments == nil {
		return nil
	}

	result := make([]*discordgopb.MessageAttachment, len(attachments))

	for i, attachment := range attachments {
		result[i] = encodeMessageAttachment(attachment)
	}

	return result
}

// decodeMessageAttachment decodes a discordgopb.MessageAttachment into a discordgo.MessageAttachment.
func decodeMessageAttachment(attachment *discordgopb.MessageAttachment) *discordgo.MessageAttachment {
	if attachment == nil {
		return nil
	}

	return &discordgo.MessageAttachment{
		ID:          attachment.Id,
		URL:         attachment.Url,
		ProxyURL:    attachment.ProxyUrl,
		Filename:    attachment.Filename,
		ContentType: attachment.ContentType,

		Height: int(attachment.Height),
		Width:  int(attachment.Width),
		Size:   int(attachment.Size),

		Ephemeral: attachment.Ephemeral,
	}
}

// decodeMessageAttachments decodes multiple discordgopb.MessageAttachment into multiple discordgo.MessageAttachment.
func decodeMessageAttachments(attachments []*discordgopb.MessageAttachment) []*discordgo.MessageAttachment {
	if attachments == nil {
		return nil
	}

	result := make([]*discordgo.MessageAttachment, len(attachments))

	for i, attachment := range attachments {
		result[i] = decodeMessageAttachment(attachment)
	}

	return result
}

// encodeMessageEmbed encodes a discordgo.MessageEmbed into a discordgopb.MessageEmbed.
func encodeMessageEmbed(embed *discordgo.MessageEmbed) *discordgopb.MessageEmbed {
	if embed == nil {
		return nil
	}

	return &discordgopb.MessageEmbed{
		Url:         embed.URL,
		Type:        string(embed.Type),
		Title:       embed.Title,
		Description: embed.Description,
		Timestamp:   embed.Timestamp,
		Color:       int32(embed.Color),

		Footer:    encodeMessageEmbedFooter(embed.Footer),
		Image:     encodeMessageEmbedImage(embed.Image),
		Thumbnail: encodeMessageEmbedThumbnail(embed.Thumbnail),
		Video:     encodeMessageEmbedVideo(embed.Video),
		Provider:  encodeMessageEmbedProvider(embed.Provider),
		Author:    encodeMessageEmbedAuthor(embed.Author),
		Fields:    encodeMessageEmbedFields(embed.Fields),
	}
}

// encodeMessageEmbeds encodes multiple discordgo.MessageEmbed into multiple discordgopb.MessageEmbed.
func encodeMessageEmbeds(embeds []*discordgo.MessageEmbed) []*discordgopb.MessageEmbed {
	if embeds == nil {
		return nil
	}

	result := make([]*discordgopb.MessageEmbed, len(embeds))

	for i, embed := range embeds {
		result[i] = encodeMessageEmbed(embed)
	}

	return result
}

// decodeMessageEmbed decodes a discordgopb.MessageEmbed into a discordgo.MessageEmbed.
func decodeMessageEmbed(embed *discordgopb.MessageEmbed) *discordgo.MessageEmbed {
	if embed == nil {
		return nil
	}

	return &discordgo.MessageEmbed{
		URL:         embed.Url,
		Type:        discordgo.EmbedType(embed.Type),
		Title:       embed.Title,
		Description: embed.Description,
		Timestamp:   embed.Timestamp,
		Color:       int(embed.Color),

		Footer:    decodeMessageEmbedFooter(embed.Footer),
		Image:     decodeMessageEmbedImage(embed.Image),
		Thumbnail: decodeMessageEmbedThumbnail(embed.Thumbnail),
		Video:     decodeMessageEmbedVideo(embed.Video),
		Provider:  decodeMessageEmbedProvider(embed.Provider),
		Author:    decodeMessageEmbedAuthor(embed.Author),
		Fields:    decodeMessageEmbedFields(embed.Fields),
	}
}

// decodeMessageEmbeds decodes multiple discordgopb.MessageEmbed into multiple discordgo.MessageEmbed.
func decodeMessageEmbeds(embeds []*discordgopb.MessageEmbed) []*discordgo.MessageEmbed {
	if embeds == nil {
		return nil
	}

	result := make([]*discordgo.MessageEmbed, len(embeds))

	for i, embed := range embeds {
		result[i] = decodeMessageEmbed(embed)
	}

	return result
}

// encodeMessageEmbedFooter encodes a discordgo.MessageEmbedFooter into a discordgopb.MessageEmbedFooter.
func encodeMessageEmbedFooter(footer *discordgo.MessageEmbedFooter) *discordgopb.MessageEmbedFooter {
	if footer == nil {
		return nil
	}

	return &discordgopb.MessageEmbedFooter{
		Text:         footer.Text,
		IconUrl:      footer.IconURL,
		ProxyIconUrl: footer.ProxyIconURL,
	}
}

// decodeMessageEmbedFooter decodes a discordgopb.MessageEmbedFooter into a discordgo.MessageEmbedFooter.
func decodeMessageEmbedFooter(footer *discordgopb.MessageEmbedFooter) *discordgo.MessageEmbedFooter {
	if footer == nil {
		return nil
	}

	return &discordgo.MessageEmbedFooter{
		Text:         footer.Text,
		IconURL:      footer.IconUrl,
		ProxyIconURL: footer.ProxyIconUrl,
	}
}

// encodeMessageEmbedImage encodes a discordgo.MessageEmbedImage into a discordgopb.MessageEmbedImage.
func encodeMessageEmbedImage(image *discordgo.MessageEmbedImage) *discordgopb.MessageEmbedImage {
	if image == nil {
		return nil
	}

	return &discordgopb.MessageEmbedImage{
		Url:      image.URL,
		ProxyUrl: image.ProxyURL,
		Width:    int32(image.Width),
		Height:   int32(image.Height),
	}
}

// decodeMessageEmbedImage decodes a discordgopb.MessageEmbedImage into a discordgo.MessageEmbedImage.
func decodeMessageEmbedImage(image *discordgopb.MessageEmbedImage) *discordgo.MessageEmbedImage {
	if image == nil {
		return nil
	}

	return &discordgo.MessageEmbedImage{
		URL:      image.Url,
		ProxyURL: image.ProxyUrl,
		Width:    int(image.Width),
		Height:   int(image.Height),
	}
}

// encodeMessageEmbedThumbnail encodes a discordgo.MessageEmbedThumbnail into a discordgopb.MessageEmbedThumbnail.
func encodeMessageEmbedThumbnail(thumbnail *discordgo.MessageEmbedThumbnail) *discordgopb.MessageEmbedThumbnail {
	if thumbnail == nil {
		return nil
	}

	return &discordgopb.MessageEmbedThumbnail{
		Url:      thumbnail.URL,
		ProxyUrl: thumbnail.ProxyURL,
		Width:    int32(thumbnail.Width),
		Height:   int32(thumbnail.Height),
	}
}

// decodeMessageEmbedThumbnail decodes a discordgopb.MessageEmbedThumbnail into a discordgo.MessageEmbedThumbnail.
func decodeMessageEmbedThumbnail(thumbnail *discordgopb.MessageEmbedThumbnail) *discordgo.MessageEmbedThumbnail {
	if thumbnail == nil {
		return nil
	}

	return &discordgo.MessageEmbedThumbnail{
		URL:      thumbnail.Url,
		ProxyURL: thumbnail.ProxyUrl,
		Width:    int(thumbnail.Width),
		Height:   int(thumbnail.Height),
	}
}

// encodeMessageEmbedVideo encodes a discordgo.MessageEmbedVideo into a discordgopb.MessageEmbedVideo.
func encodeMessageEmbedVideo(video *discordgo.MessageEmbedVideo) *discordgopb.MessageEmbedVideo {
	if video == nil {
		return nil
	}

	return &discordgopb.MessageEmbedVideo{
		Url:    video.URL,
		Width:  int32(video.Width),
		Height: int32(video.Height),
	}
}

// decodeMessageEmbedVideo decodes a discordgopb.MessageEmbedVideo into a discordgo.MessageEmbedVideo.
func decodeMessageEmbedVideo(video *discordgopb.MessageEmbedVideo) *discordgo.MessageEmbedVideo {
	if video == nil {
		return nil
	}

	return &discordgo.MessageEmbedVideo{
		URL:    video.Url,
		Width:  int(video.Width),
		Height: int(video.Height),
	}
}

// encodeMessageEmbedProvider encodes a discordgo.MessageEmbedProvider into a discordgopb.MessageEmbedProvider.
func encodeMessageEmbedProvider(provider *discordgo.MessageEmbedProvider) *discordgopb.MessageEmbedProvider {
	if provider == nil {
		return nil
	}

	return &discordgopb.MessageEmbedProvider{
		Url:  provider.URL,
		Name: provider.Name,
	}
}

// decodeMessageEmbedProvider decodes a discordgopb.MessageEmbedProvider into a discordgo.MessageEmbedProvider.
func decodeMessageEmbedProvider(provider *discordgopb.MessageEmbedProvider) *discordgo.MessageEmbedProvider {
	if provider == nil {
		return nil
	}

	return &discordgo.MessageEmbedProvider{
		URL:  provider.Url,
		Name: provider.Name,
	}
}

// encodeMessageEmbedAuthor encodes a discordgo.MessageEmbedAuthor into a discordgopb.MessageEmbedAuthor.
func encodeMessageEmbedAuthor(author *discordgo.MessageEmbedAuthor) *discordgopb.MessageEmbedAuthor {
	if author == nil {
		return nil
	}

	return &discordgopb.MessageEmbedAuthor{
		Url:          author.URL,
		Name:         author.Name,
		IconUrl:      author.IconURL,
		ProxyIconUrl: author.ProxyIconURL,
	}
}

// decodeMessageEmbedAuthor decodes a discordgopb.MessageEmbedAuthor into a discordgo.MessageEmbedAuthor.
func decodeMessageEmbedAuthor(author *discordgopb.MessageEmbedAuthor) *discordgo.MessageEmbedAuthor {
	if author == nil {
		return nil
	}

	return &discordgo.MessageEmbedAuthor{
		URL:          author.Url,
		Name:         author.Name,
		IconURL:      author.IconUrl,
		ProxyIconURL: author.ProxyIconUrl,
	}
}

// encodeMessageEmbedField encodes a discordgo.MessageEmbedField into a discordgopb.MessageEmbedField.
func encodeMessageEmbedField(field *discordgo.MessageEmbedField) *discordgopb.MessageEmbedField {
	if field == nil {
		return nil
	}

	return &discordgopb.MessageEmbedField{
		Name:   field.Name,
		Value:  field.Value,
		Inline: field.Inline,
	}
}

// encodeMessageEmbedFields encodes multiple discordgo.MessageEmbedField into multiple discordgopb.MessageEmbedField.
func encodeMessageEmbedFields(fields []*discordgo.MessageEmbedField) []*discordgopb.MessageEmbedField {
	if fields == nil {
		return nil
	}

	result := make([]*discordgopb.MessageEmbedField, len(fields))

	for i, field := range fields {
		result[i] = encodeMessageEmbedField(field)
	}

	return result
}

// decodeMessageEmbedField decodes a discordgopb.MessageEmbedField into a discordgo.MessageEmbedField.
func decodeMessageEmbedField(field *discordgopb.MessageEmbedField) *discordgo.MessageEmbedField {
	if field == nil {
		return nil
	}

	return &discordgo.MessageEmbedField{
		Name:   field.Name,
		Value:  field.Value,
		Inline: field.Inline,
	}
}

// decodeMessageEmbedFields decodes multiple discordgopb.MessageEmbedField into multiple discordgo.MessageEmbedField.
func decodeMessageEmbedFields(fields []*discordgopb.MessageEmbedField) []*discordgo.MessageEmbedField {
	if fields == nil {
		return nil
	}

	result := make([]*discordgo.MessageEmbedField, len(fields))

	for i, field := range fields {
		result[i] = decodeMessageEmbedField(field)
	}

	return result
}

// encodeMessageReaction encodes a discordgo.MessageReactions into a discordgopb.MessageReactions.
func encodeMessageReaction(reaction *discordgo.MessageReactions) *discordgopb.MessageReactions {
	if reaction == nil {
		return nil
	}

	return &discordgopb.MessageReactions{
		Count: int32(reaction.Count),
		Me:    reaction.Me,
		Emoji: encodeEmoji(reaction.Emoji),
	}
}

// encodeMessageReactions encodes multiple discordgo.MessageReactions into multiple discordgopb.MessageReactions.
func encodeMessageReactions(reactions []*discordgo.MessageReactions) []*discordgopb.MessageReactions {
	if reactions == nil {
		return nil
	}

	result := make([]*discordgopb.MessageReactions, len(reactions))

	for i, reaction := range reactions {
		result[i] = encodeMessageReaction(reaction)
	}

	return result
}

// decodeMessageReaction decodes a discordgopb.MessageReactions into a discordgo.MessageReactions.
func decodeMessageReaction(reaction *discordgopb.MessageReactions) *discordgo.MessageReactions {
	if reaction == nil {
		return nil
	}

	return &discordgo.MessageReactions{
		Count: int(reaction.Count),
		Me:    reaction.Me,
		Emoji: decodeEmoji(reaction.Emoji),
	}
}

// decodeMessageReactions decodes multiple discordgopb.MessageReactions into multiple discordgo.MessageReactions.
func decodeMessageReactions(reactions []*discordgopb.MessageReactions) []*discordgo.MessageReactions {
	if reactions == nil {
		return nil
	}

	result := make([]*discordgo.MessageReactions, len(reactions))

	for i, reaction := range reactions {
		result[i] = decodeMessageReaction(reaction)
	}

	return result
}

// encodeEmoji encodes a discordgo.Emoji into a discordgopb.Emoji.
func encodeEmoji(emoji *discordgo.Emoji) *discordgopb.Emoji {
	if emoji == nil {
		return nil
	}

	return &discordgopb.Emoji{
		Id:   emoji.ID,
		Name: emoji.Name,

		Roles: emoji.Roles,
		User:  encodeUser(emoji.User),

		RequireColons: emoji.RequireColons,
		Managed:       emoji.Managed,
		Animated:      emoji.Animated,
		Available:     emoji.Available,
	}
}

// decodeEmoji decodes a discordgopb.Emoji into a discordgo.Emoji.
func decodeEmoji(emoji *discordgopb.Emoji) *discordgo.Emoji {
	if emoji == nil {
		return nil
	}

	return &discordgo.Emoji{
		ID:   emoji.Id,
		Name: emoji.Name,

		Roles: emoji.Roles,
		User:  decodeUser(emoji.User),

		RequireColons: emoji.RequireColons,
		Managed:       emoji.Managed,
		Animated:      emoji.Animated,
		Available:     emoji.Available,
	}
}

// encodeMessageActivity encodes a discordgo.MessageActivity into a discordgopb.MessageActivity.
func encodeMessageActivity(activity *discordgo.MessageActivity) *discordgopb.MessageActivity {
	if activity == nil {
		return nil
	}

	return &discordgopb.MessageActivity{
		Type:    int32(activity.Type),
		PartyId: activity.PartyID,
	}
}

// decodeMessageActivity decodes a discordgopb.MessageActivity into a discordgo.MessageActivity.
func decodeMessageActivity(activity *discordgopb.MessageActivity) *discordgo.MessageActivity {
	if activity == nil {
		return nil
	}

	return &discordgo.MessageActivity{
		Type:    discordgo.MessageActivityType(activity.Type),
		PartyID: activity.PartyId,
	}
}

// encodeMessageApplication encodes a discordgo.MessageApplication into a discordgopb.MessageApplication.
func encodeMessageApplication(application *discordgo.MessageApplication) *discordgopb.MessageApplication {
	if application == nil {
		return nil
	}

	return &discordgopb.MessageApplication{
		Id:          application.ID,
		CoverImage:  application.CoverImage,
		Description: application.Description,
		Icon:        application.Icon,
		Name:        application.Name,
	}
}

// decodeMessageApplication decodes a discordgopb.MessageApplication into a discordgo.MessageApplication.
func decodeMessageApplication(application *discordgopb.MessageApplication) *discordgo.MessageApplication {
	if application == nil {
		return nil
	}

	return &discordgo.MessageApplication{
		ID:          application.Id,
		CoverImage:  application.CoverImage,
		Description: application.Description,
		Icon:        application.Icon,
		Name:        application.Name,
	}
}

// encodeMessageReference encodes a discordgo.MessageReference into a discordgopb.MessageReference.
func encodeMessageReference(reference *discordgo.MessageReference) *discordgopb.MessageReference {
	if reference == nil {
		return nil
	}

	return &discordgopb.MessageReference{
		MessageId: reference.MessageID,
		ChannelId: reference.ChannelID,
		GuildId:   reference.GuildID,
	}
}

// decodeMessageReference decodes a discordgopb.MessageReference into a discordgo.MessageReference.
func decodeMessageReference(reference *discordgopb.MessageReference) *discordgo.MessageReference {
	if reference == nil {
		return nil
	}

	return &discordgo.MessageReference{
		MessageID: reference.MessageId,
		ChannelID: reference.ChannelId,
		GuildID:   reference.GuildId,
	}
}

// encodeMessageInteraction encodes a discordgo.MessageInteraction into a discordgopb.MessageInteraction.
func encodeMessageInteraction(interaction *discordgo.MessageInteraction) *discordgopb.MessageInteraction {
	if interaction == nil {
		return nil
	}

	return &discordgopb.MessageInteraction{
		Id:   interaction.ID,
		Type: uint32(interaction.Type),
		Name: interaction.Name,

		User:   encodeUser(interaction.User),
		Member: encodeMember(interaction.Member),
	}
}

// decodeMessageInteraction decodes a discordgopb.MessageInteraction into a discordgo.MessageInteraction.
func decodeMessageInteraction(interaction *discordgopb.MessageInteraction) *discordgo.MessageInteraction {
	if interaction == nil {
		return nil
	}

	return &discordgo.MessageInteraction{
		ID:   interaction.Id,
		Type: discordgo.InteractionType(interaction.Type),
		Name: interaction.Name,

		User:   decodeUser(interaction.User),
		Member: decodeMember(interaction.Member),
	}
}

// encodeSticker encodes a discordgo.Sticker into a discordgopb.Sticker.
func encodeSticker(sticker *discordgo.Sticker) *discordgopb.Sticker {
	if sticker == nil {
		return nil
	}

	return &discordgopb.Sticker{
		Id:          sticker.ID,
		PackId:      sticker.PackID,
		Name:        sticker.Name,
		Description: sticker.Description,
		Tags:        sticker.Tags,

		Type:       int32(sticker.Type),
		FormatType: int32(sticker.FormatType),

		Available: sticker.Available,
		GuildId:   sticker.GuildID,
		User:      encodeUser(sticker.User),
		SortValue: int32(sticker.SortValue),
	}
}

// encodeStickers encodes multiple discordgo.Sticker into multiple discordgopb.Sticker.
func encodeStickers(stickers []*discordgo.Sticker) []*discordgopb.Sticker {
	if stickers == nil {
		return nil
	}

	result := make([]*discordgopb.Sticker, len(stickers))

	for i, sticker := range stickers {
		result[i] = encodeSticker(sticker)
	}

	return result
}

// decodeSticker decodes a discordgopb.Sticker into a discordgo.Sticker.
func decodeSticker(sticker *discordgopb.Sticker) *discordgo.Sticker {
	if sticker == nil {
		return nil
	}

	return &discordgo.Sticker{
		ID:          sticker.Id,
		PackID:      sticker.PackId,
		Name:        sticker.Name,
		Description: sticker.Description,
		Tags:        sticker.Tags,

		Type:       discordgo.StickerType(sticker.Type),
		FormatType: discordgo.StickerFormat(sticker.FormatType),

		Available: sticker.Available,
		GuildID:   sticker.GuildId,
		User:      decodeUser(sticker.User),
		SortValue: int(sticker.SortValue),
	}
}

// decodeStickers decodes multiple discordgopb.Sticker into multiple discordgo.Sticker.
func decodeStickers(stickers []*discordgopb.Sticker) []*discordgo.Sticker {
	if stickers == nil {
		return nil
	}

	result := make([]*discordgo.Sticker, len(stickers))

	for i, sticker := range stickers {
		result[i] = decodeSticker(sticker)
	}

	return result
}

// encodeFile encodes a discordgo.File into a discordgopb.File.
func encodeFile(file *discordgo.File) *discordgopb.File {
	if file == nil {
		return nil
	}

	content, _ := io.ReadAll(file.Reader)

	return &discordgopb.File{
		Name:        file.Name,
		ContentType: file.ContentType,
		Content:     content,
	}
}

// encodeFiles encodes multiple discordgo.File into multiple discordgopb.File.
func encodeFiles(files []*discordgo.File) []*discordgopb.File {
	if files == nil {
		return nil
	}

	result := make([]*discordgopb.File, len(files))

	for i, file := range files {
		result[i] = encodeFile(file)
	}

	return result
}

// decodeFile decodes a discordgopb.File into a discordgo.File.
func decodeFile(file *discordgopb.File) *discordgo.File {
	if file == nil {
		return nil
	}

	return &discordgo.File{
		Name:        file.Name,
		ContentType: file.ContentType,
		Reader:      bytes.NewReader(file.Content),
	}
}

// decodeFiles decodes multiple discordgopb.File into multiple discordgo.File.
func decodeFiles(files []*discordgopb.File) []*discordgo.File {
	if files == nil {
		return nil
	}

	result := make([]*discordgo.File, len(files))

	for i, file := range files {
		result[i] = decodeFile(file)
	}

	return result
}
