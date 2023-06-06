// Package bacotell_common provides interfaces and structs to be used by plugins and BacoTell alike.
package bacotell_common

// A Plugin is the top-level interface used by plugins to provide implementations for commands, components, etc to BacoTell.
//
// While it is possible for all methods to return an error, it is in most cases regarded as a programming error rather than typical behavior.
type Plugin interface {
	// ID returns a unique plugin identifier.
	ID() (string, error)

	// ApplicationCommands returns all available commands that this plugin provides.
	ApplicationCommands() ([]Command, error)

	// ApplicationCommands returns all available components that this plugin can handle.
	MessageComponents() ([]Component, error)

	// Modals returns all available modals that this plugin can respond to.
	Modals() ([]Modal, error)
}
