package codec

import (
	"bytes"
	"io"

	"github.com/EliasStar/BacoTell/internal/proto/discordgopb"
	"github.com/bwmarrin/discordgo"
)

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

func encodeMessageAttachments(attachments []*discordgo.MessageAttachment) []*discordgopb.MessageAttachment {
	result := make([]*discordgopb.MessageAttachment, len(attachments))

	for i, attachment := range attachments {
		result[i] = encodeMessageAttachment(attachment)
	}

	return result
}

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

func decodeMessageAttachments(attachments []*discordgopb.MessageAttachment) []*discordgo.MessageAttachment {
	result := make([]*discordgo.MessageAttachment, len(attachments))

	for i, attachment := range attachments {
		result[i] = decodeMessageAttachment(attachment)
	}

	return result
}

func encodeMessageComponent(component discordgo.MessageComponent) *discordgopb.MessageComponent {
	return nil // TODO
}

func encodeMessageComponents(components []discordgo.MessageComponent) []*discordgopb.MessageComponent {
	result := make([]*discordgopb.MessageComponent, len(components))

	for i, component := range components {
		result[i] = encodeMessageComponent(component)
	}

	return result
}

func decodeMessageComponent(component *discordgopb.MessageComponent) discordgo.MessageComponent {
	return nil // TODO
}

func decodeMessageComponents(components []*discordgopb.MessageComponent) []discordgo.MessageComponent {
	result := make([]discordgo.MessageComponent, len(components))

	for i, component := range components {
		result[i] = decodeMessageComponent(component)
	}

	return result
}

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

func encodeMessageEmbeds(embeds []*discordgo.MessageEmbed) []*discordgopb.MessageEmbed {
	result := make([]*discordgopb.MessageEmbed, len(embeds))

	for i, embed := range embeds {
		result[i] = encodeMessageEmbed(embed)
	}

	return result
}

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

func decodeMessageEmbeds(embeds []*discordgopb.MessageEmbed) []*discordgo.MessageEmbed {
	result := make([]*discordgo.MessageEmbed, len(embeds))

	for i, embed := range embeds {
		result[i] = decodeMessageEmbed(embed)
	}

	return result
}

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

func encodeMessageEmbedProvider(provider *discordgo.MessageEmbedProvider) *discordgopb.MessageEmbedProvider {
	if provider == nil {
		return nil
	}

	return &discordgopb.MessageEmbedProvider{
		Url:  provider.URL,
		Name: provider.Name,
	}
}

func decodeMessageEmbedProvider(provider *discordgopb.MessageEmbedProvider) *discordgo.MessageEmbedProvider {
	if provider == nil {
		return nil
	}

	return &discordgo.MessageEmbedProvider{
		URL:  provider.Url,
		Name: provider.Name,
	}
}

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

func encodeMessageEmbedFields(fields []*discordgo.MessageEmbedField) []*discordgopb.MessageEmbedField {
	result := make([]*discordgopb.MessageEmbedField, len(fields))

	for i, field := range fields {
		result[i] = encodeMessageEmbedField(field)
	}

	return result
}

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

func decodeMessageEmbedFields(fields []*discordgopb.MessageEmbedField) []*discordgo.MessageEmbedField {
	result := make([]*discordgo.MessageEmbedField, len(fields))

	for i, field := range fields {
		result[i] = decodeMessageEmbedField(field)
	}

	return result
}

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

func encodeMessageReactions(reactions []*discordgo.MessageReactions) []*discordgopb.MessageReactions {
	result := make([]*discordgopb.MessageReactions, len(reactions))

	for i, reaction := range reactions {
		result[i] = encodeMessageReaction(reaction)
	}

	return result
}

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

func decodeMessageReactions(reactions []*discordgopb.MessageReactions) []*discordgo.MessageReactions {
	result := make([]*discordgo.MessageReactions, len(reactions))

	for i, reaction := range reactions {
		result[i] = decodeMessageReaction(reaction)
	}

	return result
}

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

func encodeMessageActivity(activity *discordgo.MessageActivity) *discordgopb.MessageActivity {
	if activity == nil {
		return nil
	}

	return &discordgopb.MessageActivity{
		Type:    int32(activity.Type),
		PartyId: activity.PartyID,
	}
}

func decodeMessageActivity(activity *discordgopb.MessageActivity) *discordgo.MessageActivity {
	if activity == nil {
		return nil
	}

	return &discordgo.MessageActivity{
		Type:    discordgo.MessageActivityType(activity.Type),
		PartyID: activity.PartyId,
	}
}

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

func encodeStickers(stickers []*discordgo.Sticker) []*discordgopb.Sticker {
	result := make([]*discordgopb.Sticker, len(stickers))

	for i, sticker := range stickers {
		result[i] = encodeSticker(sticker)
	}

	return result
}

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

func decodeStickers(stickers []*discordgopb.Sticker) []*discordgo.Sticker {
	result := make([]*discordgo.Sticker, len(stickers))

	for i, sticker := range stickers {
		result[i] = decodeSticker(sticker)
	}

	return result
}

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

func encodeFiles(files []*discordgo.File) []*discordgopb.File {
	result := make([]*discordgopb.File, len(files))

	for i, file := range files {
		result[i] = encodeFile(file)
	}

	return result
}

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

func decodeFiles(files []*discordgopb.File) []*discordgo.File {
	result := make([]*discordgo.File, len(files))

	for i, file := range files {
		result[i] = decodeFile(file)
	}

	return result
}

func _encodeParse(allowedMentions []discordgo.AllowedMentionType) []string {
	result := make([]string, len(allowedMentions))

	for i, allowedMention := range allowedMentions {
		result[i] = string(allowedMention)
	}

	return result
}

func _decodeParse(allowedMentions []string) []discordgo.AllowedMentionType {
	result := make([]discordgo.AllowedMentionType, len(allowedMentions))

	for i, allowedMention := range allowedMentions {
		result[i] = discordgo.AllowedMentionType(allowedMention)
	}

	return result
}
