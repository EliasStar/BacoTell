package proxy

import (
	"errors"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

// The default implementation for bacotell_common.SubmitProxy.
type submitProxy struct {
	interactionProxy
}

// submitProxy implements bacotell_common.SubmitProxy.
var _ common.SubmitProxy = submitProxy{}

// NewSubmitProxy returns a new SubmitProxy implementation.
func NewSubmitProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.SubmitProxy {
	return submitProxy{
		interactionProxy: interactionProxy{
			session:     session,
			interaction: interaction,
		},
	}
}

// InputValue implements bacotell_common.SubmitProxy.
func (p submitProxy) InputValue(customID string) (string, error) {
	component := _findComponent(p.interaction.ModalSubmitData().Components, customID)

	if component == nil {
		return "", errors.New("component not found")
	}

	switch cpt := component.(type) {
	case discordgo.TextInput:
		return cpt.Value, nil
	case *discordgo.TextInput:
		return cpt.Value, nil

	default:
		return "", errors.New("component is not a text input")
	}
}

// _findComponent returns the component with the specified custom id if it exists.
func _findComponent(components []discordgo.MessageComponent, customID string) discordgo.MessageComponent {
	for _, component := range components {
		var cptCustomID string

		switch cpt := component.(type) {
		case discordgo.ActionsRow:
			return _findComponent(cpt.Components, customID)
		case *discordgo.ActionsRow:
			return _findComponent(cpt.Components, customID)

		case discordgo.Button:
			cptCustomID = cpt.CustomID
		case *discordgo.Button:
			cptCustomID = cpt.CustomID

		case discordgo.SelectMenu:
			cptCustomID = cpt.CustomID
		case *discordgo.SelectMenu:
			cptCustomID = cpt.CustomID

		case discordgo.TextInput:
			cptCustomID = cpt.CustomID
		case *discordgo.TextInput:
			cptCustomID = cpt.CustomID

		default:
			continue
		}

		if cptCustomID == customID {
			return component
		}
	}

	return nil
}
