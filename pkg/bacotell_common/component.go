package bacotell_common

import "github.com/bwmarrin/discordgo"

type Component interface {
	CustomID() (string, error)
	Handle(HandleProxy) error
}

type HandleProxy interface {
	InteractionProxy

	ComponentType() (discordgo.ComponentType, error)
	SelectedValues() ([]string, error)
}
