package codec

import "github.com/bwmarrin/discordgo"

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
