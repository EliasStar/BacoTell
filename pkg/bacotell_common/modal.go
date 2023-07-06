package bacotell_common

// A Modal represents a Discord modal.
//
// It provides the modal handler implementation for a specific custom id.
type Modal interface {
	// CustomID returns the custom id associated with the modal.
	//
	// To ensure that modals from different plugins do not conflict with each other the custom id has to be prefixed with the plugin id.
	// While it is possible for this method to return an error, it is in most cases regarded as a programming error rather than typical behavior.
	CustomID() (string, error)

	// Submit gets called when a user submits a matching modal.
	Submit(SubmitProxy) error
}

// A SubmitProxy provides methods for responding to the interaction as well as accessing interaction and modal data and during modal handling.
type SubmitProxy interface {
	InteractionProxy

	// InputValue retrieves the value of a text input by custom id. Returns an error if it was not set or the component is not a text input.
	InputValue(customID string) (string, error)
}
