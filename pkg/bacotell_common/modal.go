package bacotell_common

type Modal interface {
	CustomID() (string, error)
	Submit(SubmitProxy) error
}

type SubmitProxy interface {
	InteractionProxy
}
