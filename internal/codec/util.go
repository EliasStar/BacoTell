package codec

import "github.com/bwmarrin/discordgo"

// _encodeLocalizations encodes a map of locales and strings to a map of strings.
func _encodeLocalizations(localizations map[discordgo.Locale]string) map[string]string {
	if localizations == nil {
		return nil
	}

	result := make(map[string]string)

	for locale, str := range localizations {
		result[string(locale)] = str
	}

	return result
}

// _decodeLocalizations decodes a map of strings to a map of locales and strings.
func _decodeLocalizations(localizations map[string]string) map[discordgo.Locale]string {
	if localizations == nil {
		return nil
	}

	result := make(map[discordgo.Locale]string)

	for locale, str := range localizations {
		result[discordgo.Locale(locale)] = str
	}

	return result
}

// _encodeChannelTypes encodes a slice of channel types to a slice of int32.
func _encodeChannelTypes(types []discordgo.ChannelType) []int32 {
	if types == nil {
		return nil
	}

	result := make([]int32, len(types))

	for i, typ := range types {
		result[i] = int32(typ)
	}

	return result
}

// _decodeChannelTypes decodes a slice of int32 to a slice of channel types.
func _decodeChannelTypes(types []int32) []discordgo.ChannelType {
	if types == nil {
		return nil
	}

	result := make([]discordgo.ChannelType, len(types))

	for i, typ := range types {
		result[i] = discordgo.ChannelType(typ)
	}

	return result
}

// _encodeParse encodes a slice of allowed mention types to a slice of strings.
func _encodeParse(allowedMentions []discordgo.AllowedMentionType) []string {
	if allowedMentions == nil {
		return nil
	}

	result := make([]string, len(allowedMentions))

	for i, allowedMention := range allowedMentions {
		result[i] = string(allowedMention)
	}

	return result
}

// _decodeParse decodes a slice of strings to a slice of allowed mention types.
func _decodeParse(allowedMentions []string) []discordgo.AllowedMentionType {
	if allowedMentions == nil {
		return nil
	}

	result := make([]discordgo.AllowedMentionType, len(allowedMentions))

	for i, allowedMention := range allowedMentions {
		result[i] = discordgo.AllowedMentionType(allowedMention)
	}

	return result
}
