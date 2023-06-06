package codec

import (
	"github.com/EliasStar/BacoTell/internal/proto/discordgopb"
	util "github.com/EliasStar/BacoTell/pkg/bacotell_util"
	"github.com/bwmarrin/discordgo"
)

func encodeMessageComponent(component discordgo.MessageComponent) *discordgopb.MessageComponent {
	switch cpt := component.(type) {
	case discordgo.ActionsRow:
		return &discordgopb.MessageComponent{Component: &discordgopb.MessageComponent_ActionsRow{ActionsRow: encodeActionsRow(&cpt)}}
	case discordgo.Button:
		return &discordgopb.MessageComponent{Component: &discordgopb.MessageComponent_Button{Button: encodeButton(&cpt)}}
	case discordgo.SelectMenu:
		return &discordgopb.MessageComponent{Component: &discordgopb.MessageComponent_SelectMenu{SelectMenu: encodeSelectMenu(&cpt)}}
	case discordgo.TextInput:
		return &discordgopb.MessageComponent{Component: &discordgopb.MessageComponent_TextInput{TextInput: encodeTextInput(&cpt)}}
	default:
		return nil
	}
}

func encodeMessageComponents(components []discordgo.MessageComponent) []*discordgopb.MessageComponent {
	if components == nil {
		return nil
	}

	result := make([]*discordgopb.MessageComponent, len(components))

	for i, component := range components {
		result[i] = encodeMessageComponent(component)
	}

	return result
}

func decodeMessageComponent(component *discordgopb.MessageComponent) (result discordgo.MessageComponent) {
	switch cpt := component.Component.(type) {
	case *discordgopb.MessageComponent_ActionsRow:
		return decodeActionsRow(cpt.ActionsRow)
	case *discordgopb.MessageComponent_Button:
		return decodeButton(cpt.Button)
	case *discordgopb.MessageComponent_SelectMenu:
		return decodeSelectMenu(cpt.SelectMenu)
	case *discordgopb.MessageComponent_TextInput:
		return decodeTextInput(cpt.TextInput)
	default:
		return nil
	}
}

func decodeMessageComponents(components []*discordgopb.MessageComponent) []discordgo.MessageComponent {
	if components == nil {
		return nil
	}

	result := make([]discordgo.MessageComponent, len(components))

	for i, component := range components {
		result[i] = decodeMessageComponent(component)
	}

	return result
}

func encodeActionsRow(row *discordgo.ActionsRow) *discordgopb.ActionsRow {
	if row == nil {
		return nil
	}

	return &discordgopb.ActionsRow{Components: encodeMessageComponents(row.Components)}
}

func decodeActionsRow(row *discordgopb.ActionsRow) *discordgo.ActionsRow {
	if row == nil {
		return nil
	}

	return &discordgo.ActionsRow{Components: decodeMessageComponents(row.Components)}
}

func encodeButton(button *discordgo.Button) *discordgopb.Button {
	if button == nil {
		return nil
	}

	return &discordgopb.Button{
		Label:    button.Label,
		Style:    uint32(button.Style),
		Disabled: button.Disabled,
		Emoji:    encodeComponentEmoji(&button.Emoji),
		Url:      button.URL,
		CustomId: button.CustomID,
	}
}

func decodeButton(button *discordgopb.Button) *discordgo.Button {
	if button == nil {
		return nil
	}

	return &discordgo.Button{
		Label:    button.Label,
		Style:    discordgo.ButtonStyle(button.Style),
		Disabled: button.Disabled,
		Emoji:    *decodeComponentEmoji(button.Emoji),
		URL:      button.Url,
		CustomID: button.CustomId,
	}
}

func encodeSelectMenu(menu *discordgo.SelectMenu) *discordgopb.SelectMenu {
	if menu == nil {
		return nil
	}

	return &discordgopb.SelectMenu{
		MenuType:     uint32(menu.MenuType),
		CustomId:     menu.CustomID,
		Placeholder:  menu.Placeholder,
		MinValues:    util.PtrConv[int, int32](menu.MinValues),
		MaxValues:    int32(menu.MaxValues),
		Options:      encodeSelectMenuOptions(menu.Options),
		Disabled:     menu.Disabled,
		ChannelTypes: _encodeChannelTypes(menu.ChannelTypes),
	}
}

