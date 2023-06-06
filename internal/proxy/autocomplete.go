package proxy

import (
	"errors"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

// The default implementation for bacotell_common.AutocompleteProxy.
type autocompleteProxy struct {
	session     *discordgo.Session
	interaction *discordgo.Interaction
}

// autocompleteProxy implements bacotell_common.AutocompleteProxy.
var _ common.AutocompleteProxy = autocompleteProxy{}

// NewAutocompleteProxy returns a new AutocompleteProxy implementation.
func NewAutocompleteProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.AutocompleteProxy {
	return autocompleteProxy{
		session:     session,
		interaction: interaction,
	}
}

// Respond implements bacotell_common.AutocompleteProxy.
func (p autocompleteProxy) Respond(choices ...*discordgo.ApplicationCommandOptionChoice) error {
	return p.session.InteractionRespond(p.interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{Choices: choices},
	})
}

// FocusedOption implements bacotell_common.AutocompleteProxy.
func (p autocompleteProxy) FocusedOption() (string, any, error) {
	option := _findFocusedOption(p.interaction.ApplicationCommandData().Options)
	if option == nil {
		return "", nil, errors.New("no focused option")
	}

	return option.Name, option.Value, nil
}

// _findFocusedOption returns the focused option in an autocomplete interaction if it exists.
func _findFocusedOption(options []*discordgo.ApplicationCommandInteractionDataOption) *discordgo.ApplicationCommandInteractionDataOption {
	for _, opt := range options {
		if opt.Focused {
			return opt
		}

		if opt.Options != nil {
			return _findFocusedOption(opt.Options)
		}
	}

	return nil
}
