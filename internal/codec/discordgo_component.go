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
	result := make([]discordgo.MessageComponent, len(components))

	for i, component := range components {
		result[i] = decodeMessageComponent(component)
	}

	return result
}

func encodeActionsRow(row *discordgo.ActionsRow) *discordgopb.ActionsRow {
	return &discordgopb.ActionsRow{Components: encodeMessageComponents(row.Components)}
}

func decodeActionsRow(row *discordgopb.ActionsRow) *discordgo.ActionsRow {
	return &discordgo.ActionsRow{Components: decodeMessageComponents(row.Components)}
}

func encodeButton(button *discordgo.Button) *discordgopb.Button {
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
	return &discordgo.Button{
		Label:    button.Label,
		Style:    discordgo.ButtonStyle(button.Style),
		Disabled: button.Disabled,
		Emoji:    *decodeComponentEmoji(button.Emoji),
		URL:      button.Url,
		CustomID: button.CustomId,
	}
}

func encodeSelectMenu(input *discordgo.SelectMenu) *discordgopb.SelectMenu {
	return &discordgopb.SelectMenu{
		MenuType:     uint32(input.MenuType),
		CustomId:     input.CustomID,
		Placeholder:  input.Placeholder,
		MinValues:    util.PtrConv[int, int32](input.MinValues),
		MaxValues:    int32(input.MaxValues),
		Options:      encodeSelectMenuOptions(input.Options),
		Disabled:     input.Disabled,
		ChannelTypes: _encodeChannelTypes(input.ChannelTypes),
	}
}

func decodeSelectMenu(input *discordgopb.SelectMenu) *discordgo.SelectMenu {
	return &discordgo.SelectMenu{
		MenuType:     discordgo.SelectMenuType(input.MenuType),
		CustomID:     input.CustomId,
		Placeholder:  input.Placeholder,
		MinValues:    util.PtrConv[int32, int](input.MinValues),
		MaxValues:    int(input.MaxValues),
		Options:      decodeSelectMenuOptions(input.Options),
		Disabled:     input.Disabled,
		ChannelTypes: _decodeChannelTypes(input.ChannelTypes),
	}
}

func encodeSelectMenuOption(input *discordgo.SelectMenuOption) *discordgopb.SelectMenuOption {
	return &discordgopb.SelectMenuOption{
		Label:       input.Label,
		Value:       input.Value,
		Description: input.Description,
		Emoji:       encodeComponentEmoji(&input.Emoji),
		Default:     input.Default,
	}
}

func encodeSelectMenuOptions(inputs []discordgo.SelectMenuOption) []*discordgopb.SelectMenuOption {
	result := make([]*discordgopb.SelectMenuOption, len(inputs))

	for i, input := range inputs {
		result[i] = encodeSelectMenuOption(&input)
	}

	return result
}

func decodeSelectMenuOption(input *discordgopb.SelectMenuOption) *discordgo.SelectMenuOption {
	return &discordgo.SelectMenuOption{
		Label:       input.Label,
		Value:       input.Value,
		Description: input.Description,
		Emoji:       *decodeComponentEmoji(input.Emoji),
		Default:     input.Default,
	}
}

func decodeSelectMenuOptions(inputs []*discordgopb.SelectMenuOption) []discordgo.SelectMenuOption {
	result := make([]discordgo.SelectMenuOption, len(inputs))

	for i, input := range inputs {
		result[i] = *decodeSelectMenuOption(input)
	}

	return result
}

func encodeTextInput(input *discordgo.TextInput) *discordgopb.TextInput {
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
	return &discordgopb.ComponentEmoji{
		Name:     emoji.Name,
		Id:       emoji.ID,
		Animated: emoji.Animated,
	}
}

func decodeComponentEmoji(emoji *discordgopb.ComponentEmoji) *discordgo.ComponentEmoji {
	return &discordgo.ComponentEmoji{
		Name:     emoji.Name,
		ID:       emoji.Id,
		Animated: emoji.Animated,
	}
}