func decodeSelectMenu(menu *discordgopb.SelectMenu) *discordgo.SelectMenu {
	if menu == nil {
		return nil
	}

	return &discordgo.SelectMenu{
		MenuType:     discordgo.SelectMenuType(menu.MenuType),
		CustomID:     menu.CustomId,
		Placeholder:  menu.Placeholder,
		MinValues:    util.PtrConv[int32, int](menu.MinValues),
		MaxValues:    int(menu.MaxValues),
		Options:      decodeSelectMenuOptions(menu.Options),
		Disabled:     menu.Disabled,
		ChannelTypes: _decodeChannelTypes(menu.ChannelTypes),
	}
}

func encodeSelectMenuOption(option *discordgo.SelectMenuOption) *discordgopb.SelectMenuOption {
	if option == nil {
		return nil
	}

	return &discordgopb.SelectMenuOption{
		Label:       option.Label,
		Value:       option.Value,
		Description: option.Description,
		Emoji:       encodeComponentEmoji(&option.Emoji),
		Default:     option.Default,
	}
}

func encodeSelectMenuOptions(options []discordgo.SelectMenuOption) []*discordgopb.SelectMenuOption {
	if options == nil {
		return nil
	}

	result := make([]*discordgopb.SelectMenuOption, len(options))

	for i, option := range options {
		result[i] = encodeSelectMenuOption(&option)
	}

	return result
}

func decodeSelectMenuOption(option *discordgopb.SelectMenuOption) *discordgo.SelectMenuOption {
	if option == nil {
		return nil
	}

	return &discordgo.SelectMenuOption{
		Label:       option.Label,
		Value:       option.Value,
		Description: option.Description,
		Emoji:       *decodeComponentEmoji(option.Emoji),
		Default:     option.Default,
	}
}

func decodeSelectMenuOptions(options []*discordgopb.SelectMenuOption) []discordgo.SelectMenuOption {
	if options == nil {
		return nil
	}

	result := make([]discordgo.SelectMenuOption, len(options))

	for i, option := range options {
		result[i] = *decodeSelectMenuOption(option)
	}

	return result
}

func encodeTextInput(input *discordgo.TextInput) *discordgopb.TextInput {
	if input == nil {
		return nil
	}

	return &discordgopb.TextInput{
		CustomId:    input.CustomID,
		Label:       input.Label,
		Style:       uint32(input.Style),
		Placeholder: input.Placeholder,
		Value:       input.Value,
		Required:    input.Required,
		MinLength:   int32(input.MinLength),
		MaxLength:   int32(input.MaxLength),
	}
}

func decodeTextInput(input *discordgopb.TextInput) *discordgo.TextInput {
	if input == nil {
		return nil
	}

	return &discordgo.TextInput{
		CustomID:    input.CustomId,
		Label:       input.Label,
		Style:       discordgo.TextInputStyle(input.Style),
		Placeholder: input.Placeholder,
		Value:       input.Value,
		Required:    input.Required,
		MinLength:   int(input.MinLength),
		MaxLength:   int(input.MaxLength),
	}
}

func encodeComponentEmoji(emoji *discordgo.ComponentEmoji) *discordgopb.ComponentEmoji {
	if emoji == nil {
		return nil
	}

	return &discordgopb.ComponentEmoji{
		Name:     emoji.Name,
		Id:       emoji.ID,
		Animated: emoji.Animated,
	}
}

func decodeComponentEmoji(emoji *discordgopb.ComponentEmoji) *discordgo.ComponentEmoji {
	if emoji == nil {
		return nil
	}

	return &discordgo.ComponentEmoji{
		Name:     emoji.Name,
		ID:       emoji.Id,
		Animated: emoji.Animated,
	}
}
