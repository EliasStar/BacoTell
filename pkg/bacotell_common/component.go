package bacotell_common

import "github.com/bwmarrin/discordgo"

// A Component represents a message component.
//
// It provides the component handler implementation for a specific custom id.
type Component interface {
	// CustomID returns the custom id associated with the component.
	//
	// To ensure that components from different plugins do not conflict with each other the custom id has to be prefixed with the plugin id.
	// While it is possible for this method to return an error, it is in most cases regarded as a programming error rather than typical behavior.
	CustomID() (string, error)

	// Handle gets called when a user interacts with a matching component.
	Handle(HandleProxy) error
}

// A HandleProxy provides methods for responding to the interaction as well as accessing interaction and component data and during component handling.
type HandleProxy interface {
	InteractionProxy

	// ComponentType returns the type of the component with which the user interacted.
	ComponentType() (discordgo.ComponentType, error)

	// SelectedValues returns the selected values of a select menu component or nil if the component is not a select menu.
	SelectedValues() ([]string, error)
}
