package bacotell

type Plugin interface {
	ID() (string, error)
	ApplicationCommands() ([]Command, error)
	MessageComponents() ([]Component, error)
	// Modals() ([]Modal, error)
}
